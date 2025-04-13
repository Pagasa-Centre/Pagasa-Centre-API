package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type ApprovalRepository interface {
	Insert(ctx context.Context, approval *entity.Approval) error
}

type repository struct {
	db *sqlx.DB
}

func NewApprovalRepository(db *sqlx.DB) ApprovalRepository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, approval *entity.Approval) error {
	err := approval.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
