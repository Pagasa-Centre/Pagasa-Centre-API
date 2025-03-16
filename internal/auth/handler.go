package authentication

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler interface {
	Register() http.HandlerFunc
}

type authHandler struct {
	logger      zap.SugaredLogger
	authService Service
}

func NewHandler(
	logger zap.SugaredLogger,
	service Service,
) AuthHandler {
	return &authHandler{
		logger:      logger,
		authService: service,
	}
}

func (h *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := h.authService.RegisterNewUser(ctx)
		if err != nil {
			h.logger.Errorw("Error registering new user", "error", err)
		}
		render.Json(w, http.StatusCreated, "Successfully created user")
	}
}
