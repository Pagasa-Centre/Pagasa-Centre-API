package roles

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/storage"
)

type RolesService interface {
	AssignRole(ctx context.Context, userID, role string) error
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
	AssignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, role string, ministryID *string) error
}

type service struct {
	logger     *zap.Logger
	repository storage.RolesRepository
}

func NewRoleService(logger *zap.Logger, repository storage.RolesRepository) RolesService {
	return &service{
		logger:     logger,
		repository: repository,
	}
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

func (s *service) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	roles, err := s.repository.GetUserRoles(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	return roles, nil
}
