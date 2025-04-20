package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type ApprovalRepository interface {
	Insert(ctx context.Context, approval *entity.Approval) error
	GetAllPendingApprovalsByUserID(ctx context.Context, userID string) (entity.ApprovalSlice, error)
	GetAllPendingApprovals(ctx context.Context) (entity.ApprovalSlice, error)
	GetApprovalByID(ctx context.Context, approvalID string) (*entity.Approval, error)
	Update(ctx context.Context, approval *entity.Approval) error
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

func (r *repository) GetAllPendingApprovalsByUserID(ctx context.Context, userID string) (entity.ApprovalSlice, error) {
	approvals, err := entity.Approvals(
		entity.ApprovalWhere.ApproverID.EQ(null.StringFrom(userID)),
		entity.ApprovalWhere.Status.EQ(domain.Pending),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return approvals, nil
}

func (r *repository) GetApprovalByID(ctx context.Context, approvalID string) (*entity.Approval, error) {
	approval, err := entity.Approvals(entity.ApprovalWhere.ID.EQ(approvalID)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return approval, nil
}

func (r *repository) Update(ctx context.Context, approval *entity.Approval) error {
	_, err := approval.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllPendingApprovals(ctx context.Context) (entity.ApprovalSlice, error) {
	approvals, err := entity.Approvals(
		entity.ApprovalWhere.Status.EQ(domain.Pending),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return approvals, nil
}
