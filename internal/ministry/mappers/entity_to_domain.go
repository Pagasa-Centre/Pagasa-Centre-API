package mappers

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
)

func ToDomain(ministry *entity.Ministry, leaderNames []string) *domain.Ministry {
	var endTime *time.Time
	if ministry.EndTime.Valid {
		endTime = &ministry.EndTime.Time
	}

	var shortDescription string
	if ministry.ShortDescription.Valid {
		shortDescription = ministry.ShortDescription.String
	}

	var longDescription string
	if ministry.LongDescription.Valid {
		longDescription = ministry.LongDescription.String
	}

	var thumbnailURL string
	if ministry.ThumbnailURL.Valid {
		thumbnailURL = ministry.ThumbnailURL.String
	}

	return &domain.Ministry{
		ID:               ministry.ID,
		OutreachID:       ministry.OutreachID,
		Name:             ministry.Name,
		Day:              ministry.MeetingDay.String,
		StartTime:        &ministry.StartTime,
		EndTime:          endTime,
		MeetingLocation:  ministry.MeetingLocation.String,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
		ThumbnailURL:     thumbnailURL,
		MinistryLeaders:  leaderNames,
	}
}
