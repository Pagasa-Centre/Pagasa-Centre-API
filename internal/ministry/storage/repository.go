package storage

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type MinistryRepository interface {
	GetAll(ctx context.Context) (entity.MinistrySlice, error)
	AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error
	GetMinistryLeaderByMinistryID(ctx context.Context, ministryID string) (*string, error)
}

type repository struct {
	db *sqlx.DB
}

func NewMinistryRepository(db *sqlx.DB) MinistryRepository {
	return &repository{
		db: db,
	}
}

func (repo *repository) AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error {
	db := repo.db.DB

	// Fetch the ministry to update
	ministry, err := entity.FindMinistry(ctx, db, ministryID)
	if err != nil {
		return err // returns sql.ErrNoRows if not found
	}

	ministry.LeaderID = null.StringFrom(userID)

	// Update the record in the DB
	_, err = ministry.Update(ctx, db, boil.Whitelist("leader_id"))

	return err
}

func (repo *repository) GetAll(ctx context.Context) (entity.MinistrySlice, error) {
	// sqlx.DB -> *sql.DB (for sqlboiler)
	db := repo.db.DB

	ministries, err := entity.Ministries().All(ctx, db)
	if err != nil {
		return nil, err
	}

	return ministries, nil
}

func (repo *repository) GetMinistryLeaderByMinistryID(ctx context.Context, ministryID string) (*string, error) {
	db := repo.db.DB

	ministry, err := entity.FindMinistry(ctx, db, ministryID)
	if err != nil {
		return nil, err
	}

	var leaderID string
	if ministry.LeaderID.Valid {
		leaderID = ministry.LeaderID.String
	} else {
		return nil, errors.New("ministry has no leader")
	}

	return &leaderID, nil
}
