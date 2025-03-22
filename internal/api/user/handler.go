package user

import (
	"net/http"
	"time"

	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/request"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/commoncontext"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

type UserHandler interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	UpdateDetails() http.HandlerFunc
}

type userHandler struct {
	logger       zap.SugaredLogger
	userService  user.UserService
	rolesService roles.RolesService
}

func NewUserHandler(
	logger zap.SugaredLogger,
	userService user.UserService,
	rolesService roles.RolesService,
) UserHandler {
	return &userHandler{
		logger:       logger,
		userService:  userService,
		rolesService: rolesService,
	}
}

func (h *userHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.RegisterRequest

		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Errorw("failed to decode and validate register request body",
				"context", ctx, "error", err)

			render.Json(w, http.StatusBadRequest, err.Error())

			return
		}

		userDomain, err := domain.RegisterRequestToUserDomain(req)
		if err != nil {
			h.logger.Errorw("failed to map register request to user domain",
				"context", ctx, "error", err)
			render.Json(w, http.StatusBadRequest, err.Error())

			return
		}

		userID, err := h.userService.RegisterNewUser(ctx, userDomain)
		if err != nil {
			h.logger.Errorw("Error registering new user", "error", err)

			render.Json(w, http.StatusInternalServerError, "internal server error")

			return
		}

		if req.IsLeader {
			err = h.rolesService.AssignLeaderRole(ctx, *userID)
			if err != nil {
				h.logger.Errorw("failed to assign leader role", "error", err)
				render.Json(w, http.StatusInternalServerError, err.Error())

				return
			}
		}

		if req.IsPrimary {
			err = h.rolesService.AssignPrimaryRole(ctx, *userID)
			if err != nil {
				h.logger.Errorw("failed to assign primary role", "error", err)
				render.Json(w, http.StatusInternalServerError, err.Error())

				return
			}
		}

		if req.IsPastor {
			err = h.rolesService.AssignPastorRole(ctx, *userID)
			if err != nil {
				h.logger.Errorw("failed to assign pastor role", "error", err)
				render.Json(w, http.StatusInternalServerError, err.Error())

				return
			}
		}

		render.Json(w, http.StatusCreated, "Successfully created user")
	}
}

func (h *userHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.LoginRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Errorw("failed to decode and validate login request body", "error", err)
			render.Json(w, http.StatusBadRequest, err.Error())

			return
		}

		userDomain, err := h.userService.GetUserByEmail(ctx, req.Email)
		if err != nil {
			h.logger.Errorw("failed to find user", "error", err)
			// For security, use the same error for "not found" or "wrong password"
			render.Json(w, http.StatusUnauthorized, "invalid credentials")

			return
		}

		// Compare the provided password with the stored hashed password.
		err = bcrypt.CompareHashAndPassword([]byte(userDomain.HashedPassword), []byte(req.Password))
		if err != nil {
			h.logger.Errorw("password mismatch", "error", err)
			render.Json(w, http.StatusUnauthorized, "invalid credentials")

			return
		}

		// Generate an authentication token (e.g., JWT)
		token, err := h.userService.GenerateToken(userDomain)
		if err != nil {
			h.logger.Errorw("failed to generate token", "error", err)
			render.Json(w, http.StatusInternalServerError, "internal server error")

			return
		}

		// Step 2e: Return the token in the response.
		render.Json(w, http.StatusOK, map[string]string{"token": token})
	}
}

func (h *userHandler) UpdateDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Decode and validate the update request.
		var req dto.UpdateDetailsRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Errorw("failed to decode update details request", "error", err)
			render.Json(w, http.StatusBadRequest, err.Error())

			return
		}

		// Get the user ID from the context
		userID, err := commoncontext.GetUserID(ctx)
		if err != nil {
			h.logger.Errorw("user ID not found in session", "error", err)
			render.Json(w, http.StatusUnauthorized, "unauthorized")

			return
		}

		// Retrieve the current user from the database.
		currentUser, err := h.userService.GetUserById(ctx, userID)
		if err != nil {
			h.logger.Errorw("failed to retrieve user", "error", err)
			render.Json(w, http.StatusInternalServerError, "internal server error")

			return
		}

		// Update the fields if provided in the request.
		if req.FirstName != "" {
			currentUser.FirstName = req.FirstName
		}

		if req.LastName != "" {
			currentUser.LastName = req.LastName
		}

		if req.Email != "" {
			currentUser.Email = req.Email
		}

		if req.PhoneNumber != "" {
			currentUser.Phone = null.StringFrom(req.PhoneNumber)
		}

		if req.Birthday != "" {
			parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
			if err != nil {
				h.logger.Errorw("failed to parse birthday", "error", err)
				render.Json(w, http.StatusBadRequest, "invalid birthday format; expected YYYY-MM-DD")

				return
			}

			currentUser.Birthday = null.TimeFrom(parsedBirthday)
		}

		if req.Password != "" {
			// Hash the new password before updating.
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				h.logger.Errorw("failed to hash new password", "error", err)
				render.Json(w, http.StatusInternalServerError, "failed to update password")

				return
			}

			currentUser.HashedPassword = string(hashedPassword)
		}

		if req.CellLeaderID != nil {
			currentUser.CellLeaderID = null.IntFrom(*req.CellLeaderID)
		}

		if req.OutreachID != 0 {
			currentUser.OutreachID = null.IntFrom(req.OutreachID)
		}

		// Call the service to update the user details.
		updatedUserDetails, err := h.userService.UpdateUserDetails(ctx, currentUser)
		if err != nil {
			h.logger.Errorw("failed to update user details", "error", err)
			render.Json(w, http.StatusInternalServerError, "failed to update user")

			return
		}

		resp := dto.ToResponse(updatedUserDetails)

		render.Json(w, http.StatusOK, map[string]interface{}{
			"message": "user details updated successfully",
			"data":    resp,
		})
	}
}
