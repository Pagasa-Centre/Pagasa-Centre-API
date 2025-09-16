package outreach

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach/dto/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type Handler interface {
	AllOutreaches() http.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	outreachService outreach.Service
}

func NewOutreachHandler(logger *zap.Logger, outreachService outreach.Service) Handler {
	return &handler{
		logger:          logger,
		outreachService: outreachService,
	}
}

func (oh *handler) AllOutreaches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := oh.outreachService.GetAllOutreaches(ctx)
		if err != nil {
			oh.logger.Sugar().Infow("Failed to get all outreaches", "error", err)
			render.Json(w, http.StatusInternalServerError, mappers.ToGetAllOutreachesResponse(nil, "Failed to fetch outreaches"))

			return
		}

		resp := mappers.ToGetAllOutreachesResponse(result, "Successfully fetched outreaches")
		render.Json(w, http.StatusOK, resp)
	}
}
