package approvals

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto"
	dto2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	context2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
)

type ApprovalService interface {
	CreateNewApproval(ctx context.Context, Approval *domain.Approval) error
	GetAllApprovals(ctx context.Context) ([]dto.Approval, error)
	SetUserService(us user.UserService)
}

type service struct {
	logger       *zap.Logger
	approvalRepo storage.ApprovalRepository
	userService  user.UserService
}

func NewApprovalService(
	logger *zap.Logger,
	approvalRepo storage.ApprovalRepository,
	userService user.UserService,
) ApprovalService {
	return &service{
		logger:       logger,
		approvalRepo: approvalRepo,
		userService:  userService,
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

func (s *service) GetAllApprovals(ctx context.Context) ([]dto.Approval, error) {
	// Get the user ID from the context
	userID, err := context2.GetUserIDString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	_, err = s.userService.GetUserById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
	}

	// 1. Get all approvals by requesterID
	approvalsSlice, err := s.approvalRepo.GetAll(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get approvals: %w", err)
	}
	// 2. For each Approval, get userDetails and add to response
	var approvals []dto.Approval

	for _, apr := range approvalsSlice {
		u, err := s.userService.GetUserById(ctx, apr.RequesterID)
		if err != nil {
			return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
		}

		var phoneNumber string
		if u.Phone.Valid {
			phoneNumber = u.Phone.String
		}

		approval := dto.Approval{
			ID:     apr.ID,
			Type:   apr.Type,
			Status: apr.Status,
			RequesterDetails: dto2.UserDetails{
				FirstName:   u.FirstName,
				LastName:    u.LastName,
				Email:       u.Email,
				PhoneNumber: phoneNumber,
			},
		}
		approvals = append(approvals, approval)
	}

	return approvals, nil
}

func (s *service) SetUserService(us user.UserService) {
	s.userService = us
}
