package roles

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/storage"
)

type RolesService interface {
	AssignLeaderRole(ctx context.Context, userID int) error
	AssignPrimaryRole(ctx context.Context, userID int) error
	AssignPastorRole(ctx context.Context, userID int) error
	AssignMinistryLeaderRole(ctx context.Context, userID int) error
}

type service struct {
	logger     zap.SugaredLogger
	repository storage.RolesRepository
}

func NewRoleService(logger zap.SugaredLogger, repository storage.RolesRepository) RolesService {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) AssignMinistryLeaderRole(ctx context.Context, userID int) error {
	s.logger.Info("Assigning Ministry Leader role")

	err := s.repository.AssignPastorRole(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to assign Ministry Leader role", zap.Error(err))

		return fmt.Errorf("failed to assign Ministry Leader role: %w", err)
	}

	return nil
}

func (s *service) AssignLeaderRole(ctx context.Context, userID int) error {
	s.logger.Info("Assigning Leader role")

	err := s.repository.AssignLeaderRole(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to assign leader role", zap.Error(err))

		return fmt.Errorf("failed to assign leader role: %w", err)
	}

	return nil
}

func (s *service) AssignPrimaryRole(ctx context.Context, userID int) error {
	s.logger.Info("Assigning Primary role")

	err := s.repository.AssignPrimaryRole(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to assign primary role", zap.Error(err))

		return fmt.Errorf("failed to assign primary role: %w", err)
	}

	return nil
}

func (s *service) AssignPastorRole(ctx context.Context, userID int) error {
	s.logger.Info("Assigning Pastor role")

	err := s.repository.AssignPastorRole(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to assign Pastor role", zap.Error(err))

		return fmt.Errorf("failed to assign Pastor role: %w", err)
	}

	return nil
}
