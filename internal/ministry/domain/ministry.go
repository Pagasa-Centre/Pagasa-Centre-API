package domain

import (
	"time"
)

type Ministry struct {
	ID              string
	OutreachID      string
	Name            string
	Description     string
	Day             string
	StartTime       *time.Time
	EndTime         *time.Time
	MeetingLocation string
}
