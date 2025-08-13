package domain

import (
	"time"
)

type User struct {
	ID             string
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
