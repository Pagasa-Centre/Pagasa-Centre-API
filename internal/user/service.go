package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mapper"
	userstorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	commoncontext "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/utils"
)

type Service interface {
	Create(ctx context.Context, user domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, req dto.UpdateDetailsRequest) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	AuthenticateUser(ctx context.Context, email, password string) (*domain.User, error)
}

type userService struct {
	logger    *zap.Logger
	userRepo  userstorage.UserRepository
	jwtSecret string
}

func NewUserService(
	logger *zap.Logger,
	userRepo userstorage.UserRepository,
	jwtSecret string,
) Service {
	return &userService{
		logger:    logger,
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

var (
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidLoginDetails = errors.New("invalid email or password")
	ErrInvalidOutreach     = errors.New("invalid outreach")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)

// AuthenticateUser checks credentials and returns the user if valid, otherwise an error.
func (s *userService) AuthenticateUser(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	userDomain := mapper.EntityToDomain(user)

	return userDomain, nil
}

// Create inserts a new user into the user table and applies any roles that were provided by the user.
func (s *userService) Create(ctx context.Context, domainUser domain.User) (*domain.User, error) {
	domainUser.PhoneNumber = utils.NormalizePhoneNumber(domainUser.PhoneNumber)

	userEntity := mapper.ToEntity(domainUser)

	userID, err := s.userRepo.Insert(ctx, userEntity)
	if err != nil {
		// Check if the error is a unique constraint violation (email already exists)
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			// PostgreSQL error code 23505 = unique_violation
			return nil, ErrEmailAlreadyExists
		} else if errors.As(err, &pqErr) && pqErr.Code == "23503" && pqErr.Constraint == "users_outreach_id_fkey" {
			return nil, ErrInvalidOutreach
		}

		return nil, err
	}

	userEntity, err = s.userRepo.GetById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return mapper.EntityToDomain(userEntity), nil
}

// GetByEmail retrieves a user by their email.
func (s *userService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	userEntity, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	userDomain := mapper.EntityToDomain(userEntity)

	return userDomain, nil
}

func (s *userService) GetById(ctx context.Context, id string) (*domain.User, error) {
	s.logger.With(zap.String("user id", id)).Info("Getting user by id")

	userEntity, err := s.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.EntityToDomain(userEntity), nil
}

// Update updates a user in the database.
// todo: fix this. should not be passing in request. should be sending domain object
func (s *userService) Update(ctx context.Context, req dto.UpdateDetailsRequest) (*domain.User, error) {
	// Get the user ID from the context
	userID, err := commoncontext.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	// Retrieve the current user from the database.
	currentUser, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	// Update user fields based on the request.
	s.updateUserFields(currentUser, req)

	user, err := s.userRepo.Update(ctx, currentUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return mapper.EntityToDomain(user), nil
}

// Delete deletes a user by their ID.
func (s *userService) Delete(ctx context.Context, userID string) error {
	return s.userRepo.Delete(ctx, userID)
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
		user.Phone = null.StringFrom(utils.NormalizePhoneNumber(req.PhoneNumber))
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
