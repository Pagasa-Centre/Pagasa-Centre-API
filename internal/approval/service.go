package approval

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	roledomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role/domain"
	userservice "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	userdomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	commoncontext "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/interfaces"
)

type Service interface {
	CreateNewApproval(ctx context.Context, Approval *domain.Approval) error
	GetAll(ctx context.Context) (*GetAllResult, error)
	UpdateApprovalStatus(ctx context.Context, approvalID, status string) error
}

type service struct {
	logger       *zap.Logger
	approvalRepo storage.ApprovalRepository
	userService  userservice.Service
	roleService  interfaces.RoleService
}

func NewApprovalService(
	logger *zap.Logger,
	approvalRepo storage.ApprovalRepository,
	userService userservice.Service,
	roleService interfaces.RoleService,
) Service {
	return &service{
		logger:       logger,
		approvalRepo: approvalRepo,
		userService:  userService,
		roleService:  roleService,
	}
}

type (
	GetAllResult struct {
		Approvals []*domain.Approval
		Users     []*userdomain.User
	}
)

var ErrNoPermission = errors.New("you do not have permission to update this approval")

func (s *service) CreateNewApproval(ctx context.Context, approval *domain.Approval) error {
	approvalEntity := mapper.ToApprovalEntity(approval)

	err := s.approvalRepo.Insert(ctx, approvalEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context) (*GetAllResult, error) {
	// Get the user ID from the context
	userID, err := commoncontext.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	_, err = s.userService.GetById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
	}

	// Get user roles
	userRoles, err := s.roleService.Fetch(ctx, userID)
	if err != nil {
		return nil, err
	}

	var approvalsSlice entity.ApprovalSlice

	// Role-based logic
	switch {
	case slices.Contains(userRoles, string(roledomain.RoleAdmin)):
		approvalsSlice, err = s.approvalRepo.GetAllPendingApprovals(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get approvals: %w", err)
		}
	default:
		// get all approvals a user is able to approve
		for leaderRole, memberRole := range roledomain.RoleToLeaderMap {
			if slices.Contains(userRoles, string(leaderRole)) {
				roleApprovals, err := s.approvalRepo.GetAllPendingApprovalsByRequestedRole(ctx, string(memberRole))
				if err != nil {
					return nil, fmt.Errorf("failed to get approvals for %s: %w", memberRole, err)
				}

				approvalsSlice = append(approvalsSlice, roleApprovals...)
			}
		}
	}

	approvalsDomain := mapper.EntityToDomainApprovals(approvalsSlice)

	// 2.Get All the details of the requester on each approval
	var users []*userdomain.User

	for _, apr := range approvalsSlice {
		u, err := s.userService.GetById(ctx, apr.RequesterID)
		if err != nil {
			return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
		}

		users = append(users, u)
	}

	result := &GetAllResult{
		Approvals: approvalsDomain,
		Users:     users,
	}

	return result, nil
}

func (s *service) UpdateApprovalStatus(ctx context.Context, approvalID, status string) error {
	s.logger.With(zap.String("approvalID", approvalID), zap.String("status", status)).Info("Updating approval status")
	// Get the user ID from the context
	userID, err := commoncontext.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get user id from context: %w", err)
	}

	// check if user exists in db.
	_, err = s.userService.GetById(ctx, userID)
	if err != nil {
		return fmt.Errorf("user does not exist or has been deleted: %w", err)
	}

	// Get user roles
	userRoles, err := s.roleService.Fetch(ctx, userID)
	if err != nil {
		return err
	}

	// ✅ Check permissions: must be Admin or have a role with "Leader" in it
	isAdminOrLeader := false

	for _, role := range userRoles {
		if role == "Admin" || strings.Contains(role, "Leader") {
			isAdminOrLeader = true
			break
		}
	}

	if !isAdminOrLeader {
		return ErrNoPermission
	}

	tx, err := s.approvalRepo.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback() // safe rollback; will be a no-op if committed
	}()

	// 1. Get approval by id
	approvalEntity, err := s.approvalRepo.GetApprovalByIDTx(ctx, tx, approvalID)
	if err != nil {
		return err
	}

	// todo: don't update entity. convert to domain, manipulate, convert back entity
	switch status {
	case string(domain.Approved):
		approvalEntity.Status = string(domain.Approved)
	case string(domain.Rejected):
		approvalEntity.Status = string(domain.Rejected)
	default:
		return fmt.Errorf("invalid approval status: %s", status)
	}

	approvalEntity.UpdatedBy = null.StringFrom(userID)

	// 2. Update status in transaction
	err = s.approvalRepo.UpdateTx(ctx, tx, approvalEntity)
	if err != nil {
		return err
	}

	if approvalEntity.Status == string(domain.Approved) {
		var ministryID *string
		if approvalEntity.MinistryID.Valid {
			ministryID = &approvalEntity.MinistryID.String
		}

		err = s.roleService.AssignRoleTx(ctx, tx, approvalEntity.RequesterID, approvalEntity.RequestedRole, ministryID)
		if err != nil {
			return err
		}
	}

	// 3. Commit if all succeeds
	return tx.Commit()
}
