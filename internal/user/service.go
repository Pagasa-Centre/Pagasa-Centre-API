package user

import (
	"context"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mappers"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
)

type UserService interface {
	RegisterNewUser(ctx context.Context, user *domain.User) (*int, error)
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

func (s *userservice) RegisterNewUser(ctx context.Context, user *domain.User) (*int, error) {
	s.logger.Infow("Registering new user")

	userEntity := mappers.ToUserEntity(*user)

	userID, err := s.userRepo.InsertUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return userID, nil
}
