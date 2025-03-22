package user

import (
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/request"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

type UserHandler interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
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
