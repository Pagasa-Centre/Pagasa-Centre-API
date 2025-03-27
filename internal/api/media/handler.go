package media

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
)

type MediaHandler interface {
	All() http.HandlerFunc
}

type handler struct {
	logger  zap.SugaredLogger
	service media.MediaService
}

func NewMediaHandler(logger zap.SugaredLogger, service media.MediaService) MediaHandler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		media, err := h.service.All(ctx)
		if err != nil {
			h.logger.Errorw("Failed to get all medias", "error", err)
			render.Json(w, http.StatusInternalServerError, err.Error())

			return
		}

		render.Json(w, http.StatusOK, map[string]any{"media": media})
	}
}
