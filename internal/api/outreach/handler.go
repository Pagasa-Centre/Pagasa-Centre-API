package outreach

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
)

type OutreachHandler interface {
	All() http.HandlerFunc
}

type handler struct {
	logger          zap.SugaredLogger
	outreachService outreach.OutreachService
}

func NewOutreachHandler(logger zap.SugaredLogger, outreachService outreach.OutreachService) OutreachHandler {
	return &handler{
		logger:          logger,
		outreachService: outreachService,
	}
}

func (oh *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		outreaches, err := oh.outreachService.All(ctx)
		if err != nil {
			oh.logger.Infow("Failed to get all outreaches", "error", err)
			render.Json(w, http.StatusInternalServerError, err.Error())

			return
		}

		resp := dto.ToGetAllOutreachesResponse(outreaches)

		render.Json(w, http.StatusOK, resp)
	}
}
