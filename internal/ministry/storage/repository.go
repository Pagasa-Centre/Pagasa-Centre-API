package storage

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type MinistryRepository interface {
	GetAll(ctx context.Context) (entity.MinistrySlice, error)
}

type ministryRepository struct {
	db *sqlx.DB
}

func NewMinistryRepository(db *sqlx.DB) MinistryRepository {
	return &ministryRepository{
		db: db,
	}
}

func (repo *ministryRepository) GetAll(ctx context.Context) (entity.MinistrySlice, error) {
	// sqlx.DB -> *sql.DB (for sqlboiler)
	db := repo.db.DB

	ministries, err := entity.Ministries().All(ctx, db)
	if err != nil {
		return nil, err
	}

	return ministries, nil
}
