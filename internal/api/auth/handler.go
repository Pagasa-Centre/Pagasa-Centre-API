package auth

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/request"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler interface {
	Register() http.HandlerFunc
}

type authHandler struct {
	logger      zap.SugaredLogger
	authService authentication.Service
}

func NewHandler(
	logger zap.SugaredLogger,
	service authentication.Service,
) AuthHandler {
	return &authHandler{
		logger:      logger,
		authService: service,
	}
}

func (h *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.CreateUserRequest

		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Errorw("failed to decode and validate user request body", "context", ctx)

			render.Json(w, http.StatusBadRequest, "invalid body")

			return
		}

		userDomain := domain.CreateUserRequestToUserDomain(req)

		err := h.authService.RegisterNewUser(ctx, userDomain)
		if err != nil {
			h.logger.Errorw("Error registering new user", "error", err)
		}
		render.Json(w, http.StatusCreated, "Successfully created user")
	}
}
