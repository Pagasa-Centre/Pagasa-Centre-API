package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role"
	roleDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	userDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

type Service interface {
	Register(ctx context.Context, registrationDetails *domain.Registration) (*domain.AuthResult, error)
	Login(ctx context.Context, loginInput domain.Login) (*domain.AuthResult, error)
}

type authService struct {
	logger         *zap.Logger
	jwtSecret      string
	userService    user.Service
	authRepository storage.AuthRepository
	rolesService   role.RolesService
}

func NewAuthService(
	logger *zap.Logger,
	jwtSecret string,
	authRepository storage.AuthRepository,
	userService user.Service,
	rolesService role.RolesService,
) Service {
	return &authService{
		logger:         logger,
		jwtSecret:      jwtSecret,
		userService:    userService,
		authRepository: authRepository,
		rolesService:   rolesService,
	}
}

func (s *authService) Register(ctx context.Context, registrationDetails *domain.Registration) (*domain.AuthResult, error) {
	domainUser, err := s.userService.Create(ctx, userDomain.User{
		FirstName:      registrationDetails.FirstName,
		LastName:       registrationDetails.LastName,
		HashedPassword: registrationDetails.Password,
		Email:          registrationDetails.Email,
		PhoneNumber:    registrationDetails.PhoneNumber,
		Birthday:       registrationDetails.Birthday,
		CellLeaderID:   registrationDetails.CellLeaderID,
		OutreachID:     registrationDetails.OutreachID,
		MinistryID:     registrationDetails.MinistryID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	err = s.rolesService.HandleRoleApprovals(ctx, roleDomain.RoleApplication{
		UserID:           domainUser.ID,
		IsLeader:         registrationDetails.IsLeader,
		IsPrimary:        registrationDetails.IsPrimary,
		IsPastor:         registrationDetails.IsPastor,
		IsMinistryLeader: registrationDetails.IsMinistryLeader,
		MinistryID:       registrationDetails.MinistryID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to handle role approvals: %w", err)
	}

	token, err := s.generateToken(domainUser)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	userRoles, err := s.rolesService.Fetch(ctx, domainUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	return &domain.AuthResult{
		AccessToken: token,
		User:        domainUser,
		Roles:       userRoles,
	}, nil
}

func (s *authService) Login(ctx context.Context, loginDetails domain.Login) (*domain.AuthResult, error) {
	domainUser, err := s.userService.AuthenticateUser(ctx, loginDetails.Email, loginDetails.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %w", err)
	}

	token, err := s.generateToken(domainUser)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	userRoles, err := s.rolesService.Fetch(ctx, domainUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	return &domain.AuthResult{
		AccessToken: token,
		User:        domainUser,
		Roles:       userRoles,
	}, nil
}

// generateToken generates a JWT for the authenticated user.
func (s *authService) generateToken(user *userDomain.User) (string, error) {
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
