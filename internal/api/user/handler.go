package user

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/request"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler interface {
	Register() http.HandlerFunc
}

type userHandler struct {
	logger      zap.SugaredLogger
	userService user.UserService
}

func NewUserHandler(
	logger zap.SugaredLogger,
	service user.UserService,
) UserHandler {
	return &userHandler{
		logger:      logger,
		userService: service,
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

		err = h.userService.RegisterNewUser(ctx, userDomain)
		if err != nil {
			h.logger.Errorw("Error registering new user", "error", err)

			render.Json(w, http.StatusInternalServerError, "internal server error")
			return
		}

		render.Json(w, http.StatusCreated, "Successfully created user")
	}
}
