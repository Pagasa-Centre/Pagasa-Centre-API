package user

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
	"go.uber.org/zap"
	"net/http"
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
	InvalidInputMsg        = "Please check your input. Some required fields might be missing or incorrectly formatted."
	InternalServerErrorMsg = "Internal server error. Please try again later."
	InvalidCredentialsMsg  = "Please check your username and password."
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
				dto.ToRegisterResponse(
					nil,
					nil,
					InvalidInputMsg,
				))

			return
		}

		userDomain, err := domain.RegisterRequestToUserDomain(req)
		if err != nil {
			h.logger.Sugar().Errorw("failed to map register request to user domain",
				"context", ctx, "error", err)
			render.Json(w, http.StatusBadRequest,
				dto.ToRegisterResponse(
					nil,
					nil,
					InvalidInputMsg,
				),
			)

			return
		}

		userEntity, err := h.userService.RegisterNewUser(ctx, userDomain, req)
		if err != nil {
			h.logger.Sugar().Errorw("Error registering new user", "error", err)

			render.Json(w, http.StatusInternalServerError,
				dto.ToRegisterResponse(
					nil,
					nil,
					InternalServerErrorMsg,
				),
			)

			return
		}

		// Generate an authentication token (e.g., JWT)
		token, err := h.userService.GenerateToken(userEntity)
		if err != nil {
			h.logger.Sugar().Errorw("failed to generate token", "error", err)
			render.Json(w, http.StatusInternalServerError,
				dto.ToRegisterResponse(
					nil,
					nil,
					InternalServerErrorMsg,
				),
			)

			return
		}

		resp := dto.ToRegisterResponse(userEntity, &token, "Registration successful")

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
				dto.ToRegisterResponse(
					nil,
					nil,
					InvalidCredentialsMsg,
				),
			)

			return
		}

		authResult, err := h.userService.AuthenticateAndGenerateToken(ctx, req.Email, req.Password)
		if err != nil {
			h.logger.Sugar().Errorw("authentication failed", "error", err)
			render.Json(w, http.StatusUnauthorized,
				dto.ToRegisterResponse(nil, nil, InvalidCredentialsMsg),
			)
			return
		}

		resp := dto.ToLoginResponse(authResult.User, authResult.Token, "Successfully logged in")

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
				dto.ToUpdateUserDetailsResponse(
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
				dto.ToUpdateUserDetailsResponse(
					nil,
					"failed to update user details",
				))

			return
		}

		resp := dto.ToUpdateUserDetailsResponse(updatedUserDetails, "Successfully updated user details")

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
				dto.ToDeleteUserResponse("unauthorized"))

			return
		}

		// Call the service to delete the user.
		if err := h.userService.DeleteUser(ctx, userID); err != nil {
			h.logger.Sugar().Errorw("failed to delete user", "error", err)
			render.Json(
				w,
				http.StatusInternalServerError,
				dto.ToDeleteUserResponse("failed to delete user"),
			)

			return
		}

		// Return a success response.
		render.Json(w, http.StatusOK, dto.ToDeleteUserResponse("user deleted successfully"))
	}
}
