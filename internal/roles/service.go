package roles

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/storage"
)

type RolesService interface {
	AssignRole(ctx context.Context, userID, role string) error
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

func (s *service) AssignRole(ctx context.Context, userID, role string) error {
	s.logger.Sugar().Infof("Assigning %s role", role)

	err := s.repository.AssignRole(ctx, userID, role)
	if err != nil {
		s.logger.Sugar().With("error", err).Errorf("failed to assign %s role", role)

		return fmt.Errorf("failed to assign %s role: %w", role, err)
	}

	return nil
}
