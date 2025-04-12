package user

import (
	"context"
	"fmt"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
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
	RegisterNewUser(ctx context.Context, user *domain.User, req dto.RegisterRequest) (*entity.User, error)
	GenerateToken(user *entity.User) (string, error)
	UpdateUserDetails(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	logger          *zap.Logger
	userRepo        userStorage.UserRepository
	jwtSecret       string
	rolesService    roles.RolesService
	ministryService ministry.MinistryService
}

func NewUserService(
	logger *zap.Logger,
	userRepo userStorage.UserRepository,
	jwtSecret string,
	rolesService roles.RolesService,
	ministryService ministry.MinistryService,
) UserService {
	return &userService{
		logger:          logger,
		userRepo:        userRepo,
		jwtSecret:       jwtSecret,
		rolesService:    rolesService,
		ministryService: ministryService,
	}
}

// RegisterNewUser inserts a new user into the user table and applies any roles that were provided by the user
func (s *userService) RegisterNewUser(ctx context.Context, user *domain.User, req dto.RegisterRequest) (*entity.User, error) {
	s.logger.Info("Registering new user")

	userEntity := mappers.ToUserEntity(*user)

	userID, err := s.userRepo.InsertUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	if req.IsLeader {
		err = s.rolesService.AssignLeaderRole(ctx, *userID)
		if err != nil {
			return nil, fmt.Errorf("failed to assign leader role: %w", err)
		}
	}

	if req.IsPrimary {
		err = s.rolesService.AssignPrimaryRole(ctx, *userID)
		if err != nil {
			return nil, fmt.Errorf("failed to assign primary role: %w", err)
		}
	}

	if req.IsPastor {
		err = s.rolesService.AssignPastorRole(ctx, *userID)
		if err != nil {
			return nil, fmt.Errorf("failed to assign pastor role: %w", err)
		}
	}

	if req.IsMinistryLeader {
		err = s.rolesService.AssignMinistryLeaderRole(ctx, *userID)
		if err != nil {
			return nil, fmt.Errorf("failed to assign ministry leader role: %w", err)
		}

		err = s.ministryService.AssignLeaderToMinistry(ctx, *req.MinistryID, *userID)
		if err != nil {
			return nil, fmt.Errorf("failed to assign leader to ministry: %w", err)
		}
	}

	u, err := s.GetUserById(ctx, *userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return u, nil
}

// GetUserByEmail retrieves a user by their email.
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	s.logger.Info("Getting user by email")

	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GenerateToken generates a JWT for the authenticated user.
func (s *userService) GenerateToken(user *entity.User) (string, error) {
	s.logger.Info("Generating token")
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

func (s *userService) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	s.logger.Info("Getting user by id")

	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserDetails updates a user in the database.
func (s *userService) UpdateUserDetails(ctx context.Context, user *entity.User) (*entity.User, error) {
	s.logger.Info("Updating user details")

	user, err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by their ID.
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	s.logger.Info("Deleting user")

	if err := s.userRepo.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("service error deleting user: %w", err)
	}

	return nil
}
