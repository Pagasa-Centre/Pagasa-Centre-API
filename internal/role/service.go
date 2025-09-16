package role

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval"
	approvaldomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/domain"
	ministrydomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/interfaces"
)

type RolesService interface {
	Fetch(ctx context.Context, userID string) ([]string, error)
	HandleRoleApprovals(ctx context.Context, app domain.RoleApplication) error
	AssignRole(ctx context.Context, userID, role string) error // todo: update all to take domain role instead of string
	AssignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, role string, ministryID *string) error
	SetApprovalService(service approval.Service)
	SetMinistryService(service interfaces.MinistryService)
}

type service struct {
	logger           *zap.Logger
	repository       storage.RolesRepository
	approvalsService approval.Service
	ministryService  interfaces.MinistryService
}

func NewRoleService(
	logger *zap.Logger,
	repository storage.RolesRepository,
	approvalsService approval.Service,
	ministryService interfaces.MinistryService,
) RolesService {
	return &service{
		logger:           logger,
		repository:       repository,
		approvalsService: approvalsService,
		ministryService:  ministryService,
	}
}

func (s *service) SetApprovalService(service approval.Service) {
	s.approvalsService = service
}

func (s *service) SetMinistryService(service interfaces.MinistryService) {
	s.ministryService = service
}

func (s *service) HandleRoleApprovals(ctx context.Context, app domain.RoleApplication) error {
	s.logger.Sugar().Infof("Handling role application for %s", app.UserID)
	//todo: remove logs like this and add context to errors instead with helpful non sensityve data.

	createApproval := func(requestedRole, approvalType string, ministryID *string) error {
		return s.approvalsService.CreateNewApproval(ctx, &approvaldomain.Approval{
			RequesterID:   app.UserID,
			RequestedRole: requestedRole,
			Type:          approvalType,
			Status:        string(approvaldomain.Pending),
			MinistryID:    ministryID,
		})
	}

	var err error
	if app.IsLeader {
		err = createApproval(
			string(domain.RoleLeader),
			string(approvaldomain.LeaderStatusConfirmation),
			nil,
		)
		if err != nil {
			return err
		}
	}

	if app.IsPrimary {
		err = createApproval(
			string(domain.RolePrimary),
			string(approvaldomain.PrimaryStatusConfirmation),
			nil,
		)
		if err != nil {
			return err
		}
	}

	if app.IsPastor {
		err = createApproval(
			string(domain.RolePastor),
			string(approvaldomain.PastorStatusConfirmation),
			nil,
		)
		if err != nil {
			return err
		}
	}

	if app.IsMinistryLeader {
		if app.MinistryID == nil {
			return fmt.Errorf("ministry_id is required for ministry leader role")
		}

		ministry, err := s.ministryService.GetByID(ctx, *app.MinistryID)
		if err != nil {
			return fmt.Errorf("failed to get ministry by id: %w", err)
		}

		requestedLeaderRole, err := ministrydomain.ToMinistryLeaderRole(ministry.Name) // TODO: MAYBE Move to mapper
		if err != nil {
			return fmt.Errorf("failed to convert ministry name to leader role: %w", err)
		}

		err = createApproval(
			requestedLeaderRole,
			string(approvaldomain.MinistryLeaderStatusConfirmation),
			app.MinistryID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) AssignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, role string, ministryID *string) error {
	s.logger.Sugar().Infof("Assigning %s role", role)

	err := s.repository.AssignRoleTx(ctx, tx, userID, role, ministryID)
	if err != nil {
		s.logger.Sugar().With("error", err).Errorf("failed to assign %s role", role)

		return fmt.Errorf("failed to assign %s role: %w", role, err)
	}

	return nil
}

func (s *service) AssignRole(ctx context.Context, userID, role string) error {
	s.logger.Sugar().Infof("Assigning %s role", role)

	err := s.repository.AssignRole(ctx, userID, role)
	if err != nil {
		s.logger.Sugar().With("error", err).Errorf("failed to assign %s role", role)

		return fmt.Errorf("failed to assign %s role: %w", role, err)
	}

	return nil
}

func (s *service) Fetch(ctx context.Context, userID string) ([]string, error) {
	roles, err := s.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	return roles, nil
}
