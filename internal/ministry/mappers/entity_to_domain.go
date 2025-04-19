package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
	"time"
)

func ToDomain(ministry *entity.Ministry) *domain.Ministry {
	var endTime *time.Time
	if ministry.EndTime.Valid {
		endTime = &ministry.EndTime.Time
	}

	description := ""
	if ministry.Description.Valid {
		description = ministry.Description.String
	}

	return &domain.Ministry{
		ID:              ministry.ID,
		OutreachID:      ministry.OutreachID,
		Name:            ministry.Name,
		Description:     description,
		Day:             ministry.MeetingDay.String,
		StartTime:       &ministry.StartTime,
		EndTime:         endTime,
		MeetingLocation: ministry.MeetingLocation.String,
	}
}
