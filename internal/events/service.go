package events

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/storage"
)

type EventsService interface {
	GetAll(ctx context.Context) ([]*domain.Events, error)
	Create(ctx context.Context, events domain.Events) error
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

func (s *services) GetAll(ctx context.Context) ([]*domain.Events, error) {
	s.logger.Info("Fetching all events")

	eventEntities, err := s.repo.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}

	eventDays, err := s.repo.GetAllEventDays(ctx)
	if err != nil {
		return nil, err
	}

	events := mappers.EntityToEventsDomain(&eventEntities, &eventDays)
	if events == nil {
		return nil, errors.New("invalid event day date format")
	}

	return events, nil
}

func (s *services) Create(ctx context.Context, event domain.Events) error {
	s.logger.Info("Creating event")

	eventEntity := mappers.EventDomainToEntity(event)

	eventID, err := s.repo.CreateEvent(ctx, *eventEntity)
	if err != nil {
		return err
	}

	eventDaysEntity := mappers.EventDaysDomainToEntities(eventID, event.Days)

	err = s.repo.InsertEventDays(ctx, eventDaysEntity)
	if err != nil {
		return err
	}

	return nil
}
