package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type MinistryRepository interface {
	GetAll(ctx context.Context) (entity.MinistrySlice, error)
	GetMinistryByID(ctx context.Context, ministryID string) (*entity.Ministry, error)
	GetMinistryLeaderUsersByMinistryID(ctx context.Context, ministryID string) (entity.UserSlice, error)
	GetMinistryActivitiesByMinistryID(ctx context.Context, ministryID string) (entity.MinistryActivitySlice, error)
}

type repository struct {
	db *sqlx.DB
}

func NewMinistryRepository(db *sqlx.DB) MinistryRepository {
	return &repository{
		db: db,
	}
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

func (repo *repository) GetMinistryByID(ctx context.Context, ministryID string) (*entity.Ministry, error) {
	db := repo.db.DB

	ministry, err := entity.FindMinistry(ctx, db, ministryID)
	if err != nil {
		return nil, err
	}

	return ministry, nil
}

func (repo *repository) GetMinistryLeaderUsersByMinistryID(ctx context.Context, ministryID string) (entity.UserSlice, error) {
	db := repo.db.DB

	leaders, err := entity.MinistryLeaders(
		entity.MinistryLeaderWhere.MinistryID.EQ(ministryID),
	).All(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to get ministry leaders: %w", err)
	}

	for _, leader := range leaders {
		if err := leader.L.LoadUser(ctx, db, true, leader, nil); err != nil {
			return nil, fmt.Errorf("failed to load user: %w", err)
		}
	}

	var users entity.UserSlice

	for _, leader := range leaders {
		if leader.R != nil && leader.R.User != nil {
			users = append(users, leader.R.User)
		}
	}

	return users, nil
}

func (repo *repository) GetMinistryActivitiesByMinistryID(ctx context.Context, ministryID string) (entity.MinistryActivitySlice, error) {
	db := repo.db.DB

	activities, err := entity.MinistryActivities(
		entity.MinistryActivityWhere.MinistryID.EQ(ministryID),
	).All(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to get ministry activities: %w", err)
	}

	return activities, nil
}
