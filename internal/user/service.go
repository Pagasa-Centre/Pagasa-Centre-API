package user

import (
	"context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mappers"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	"go.uber.org/zap"
)

type UserService interface {
	RegisterNewUser(ctx context.Context, user *domain.User) error
}

type userservice struct {
	logger   zap.SugaredLogger
	userRepo userStorage.UserRepository
}

func NewService(
	logger zap.SugaredLogger,
	userRepo userStorage.UserRepository,
) UserService {
	return &userservice{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *userservice) RegisterNewUser(ctx context.Context, user *domain.User) error {
	s.logger.Infow("Registering new user")

	userEntity := mappers.ToUserEntity(*user)

	err := s.userRepo.InsertUser(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}
