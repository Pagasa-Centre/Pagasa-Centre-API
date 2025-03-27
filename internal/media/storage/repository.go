package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type MediaRepository interface {
	GetAll(ctx context.Context) ([]*entity.Medium, error)
	BulkInsert(ctx context.Context, media []*entity.Medium) error
}

type repository struct {
	db *sqlx.DB
}

func NewMediaRepository(db *sqlx.DB) MediaRepository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Medium, error) {
	db := r.db.DB
	return entity.Media().All(ctx, db)
}

func (r *repository) BulkInsert(ctx context.Context, media []*entity.Medium) error {
	db := r.db.DB

	for _, item := range media {
		// Skip if already exists
		exists, err := entity.Media(
			entity.MediumWhere.YoutubeVideoID.EQ(item.YoutubeVideoID),
		).Exists(ctx, db)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		err = item.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
