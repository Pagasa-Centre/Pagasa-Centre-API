package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) (*int, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) InsertUser(ctx context.Context, user *entity.User) (*int, error) {
	if err := user.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, fmt.Errorf("failed inserting user entity: %w", err)
	}

	return &user.ID, nil
}
