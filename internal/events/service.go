package events

import (
	"context"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/storage"
)

type EventsService interface {
	GetAll(ctx context.Context) (*[]domain.Events, error)
}

type services struct {
	logger *zap.Logger
	repo   storage.EventsRepository
}

func NewEventsService(logger *zap.Logger, eventsRepo storage.EventsRepository) EventsService {
	return &services{
		logger: logger,
		repo:   eventsRepo,
	}
}

func (s *services) GetAll(ctx context.Context) (*[]domain.Events, error) {
	s.logger.Info("Fetching all events")

	eventEntities, err := s.repo.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}

	eventDays, err := s.repo.GetAllEventDays(ctx)
	if err != nil {
		return nil, err
	}

	events := domain.EntityToEventsDomain(&eventEntities, &eventDays)

	return events, nil
}
