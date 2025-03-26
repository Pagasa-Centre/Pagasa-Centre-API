package db

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	return goose.Up(db, "./migrations")
}
