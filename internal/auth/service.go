package authentication

import (
	"context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/domain"
	"go.uber.org/zap"
)

type Service interface {
	RegisterNewUser(ctx context.Context, user domain.User) error
}

type service struct {
	logger zap.SugaredLogger
}

func NewService(
	logger zap.SugaredLogger,
) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) RegisterNewUser(ctx context.Context, user domain.User) error {
	s.logger.Infow("Registering new user", "context", ctx)
	return nil
}
