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

type AuthHandler interface {
	Register() http.HandlerFunc
}

type userHandler struct {
	logger      zap.SugaredLogger
	userService user.Service
}

func NewHandler(
	logger zap.SugaredLogger,
	service user.Service,
) AuthHandler {
	return &userHandler{
		logger:      logger,
		userService: service,
	}
}

func (h *userHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.CreateUserRequest

		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Errorw("failed to decode and validate user request body", "context", ctx)

			render.Json(w, http.StatusBadRequest, "invalid body")

			return
		}

		userDomain := domain.CreateUserRequestToUserDomain(req)

		err := h.userService.RegisterNewUser(ctx, userDomain)
		if err != nil {
			h.logger.Errorw("Error registering new user", "error", err)
		}
		render.Json(w, http.StatusCreated, "Successfully created user")
	}
}
