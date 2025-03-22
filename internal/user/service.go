package user

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mappers"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	RegisterNewUser(ctx context.Context, user *domain.User) (*int, error)
	GenerateToken(user *entity.User) (string, error)
	UpdateUserDetails(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	logger    zap.SugaredLogger
	userRepo  userStorage.UserRepository
	jwtSecret string
}

func NewService(
	logger zap.SugaredLogger,
	userRepo userStorage.UserRepository,
	jwtSecret string,
) UserService {
	return &userService{
		logger:    logger,
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *userService) RegisterNewUser(ctx context.Context, user *domain.User) (*int, error) {
	s.logger.Infow("Registering new user")

	userEntity := mappers.ToUserEntity(*user)

	userID, err := s.userRepo.InsertUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

// GetUserByEmail retrieves a user by their email.
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	s.logger.Infow("Getting user by email")

	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GenerateToken generates a JWT for the authenticated user.
func (s *userService) GenerateToken(user *entity.User) (string, error) {
	s.logger.Infow("Generating token")
	// Define claims; you can add custom claims as needed.
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (s *userService) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	s.logger.Infow("Getting user by id")

	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserDetails updates a user in the database.
func (s *userService) UpdateUserDetails(ctx context.Context, user *entity.User) (*entity.User, error) {
	s.logger.Infow("Updating user details")

	user, err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by their ID.
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	s.logger.Infow("Deleting user")

	if err := s.userRepo.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("service error deleting user: %w", err)
	}

	return nil
}
