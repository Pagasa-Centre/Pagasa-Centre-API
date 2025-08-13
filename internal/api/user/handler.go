package user

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	commonErrors "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/errors"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type UserHandler interface {
	UpdateDetails() http.HandlerFunc // UpdateDetails handles updating the logged-in user's profile details
	Delete() http.HandlerFunc        // Delete handles deleting the logged-in user's profile
}

type handler struct {
	logger      *zap.Logger
	userService user.Service
}

func NewUserHandler(
	logger *zap.Logger,
	userService user.Service,
) UserHandler {
	return &handler{
		logger:      logger,
		userService: userService,
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
				mapper.ToUpdateUserDetailsResponse(
					nil,
					commonErrors.InvalidInputMsg,
				),
			)

			return
		}

		// Call the service to update the user details.
		updatedUserDetails, err := h.userService.Update(ctx, req)
		if err != nil {
			h.logger.Sugar().Errorw("failed to update user details", "error", err)
			status, msg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(
				w,
				status,
				mapper.ToUpdateUserDetailsResponse(
					nil,
					msg,
				))

			return
		}

		render.Json(w,
			http.StatusOK,
			mapper.ToUpdateUserDetailsResponse(updatedUserDetails, "Successfully updated user details"),
		)
	}
}

func (h *handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID, err := context.GetUserID(ctx)
		if err != nil {
			h.logger.Sugar().Errorw("user ID not found in session", "error", err)
			render.Json(
				w,
				http.StatusUnauthorized,
				mapper.ToDeleteUserResponse("unauthorized"))

			return
		}

		// Call the service to delete the user.
		err = h.userService.Delete(ctx, userID)
		if err != nil {
			h.logger.Sugar().Errorw("failed to delete user", "error", err)

			status, msg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)
			render.Json(
				w,
				status,
				mapper.ToDeleteUserResponse(msg),
			)

			return
		}

		// Return a success response.
		render.Json(w, http.StatusOK, mapper.ToDeleteUserResponse("user deleted successfully"))
	}
}

func mapErrorsToStatusCodeAndUserFriendlyMessages(err error) (int, string) {
	switch {
	case errors.Is(err, user.ErrEmailAlreadyExists):
		return http.StatusConflict, "Email already exists."
	case errors.Is(err, user.ErrInvalidOutreach):
		return http.StatusBadRequest, "The selected outreach does not exist."
	default:
		return http.StatusInternalServerError, commonErrors.InternalServerErrorMsg
	}
}
