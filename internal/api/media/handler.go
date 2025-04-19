package media

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/media/dto/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type MediaHandler interface {
	All() http.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	service media.MediaService
}

func NewMediaHandler(logger *zap.Logger, service media.MediaService) MediaHandler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

const InternalServerErrorMsg = "Internal server error. Please try again later."

func (h *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		m, err := h.service.All(ctx)
		if err != nil {
			h.logger.Sugar().Errorw("Failed to get all medias", "error", err)
			render.Json(
				w,
				http.StatusInternalServerError,
				mappers.ToGetAllMediaResponse(nil, InternalServerErrorMsg),
			)

			return
		}

		render.Json(
			w,
			http.StatusOK,
			mappers.ToGetAllMediaResponse(m, "Media fetched successfully"),
		)
	}
}
