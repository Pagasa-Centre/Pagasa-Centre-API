package domain

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	FirstName      string
	LastName       string
	HashedPassword string
	Email          string
	PhoneNumber    string
	Birthday       time.Time
	CellLeaderID   *int
	OutreachID     int
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

	// Only set CellLeaderID if it's non-zero, otherwise leave it nil.
	var cellLeaderID *int
	if req.CellLeaderID != 0 {
		cellLeaderID = &req.CellLeaderID
	}

	return &User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		HashedPassword: string(hashedPassword),
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Birthday:       parsedBirthday,
		CellLeaderID:   cellLeaderID,
		OutreachID:     req.OutreachID,
	}, nil
}
