package user

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type UserHandler interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	UpdateDetails() http.HandlerFunc
	Delete() http.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	userService user.UserService
}

func NewUserHandler(
	logger *zap.Logger,
	userService user.UserService,
) UserHandler {
	return &handler{
		logger:      logger,
		userService: userService,
	}
}

const (
	InvalidInputMsg        = "Some required fields are missing or incorrectly filled out. Please check the form and try again."
	InternalServerErrorMsg = "Internal server error. Please try again later."
	InvalidCredentialsMsg  = "Incorrect email or password. Please try again."
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
				http.StatusBadRequest,
				mappers.ToRegisterResponse(
					nil,
					InvalidInputMsg,
				))

			return
		}

		userDomain, err := mappers.RegisterRequestToUserDomain(req)
		if err != nil {
			h.logger.Sugar().Errorw("failed to map register request to user domain",
				"context", ctx, "error", err)
			render.Json(w, http.StatusBadRequest,
				mappers.ToRegisterResponse(
					nil,
					InvalidInputMsg,
				),
			)

			return
		}

		result, err := h.userService.RegisterNewUser(ctx, userDomain, req)
		if err != nil {
			h.logger.Sugar().Errorw("Error registering new user", "error", err)

			statusCode, errMsg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(w, statusCode,
				mappers.ToRegisterResponse(
					nil,
					errMsg,
				),
			)

			return
		}

		resp := mappers.ToRegisterResponse(result, "Registration successful")

		render.Json(w, http.StatusCreated, resp)
	}
}

func (h *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.LoginRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode and validate login request body", "error", err)
			render.Json(w, http.StatusBadRequest,
				mappers.ToLoginResponse(
					nil,
					InvalidCredentialsMsg,
				),
			)

			return
		}

		authResult, err := h.userService.AuthenticateAndGenerateToken(ctx, req.Email, req.Password)
		if err != nil {
			h.logger.Sugar().Errorw("authentication failed", "error", err)

			statusCode, errMsg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(w, statusCode,
				mappers.ToLoginResponse(nil, errMsg),
			)

			return
		}

		resp := mappers.ToLoginResponse(authResult, "Successfully logged in")

		render.Json(w, http.StatusOK, resp)
	}
}

func (h *handler) UpdateDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Decode and validate the update request.
		var req dto.UpdateDetailsRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode update details request", "error", err)
			render.Json(
				w,
				http.StatusBadRequest,
				mappers.ToUpdateUserDetailsResponse(
					nil,
					InvalidInputMsg,
				),
			)

			return
		}

		// Call the service to update the user details.
		updatedUserDetails, err := h.userService.UpdateUserDetails(ctx, req)
		if err != nil {
			h.logger.Sugar().Errorw("failed to update user details", "error", err)
			render.Json(
				w,
				http.StatusInternalServerError,
				mappers.ToUpdateUserDetailsResponse(
					nil,
					"failed to update user details",
				))

			return
		}

		resp := mappers.ToUpdateUserDetailsResponse(updatedUserDetails, "Successfully updated user details")

		render.Json(w, http.StatusOK, resp)
	}
}

func (h *handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID, err := context.GetUserIDString(ctx)
		if err != nil {
			h.logger.Sugar().Errorw("user ID not found in session", "error", err)
			render.Json(
				w,
				http.StatusUnauthorized,
				mappers.ToDeleteUserResponse("unauthorized"))

			return
		}

		// Call the service to delete the user.
		if err := h.userService.DeleteUser(ctx, userID); err != nil {
			h.logger.Sugar().Errorw("failed to delete user", "error", err)
			render.Json(
				w,
				http.StatusInternalServerError,
				mappers.ToDeleteUserResponse("failed to delete user"),
			)

			return
		}

		// Return a success response.
		render.Json(w, http.StatusOK, mappers.ToDeleteUserResponse("user deleted successfully"))
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
		return http.StatusInternalServerError, InternalServerErrorMsg
	}
}
