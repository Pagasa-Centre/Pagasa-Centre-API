package mapper

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/domain"
)

func MapRegisterRequestToDomain(req dto.RegisterRequest) (*domain.Registration, error) {
	parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if req.CellLeaderID != nil && *req.CellLeaderID == "" { // todo: ?
		req.CellLeaderID = nil
	}

	if req.MinistryID != nil && *req.MinistryID == "" { // todo: ?
		req.MinistryID = nil
	}

	return &domain.Registration{
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Password:         string(hashedPassword),
		Email:            req.Email,
		PhoneNumber:      req.PhoneNumber,
		Birthday:         parsedBirthday,
		CellLeaderID:     req.CellLeaderID,
		OutreachID:       req.OutreachID,
		MinistryID:       req.MinistryID,
		IsMinistryLeader: req.IsMinistryLeader,
		IsLeader:         req.IsLeader,
		IsPrimary:        req.IsPrimary,
	}, nil
}

func MapLoginRequestToDomain(req dto.LoginRequest) domain.Login {
	return domain.Login{
		Email:    req.Email,
		Password: req.Password,
	}
}
