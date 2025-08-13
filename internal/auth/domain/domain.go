package domain

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

type Registration struct {
	FirstName        string
	LastName         string
	Email            string
	Password         string
	Birthday         time.Time
	OutreachID       string
	PhoneNumber      string
	CellLeaderID     *string
	MinistryID       *string
	IsLeader         bool
	IsPrimary        bool
	IsPastor         bool
	IsMinistryLeader bool
}

type Login struct {
	Email    string
	Password string
}

type AuthResult struct {
	// RefreshToken string todo: implement later
	AccessToken string
	User        *domain.User
	Roles       []string
}
