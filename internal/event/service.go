package event

import (
	"context"
	"errors"

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
		return nil, err
	}

	eventDays, err := s.repo.GetAllEventDays(ctx)
	if err != nil {
		return nil, err
	}

	events := mapper.EntityToEventsDomain(&eventEntities, &eventDays)
	if events == nil {
		return nil, errors.New("invalid event day date format")
	}

	return events, nil
}

func (s *services) Create(ctx context.Context, event domain.Events) error {
	eventEntity := mapper.EventDomainToEntity(event)

	eventID, err := s.repo.CreateEvent(ctx, *eventEntity)
	if err != nil {
		return err
	}

	eventDaysEntity := mapper.EventDaysDomainToEntities(eventID, event.Days)

	err = s.repo.InsertEventDays(ctx, eventDaysEntity)
	if err != nil {
		return err
	}

	return nil
}
