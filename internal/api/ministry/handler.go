package ministry

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
)

type MinistryHandler interface {
	All() http.HandlerFunc
}
type ministryHandler struct {
	logger          zap.SugaredLogger
	MinistryService ministry.MinistryService
}

func NewMinistryHandler(logger zap.SugaredLogger, ministryService ministry.MinistryService) MinistryHandler {
	return &ministryHandler{
		logger:          logger,
		MinistryService: ministryService,
	}
}

func (mh *ministryHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ministries, err := mh.MinistryService.All(ctx)
		if err != nil {
			mh.logger.Infow("Failed to get all ministries", "error", err)
			render.Json(w, http.StatusInternalServerError, err.Error())

			return
		}

		resp := dto.ToGetAllMinistriesResponse(ministries)

		render.Json(w, http.StatusOK, resp)
	}
}
