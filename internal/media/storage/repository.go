package storage

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"

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
	all, err := entity.Media().All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return all, nil
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
