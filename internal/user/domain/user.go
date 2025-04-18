package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
)

type User struct {
	FirstName      string
	LastName       string
	HashedPassword string
	Email          string
	PhoneNumber    string
	Birthday       time.Time
	CellLeaderID   *string
	OutreachID     string
	MinistryID     *string
}

func RegisterRequestToUserDomain(req dto.RegisterRequest) (user *User, err error) {
	parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if req.CellLeaderID != nil && *req.CellLeaderID == "" {
		req.CellLeaderID = nil
	}

	var ministryID *string
	if req.MinistryID != nil && *req.MinistryID == "" {
		ministryID = nil
	}

	if req.MinistryID != nil {
		ministryID = req.MinistryID
	}

	return &User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		HashedPassword: string(hashedPassword),
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Birthday:       parsedBirthday,
		CellLeaderID:   req.CellLeaderID,
		OutreachID:     req.OutreachID,
		MinistryID:     ministryID,
	}, nil
}
