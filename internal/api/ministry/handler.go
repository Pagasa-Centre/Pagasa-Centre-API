package ministry

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type MinistryHandler interface {
	All() http.HandlerFunc
}
type handler struct {
	logger          *zap.Logger
	MinistryService ministry.MinistryService
}

func NewMinistryHandler(logger *zap.Logger, ministryService ministry.MinistryService) MinistryHandler {
	return &handler{
		logger:          logger,
		MinistryService: ministryService,
	}
}

func (mh *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ministries, err := mh.MinistryService.All(ctx)
		if err != nil {
			mh.logger.Sugar().Infow("Failed to get all ministries", "error", err)
			render.Json(w, http.StatusInternalServerError, dto.ToErrorMinistriesResponse("Failed to fetch ministries"))
			return
		}

		resp := dto.ToGetAllMinistriesResponse(ministries, "Successfully fetched ministries")
		render.Json(w, http.StatusOK, resp)
	}
}
