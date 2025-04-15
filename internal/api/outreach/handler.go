package outreach

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type OutreachHandler interface {
	GetAllOutreaches() http.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	outreachService outreach.OutreachService
}

func NewOutreachHandler(logger *zap.Logger, outreachService outreach.OutreachService) OutreachHandler {
	return &handler{
		logger:          logger,
		outreachService: outreachService,
	}
}

func (oh *handler) GetAllOutreaches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		outreaches, services, err := oh.outreachService.GetAllOutreaches(ctx)
		if err != nil {
			oh.logger.Sugar().Infow("Failed to get all outreaches", "error", err)
			render.Json(w, http.StatusInternalServerError, dto.ToErrorOutreachesResponse("Failed to fetch outreaches"))

			return
		}

		resp := dto.ToGetAllOutreachesResponse(outreaches, services, "Successfully fetched outreaches")
		render.Json(w, http.StatusOK, resp)
	}
}
