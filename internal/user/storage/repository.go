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
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
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

// GetUserByEmail retrieves a user by their email address.
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := entity.Users(entity.UserWhere.Email.EQ(email)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email %s: %w", email, err)
	}

	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	return entity.Users(entity.UserWhere.ID.EQ(id)).One(ctx, r.db)
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := user.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}
