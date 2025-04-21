package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type ApprovalRepository interface {
	Insert(ctx context.Context, approval *entity.Approval) error
	GetAllPendingApprovalsByRequestedRole(ctx context.Context, requestedRole string) (entity.ApprovalSlice, error)
	GetAllPendingApprovals(ctx context.Context) (entity.ApprovalSlice, error)
	GetApprovalByID(ctx context.Context, approvalID string) (*entity.Approval, error)
	Update(ctx context.Context, approval *entity.Approval) error
	BeginTx(ctx context.Context) (*sqlx.Tx, error)
	GetApprovalByIDTx(ctx context.Context, tx *sqlx.Tx, approvalID string) (*entity.Approval, error)
	UpdateTx(ctx context.Context, tx *sqlx.Tx, approval *entity.Approval) error
}

type repository struct {
	db *sqlx.DB
}

func NewApprovalRepository(db *sqlx.DB) ApprovalRepository {
	return &repository{db: db}
}

func (r *repository) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.BeginTxx(ctx, nil)
}

func (r *repository) UpdateTx(ctx context.Context, tx *sqlx.Tx, approval *entity.Approval) error {
	_, err := approval.Update(ctx, tx, boil.Infer())
	return err
}

func (r *repository) GetApprovalByIDTx(ctx context.Context, tx *sqlx.Tx, approvalID string) (*entity.Approval, error) {
	return entity.Approvals(entity.ApprovalWhere.ID.EQ(approvalID)).One(ctx, tx)
}

func (r *repository) Insert(ctx context.Context, approval *entity.Approval) error {
	err := approval.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllPendingApprovalsByRequestedRole(ctx context.Context, requestedRole string) (entity.ApprovalSlice, error) {
	approvals, err := entity.Approvals(
		entity.ApprovalWhere.RequestedRole.EQ(requestedRole),
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
