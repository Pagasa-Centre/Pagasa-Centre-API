package events

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/events/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type EventHandler interface {
	All() http.HandlerFunc
}

type handler struct {
	logger        *zap.Logger
	eventsService events.EventsService
}

func NewEventsHandler(
	logger *zap.Logger,
	eventsService events.EventsService,
) EventHandler {
	return &handler{
		logger:        logger,
		eventsService: eventsService,
	}
}

const InternalServerErrorMsg = "Internal server error. Please try again later."

func (h *handler) All() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		e, err := h.eventsService.GetAll(ctx)
		if err != nil {
			h.logger.Error("Failed to get events", zap.Error(err))
			render.Json(
				w,
				http.StatusInternalServerError,
				dto.ToGetAllEventsResponse(nil, InternalServerErrorMsg),
			)
		}

		render.Json(
			w,
			http.StatusOK,
			dto.ToGetAllEventsResponse(*e, "Successfully fetched events."),
		)
	})
}
