package user

import (
	"net/http"
	"time"

	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
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
	logger          *zap.Logger
	userService     user.UserService
	rolesService    roles.RolesService
	ministryService ministry.MinistryService
}

func NewUserHandler(
	logger *zap.Logger,
	userService user.UserService,
	rolesService roles.RolesService,
	ministryService ministry.MinistryService,
) UserHandler {
	return &handler{
		logger:          logger,
		userService:     userService,
		rolesService:    rolesService,
		ministryService: ministryService,
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

		userID, err := h.userService.RegisterNewUser(ctx, userDomain)
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

		if req.IsLeader {
			err = h.rolesService.AssignLeaderRole(ctx, *userID)
			if err != nil {
				h.logger.Sugar().Errorw("failed to assign leader role", "error", err)
				render.Json(w, http.StatusInternalServerError,
					dto.ToRegisterResponse(
						nil,
						nil,
						InternalServerErrorMsg,
					),
				)

				return
			}
		}

		if req.IsPrimary {
			err = h.rolesService.AssignPrimaryRole(ctx, *userID)
			if err != nil {
				h.logger.Sugar().Errorw("failed to assign primary role", "error", err)
				render.Json(
					w,
					http.StatusInternalServerError,
					dto.ToRegisterResponse(
						nil,
						nil,
						InternalServerErrorMsg,
					),
				)

				return
			}
		}

		if req.IsPastor {
			err = h.rolesService.AssignPastorRole(ctx, *userID)
			if err != nil {
				h.logger.Sugar().Errorw("failed to assign pastor role", "error", err)
				render.Json(
					w,
					http.StatusInternalServerError,
					dto.ToRegisterResponse(
						nil,
						nil,
						InternalServerErrorMsg,
					),
				)

				return
			}
		}

		if req.IsMinistryLeader {
			err = h.rolesService.AssignMinistryLeaderRole(ctx, *userID)
			if err != nil {
				h.logger.Sugar().Errorw("failed to assign pastor role", "error", err)
				render.Json(
					w,
					http.StatusInternalServerError,
					dto.ToRegisterResponse(
						nil,
						nil,
						InternalServerErrorMsg,
					),
				)

				return
			}

			err = h.ministryService.AssignLeaderToMinistry(ctx, *req.MinistryID, *userID)
			if err != nil {
				h.logger.Sugar().Errorw("failed to assign leader to ministry", "error", err)
				render.Json(
					w,
					http.StatusInternalServerError,
					dto.ToRegisterResponse(
						nil,
						nil,
						InternalServerErrorMsg,
					),
				)

				return
			}
		}

		userEntity, err := h.userService.GetUserById(ctx, *userID)
		if err != nil {
			h.logger.Sugar().Errorw("failed to get user by id", "error", err)
			render.Json(w, http.StatusInternalServerError,
				dto.ToRegisterResponse(
					nil,
					nil,
					InternalServerErrorMsg,
				),
			)
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

		userEntity, err := h.userService.GetUserByEmail(ctx, req.Email)
		if err != nil {
			h.logger.Sugar().Errorw("failed to find user", "error", err)
			// For security, use the same error for "not found" or "wrong password"
			render.Json(w, http.StatusUnauthorized,
				dto.ToRegisterResponse(
					nil,
					nil,
					InvalidCredentialsMsg,
				),
			)

			return
		}

		// Compare the provided password with the stored hashed password.
		err = bcrypt.CompareHashAndPassword([]byte(userEntity.HashedPassword), []byte(req.Password))
		if err != nil {
			h.logger.Sugar().Errorw("password mismatch", "error", err)
			render.Json(w, http.StatusUnauthorized,
				dto.ToRegisterResponse(
					nil,
					nil,
					InvalidCredentialsMsg,
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

		resp := dto.ToLoginResponse(userEntity, token, "Successfully logged in")

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

		// Get the user ID from the context
		userID, err := context.GetUserIDString(ctx)
		if err != nil {
			h.logger.Sugar().Errorw("user ID not found in session", "error", err)
			render.Json(
				w,
				http.StatusUnauthorized,
				dto.ToUpdateUserDetailsResponse(
					nil,
					"unauthorized",
				),
			)

			return
		}

		// Retrieve the current user from the database.
		currentUser, err := h.userService.GetUserById(ctx, userID)
		if err != nil {
			h.logger.Sugar().Errorw("failed to retrieve user", "error", err)
			render.Json(
				w,
				http.StatusInternalServerError,
				dto.ToUpdateUserDetailsResponse(
					nil,
					InternalServerErrorMsg,
				))

			return
		}

		// Update user fields based on the request.
		h.updateUserFields(currentUser, req)

		// Call the service to update the user details.
		updatedUserDetails, err := h.userService.UpdateUserDetails(ctx, currentUser)
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

// updateUserFields updates the provided user entity with the values from the update request.
func (h *handler) updateUserFields(user *entity.User, req dto.UpdateDetailsRequest) {
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.PhoneNumber != "" {
		user.Phone = null.StringFrom(req.PhoneNumber)
	}

	if req.Birthday != "" {
		parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			h.logger.Sugar().Errorw("failed to parse birthday", "error", err)
		} else {
			user.Birthday = null.TimeFrom(parsedBirthday)
		}
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			h.logger.Sugar().Errorw("failed to hash new password", "error", err)
		} else {
			user.HashedPassword = string(hashedPassword)
		}
	}

	if req.CellLeaderID != nil {
		user.CellLeaderID = null.StringFrom(*req.CellLeaderID)
	}

	if req.OutreachID != "" {
		user.OutreachID = null.StringFrom(req.OutreachID)
	}
}
