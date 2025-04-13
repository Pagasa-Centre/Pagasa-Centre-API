package approvals

import (
	"context"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/storage"
)

type ApprovalService interface {
	CreateNewApproval(ctx context.Context, Approval *domain.Approval) error
}

type service struct {
	logger       *zap.Logger
	approvalRepo storage.ApprovalRepository
}

func NewApprovalService(
	logger *zap.Logger,
	approvalRepo storage.ApprovalRepository,
) ApprovalService {
	return &service{
		logger:       logger,
		approvalRepo: approvalRepo,
	}
}

func (s *service) CreateNewApproval(ctx context.Context, approval *domain.Approval) error {
	s.logger.Info("Creating new approval")

	entity := domain.ToApprovalEntity(approval)

	err := s.approvalRepo.Insert(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
