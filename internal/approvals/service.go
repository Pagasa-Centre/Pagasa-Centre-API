package approvals

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	roleDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/contracts"
	userDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	context2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
)

type ApprovalService interface {
	CreateNewApproval(ctx context.Context, Approval *domain.Approval) error
	GetAll(ctx context.Context) (*GetAllResult, error)
	SetUserService(us contracts.UserService)
	UpdateApprovalStatus(ctx context.Context, approvalID, status string) error
}

type service struct {
	logger       *zap.Logger
	approvalRepo storage.ApprovalRepository
	userService  contracts.UserService
	roleService  roles.RolesService
}

func NewApprovalService(
	logger *zap.Logger,
	approvalRepo storage.ApprovalRepository,
	userService contracts.UserService,
	roleService roles.RolesService,
) ApprovalService {
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
		Users     []*userDomain.User
	}
)

var ErrNoPermission = errors.New("you do not have permission to update this approval")

func (s *service) CreateNewApproval(ctx context.Context, approval *domain.Approval) error {
	s.logger.Info("Creating new approval")

	approvalEntity := mappers.ToApprovalEntity(approval)

	err := s.approvalRepo.Insert(ctx, approvalEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context) (*GetAllResult, error) {
	// Get the user ID from the context
	userID, err := context2.GetUserIDString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	_, err = s.userService.GetUserById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
	}

	// Get user roles
	userRoles, err := s.roleService.GetUserRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	var approvalsSlice entity.ApprovalSlice

	// Role-based logic
	switch {
	case slices.Contains(userRoles, roleDomain.RoleAdmin):
		approvalsSlice, err = s.approvalRepo.GetAllPendingApprovals(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get approvals: %w", err)
		}
	default:
		// Map of leader role => member role
		roleMap := map[string]string{
			roleDomain.RoleChildrensMinistryLeader:        roleDomain.RoleChildrensMinistryMember,
			roleDomain.RoleCreativeArtsMinistryLeader:     roleDomain.RoleCreativeArtsMinistryMember,
			roleDomain.RoleMediaMinistryLeader:            roleDomain.RoleMediaMinistryMember,
			roleDomain.RoleMusicMinistryLeader:            roleDomain.RoleMusicMinistryMember,
			roleDomain.RoleProductionMinistryLeader:       roleDomain.RoleProductionMinistryMember,
			roleDomain.RoleUsheringSecurityMinistryLeader: roleDomain.RoleUsheringSecurityMinistryMember,
			roleDomain.RolePastor:                         roleDomain.RolePrimary,
		}
		// get all approvals a user is able to approve
		for leaderRole, memberRole := range roleMap {
			if slices.Contains(userRoles, leaderRole) {
				roleApprovals, err := s.approvalRepo.GetAllPendingApprovalsByRequestedRole(ctx, memberRole)
				if err != nil {
					return nil, fmt.Errorf("failed to get approvals for %s: %w", memberRole, err)
				}

				approvalsSlice = append(approvalsSlice, roleApprovals...)
			}
		}
	}

	approvalsDomain := mappers.EntityToDomainApprovals(approvalsSlice)

	// 2.Get All the details of the requester on each approval
	var users []*userDomain.User

	for _, apr := range approvalsSlice {
		u, err := s.userService.GetUserById(ctx, apr.RequesterID)
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

func (s *service) SetUserService(us contracts.UserService) {
	s.userService = us
}

func (s *service) UpdateApprovalStatus(ctx context.Context, approvalID, status string) error {
	s.logger.Info("Updating approval status")
	// Get the user ID from the context
	userID, err := context2.GetUserIDString(ctx)
	if err != nil {
		return fmt.Errorf("failed to get user id from context: %w", err)
	}

	// check if user exists in db.
	_, err = s.userService.GetUserById(ctx, userID)
	if err != nil {
		return fmt.Errorf("user does not exist or has been deleted: %w", err)
	}

	// Get user roles
	userRoles, err := s.roleService.GetUserRoles(ctx, userID)
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
	approval, err := s.approvalRepo.GetApprovalByIDTx(ctx, tx, approvalID)
	if err != nil {
		return err
	}

	switch status {
	case domain.Approved:
		approval.Status = domain.Approved
	case domain.Rejected:
		approval.Status = domain.Rejected
	default:
		return fmt.Errorf("invalid approval status: %s", status)
	}

	approval.UpdatedBy = null.StringFrom(userID)

	// 2. Update status in transaction
	err = s.approvalRepo.UpdateTx(ctx, tx, approval)
	if err != nil {
		return err
	}

	if approval.Status == domain.Approved {
		var ministryID *string
		if approval.MinistryID.Valid {
			ministryID = &approval.MinistryID.String
		}

		err = s.roleService.AssignRoleTx(ctx, tx, approval.RequesterID, approval.RequestedRole, ministryID)
		if err != nil {
			return err
		}
	}

	// 3. Commit if all succeeds
	return tx.Commit()
}
