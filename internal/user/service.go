package user

import (
	"context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mappers"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	"go.uber.org/zap"
)

type Service interface {
	RegisterNewUser(ctx context.Context, user domain.User) error
}

type service struct {
	logger   zap.SugaredLogger
	userRepo userStorage.Repository
}

func NewService(
	logger zap.SugaredLogger,
	userRepo userStorage.Repository,
) Service {
	return &service{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *service) RegisterNewUser(ctx context.Context, user domain.User) error {
	s.logger.Infow("Registering new user", "context", ctx)

	userEntity := mappers.ToUserEntity(user)

	err := s.userRepo.InsertUser(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}
