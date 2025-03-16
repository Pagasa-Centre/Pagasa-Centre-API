package authentication

import (
	"context"
	"go.uber.org/zap"
)

type Service interface {
	RegisterNewUser(ctx context.Context) error
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

func (s *service) RegisterNewUser(ctx context.Context) error {
	s.logger.Infow("Registering new user", "context", ctx)
	return nil
}
