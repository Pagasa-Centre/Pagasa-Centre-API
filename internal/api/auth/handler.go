package auth

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto/mapper"
	auth2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	commonErrors "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/errors"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type Handler interface {
	Register() http.HandlerFunc // Register handles registering a new user
	Login() http.HandlerFunc    // Login handles logging in users
}

type handler struct {
	logger      *zap.Logger
	authService auth2.Service
}

func NewAuthHandler(
	logger *zap.Logger,
	authService auth2.Service,
) Handler {
	return &handler{
		logger:      logger,
		authService: authService,
	}
}

const (
	InvalidCredentialsMsg = "Incorrect email or password. Please try again."
)

func (h *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.RegisterRequest

		// Validates and decodes request
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode and validate register request body",
				"context", ctx, "error", err)

			render.Json(
				w,
				http.StatusBadRequest, //todo: create mapper to friendly errors in all handlers
				mapper.MapAuthResultToDTO(
					nil,
					commonErrors.InvalidInputMsg,
				))

			return
		}

		registrationDetails, err := mapper.MapRegisterRequestToDomain(req)
		if err != nil {
			h.logger.Sugar().Errorw("failed to map register request to register input",
				"context", ctx, "error", err)
			render.Json(w, http.StatusBadRequest,
				mapper.MapAuthResultToDTO(
					nil,
					commonErrors.InvalidInputMsg,
				),
			)

			return
		}

		authResult, err := h.authService.Register(ctx, registrationDetails)
		if err != nil {
			h.logger.Sugar().Errorw("Error registering new user", "error", err)

			statusCode, errMsg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(w, statusCode,
				mapper.MapAuthResultToDTO(
					nil,
					errMsg,
				),
			)

			return
		}

		render.Json(w, http.StatusCreated, mapper.MapAuthResultToDTO(authResult, "Registration successful"))
	}
}

func (h *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.LoginRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode and validate login request body", "error", err)
			render.Json(w, http.StatusBadRequest,
				mapper.MapAuthResultToDTO(
					nil,
					InvalidCredentialsMsg,
				),
			)

			return
		}

		authResult, err := h.authService.Login(ctx, mapper.MapLoginRequestToDomain(req))
		if err != nil {
			h.logger.Sugar().Errorw("authentication failed", "error", err)

			statusCode, errMsg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(w, statusCode,
				mapper.MapAuthResultToDTO(nil, errMsg),
			)

			return
		}

		render.Json(w, http.StatusOK, mapper.MapAuthResultToDTO(authResult, "Successfully logged in"))
	}
}

func mapErrorsToStatusCodeAndUserFriendlyMessages(err error) (int, string) {
	switch {
	case errors.Is(err, user.ErrEmailAlreadyExists):
		return http.StatusConflict, "Email already exists."
	case errors.Is(err, user.ErrInvalidOutreach):
		return http.StatusBadRequest, "The selected outreach does not exist."
	case errors.Is(err, user.ErrInvalidLoginDetails):
		return http.StatusBadRequest, InvalidCredentialsMsg
	default:
		return http.StatusInternalServerError, commonErrors.InternalServerErrorMsg
	}
}
