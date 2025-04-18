package storage

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type EventsRepository interface {
	GetAllEvents(ctx context.Context) (entity.EventSlice, error)
	GetAllEventDays(ctx context.Context) (entity.EventDaySlice, error)
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
