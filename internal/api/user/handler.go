package user

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/request"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

type UserHandler interface {
	Register() http.HandlerFunc
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
