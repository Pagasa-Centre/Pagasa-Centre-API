package event

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/storage"
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
	eventEntities, err := s.repo.GetAllEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all events: %w", err)
	}

	eventDays, err := s.repo.GetAllEventDays(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all event days: %w", err)
	}

	events, err := mapper.EntityToEventsDomain(&eventEntities, &eventDays)
	if err != nil {
		return nil, fmt.Errorf("failed to map events: %w", err)
	}

	return events, nil
}

func (s *services) Create(ctx context.Context, event domain.Events) error {
	eventEntity := mapper.EventDomainToEntity(event)

	eventID, err := s.repo.CreateEvent(ctx, *eventEntity)
	if err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}

	eventDaysEntity := mapper.EventDaysDomainToEntities(eventID, event.Days)

	err = s.repo.InsertEventDays(ctx, eventDaysEntity)
	if err != nil {
		return fmt.Errorf("failed to insert event days: %w", err)
	}

	return nil
}
