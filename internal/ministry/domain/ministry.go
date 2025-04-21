package domain

import (
	"time"
)

type Ministry struct {
	ID               string
	OutreachID       string
	Name             string
	Day              string
	StartTime        *time.Time
	EndTime          *time.Time
	MeetingLocation  string
	ShortDescription string
	LongDescription  string
	ThumbnailURL     string
	MinistryLeaders  []string
}
