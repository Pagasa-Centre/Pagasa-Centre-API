package ministry

import (
	"net/http"

	"go.uber.org/zap"

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
		render.Json(w, http.StatusOK, "here G")
	}
}
