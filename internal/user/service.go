package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mappers"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	context2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	RegisterNewUser(ctx context.Context, user *domain.User, req dto.RegisterRequest) (*entity.User, error)
	GenerateToken(user *entity.User) (string, error)
	UpdateUserDetails(ctx context.Context, req dto.UpdateDetailsRequest) (*entity.User, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error)
	AuthenticateAndGenerateToken(ctx context.Context, email, password string) (*AuthResult, error)
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

type AuthResult struct {
	User  *entity.User
	Token string
}

func (s *userService) AuthenticateAndGenerateToken(ctx context.Context, email, password string) (*AuthResult, error) {
	user, err := s.AuthenticateUser(ctx, email, password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &AuthResult{
		User:  user,
		Token: token,
	}, nil
}

// AuthenticateUser checks credentials and returns the user if valid, otherwise an error.
func (s *userService) AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// RegisterNewUser inserts a new user into the user table and applies any roles that were provided by the user.
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
func (s *userService) UpdateUserDetails(ctx context.Context, req dto.UpdateDetailsRequest) (*entity.User, error) {
	s.logger.Info("Updating user details")

	// Get the user ID from the context
	userID, err := context2.GetUserIDString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	// Retrieve the current user from the database.
	currentUser, err := s.GetUserById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	// Update user fields based on the request.
	s.updateUserFields(currentUser, req)

	user, err := s.userRepo.UpdateUser(ctx, currentUser)
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

// updateUserFields updates the provided user entity with the values from the update request.
func (s *userService) updateUserFields(user *entity.User, req dto.UpdateDetailsRequest) {
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.PhoneNumber != "" {
		user.Phone = null.StringFrom(req.PhoneNumber)
	}

	if req.Birthday != "" {
		parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			s.logger.Sugar().Errorw("failed to parse birthday", "error", err)
		} else {
			user.Birthday = null.TimeFrom(parsedBirthday)
		}
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			s.logger.Sugar().Errorw("failed to hash new password", "error", err)
		} else {
			user.HashedPassword = string(hashedPassword)
		}
	}

	if req.CellLeaderID != nil {
		user.CellLeaderID = null.StringFrom(*req.CellLeaderID)
	}

	if req.OutreachID != "" {
		user.OutreachID = null.StringFrom(req.OutreachID)
	}
}
