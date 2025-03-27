package domain

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type Ministry struct {
	ID              int
	OutreachID      int
	Name            string
	Description     string
	Day             string
	StartTime       *time.Time
	EndTime         *time.Time
	MeetingLocation string
}

func ToDomain(ministry *entity.Ministry) *Ministry {
	var startTime *time.Time
	if ministry.StartTime.Valid {
		startTime = &ministry.StartTime.Time
	}

	var endTime *time.Time
	if ministry.EndTime.Valid {
		endTime = &ministry.EndTime.Time
	}

	description := ""
	if ministry.Description.Valid {
		description = ministry.Description.String
	}

	return &Ministry{
		ID:              ministry.ID,
		OutreachID:      ministry.OutreachID,
		Name:            ministry.Name,
		Description:     description,
		Day:             ministry.MeetingDay.String,
		StartTime:       startTime,
		EndTime:         endTime,
		MeetingLocation: ministry.MeetingLocation.String,
	}
}
