package storage

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type OutreachRepository interface {
	GetAll(ctx context.Context) (entity.OutreachSlice, error)
	GetServicesByOutreachID(ctx context.Context, outreachID string) (entity.OutreachServiceSlice, error)
}

type repository struct {
	db *sqlx.DB
}

func NewOutreachRepository(db *sqlx.DB) OutreachRepository {
	return &repository{
		db: db,
	}
}

func (or *repository) GetAll(ctx context.Context) (entity.OutreachSlice, error) {
	// sqlx.DB -> *sql.DB (for sqlboiler)
	db := or.db.DB

	outreaches, err := entity.Outreaches().All(ctx, db)
	if err != nil {
		return nil, err
	}

	return outreaches, nil
}

func (os *repository) GetServicesByOutreachID(ctx context.Context, outreachID string) (entity.OutreachServiceSlice, error) {
	db := os.db.DB

	services, err := entity.OutreachServices(entity.OutreachServiceWhere.OutreachID.EQ(outreachID)).All(ctx, db)
	if err != nil {
		return nil, err
	}

	return services, nil
}
