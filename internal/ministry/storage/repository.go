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

func (r *repository) GetAll(ctx context.Context) (entity.MinistrySlice, error) {
	ministries, err := entity.Ministries().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return ministries, nil
}

func (r *repository) GetMinistryByID(ctx context.Context, ministryID string) (*entity.Ministry, error) {
	ministry, err := entity.FindMinistry(ctx, r.db, ministryID)
	if err != nil {
		return nil, err
	}

	return ministry, nil
}

func (r *repository) GetMinistryLeaderUsersByMinistryID(ctx context.Context, ministryID string) (entity.UserSlice, error) {
	leaders, err := entity.MinistryLeaders(
		entity.MinistryLeaderWhere.MinistryID.EQ(ministryID),
	).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get ministry leaders: %w", err)
	}

	for _, leader := range leaders {
		if err := leader.L.LoadUser(ctx, r.db, true, leader, nil); err != nil {
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

func (r *repository) GetMinistryActivitiesByMinistryID(ctx context.Context, ministryID string) (entity.MinistryActivitySlice, error) {
	activities, err := entity.MinistryActivities(
		entity.MinistryActivityWhere.MinistryID.EQ(ministryID),
	).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get ministry activities: %w", err)
	}

	return activities, nil
}
