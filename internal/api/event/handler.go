package event

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/event/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/event/dto/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type Handler interface {
	All() http.HandlerFunc
	Create() http.HandlerFunc
}

type handler struct {
	logger        *zap.Logger
	eventsService event.EventsService
}

func NewEventsHandler(
	logger *zap.Logger,
	eventsService event.EventsService,
) Handler {
	return &handler{
		logger:        logger,
		eventsService: eventsService,
	}
}

const InternalServerErrorMsg = "Internal server error. Please try again later."

func (h *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := h.eventsService.GetAll(ctx)
		if err != nil {
			h.logger.Error("Failed to get events", zap.Error(err))
			render.Json(
				w,
				http.StatusInternalServerError,
				mapper.ToGetAllEventsResponse(nil, InternalServerErrorMsg),
			)

			return
		}

		render.Json(
			w,
			http.StatusOK,
			mapper.ToGetAllEventsResponse(result, "Successfully fetched events."),
		)
	}
}

func (h *handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.CreateEventRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode and validate create event request", "error", err)
			render.Json(w, http.StatusBadRequest,
				mapper.ToCreateEventResponse(
					"Please double check your event and try again",
				),
			)

			return
		}

		eventDomain := mapper.CreateEventRequestToDomain(req)

		err := h.eventsService.Create(ctx, eventDomain)
		if err != nil {
			h.logger.Sugar().Errorw("failed to create event", "error", err)
			render.Json(w, http.StatusInternalServerError,
				mapper.ToCreateEventResponse(
					InternalServerErrorMsg,
				),
			)

			return
		}

		render.Json(
			w,
			http.StatusCreated,
			mapper.ToCreateEventResponse(
				"Event successfully created",
			))
	}
}
