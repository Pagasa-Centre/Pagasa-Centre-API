package storage

import (
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=storage
type AuthRepository interface{}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
