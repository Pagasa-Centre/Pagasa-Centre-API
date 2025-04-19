package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type EventsRepository interface {
	GetAllEvents(ctx context.Context) (entity.EventSlice, error)
	GetAllEventDays(ctx context.Context) (entity.EventDaySlice, error)
	CreateEvent(ctx context.Context, event entity.Event) (string, error)
	InsertEventDays(ctx context.Context, eventDays []entity.EventDay) error
}

type repository struct {
	db *sqlx.DB
}

func NewEventsRepository(db *sqlx.DB) EventsRepository {
	return &repository{db: db}
}

func (r *repository) GetAllEvents(ctx context.Context) (entity.EventSlice, error) {
	db := r.db.DB
	return entity.Events().All(ctx, db)
}

func (r *repository) GetAllEventDays(ctx context.Context) (entity.EventDaySlice, error) {
	db := r.db.DB
	return entity.EventDays().All(ctx, db)
}

func (r *repository) CreateEvent(ctx context.Context, event entity.Event) (string, error) {
	db := r.db.DB

	err := event.Insert(ctx, db, boil.Infer())
	if err != nil {
		return "", fmt.Errorf("failed inserting event: %w", err)
	}

	return event.ID, nil
}

func (r *repository) InsertEventDays(ctx context.Context, eventDays []entity.EventDay) error {
	db := r.db.DB

	for _, item := range eventDays {
		err := item.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
