package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
)

type Ministry struct {
	ID              int     `json:"id"`
	OutreachID      int     `json:"outreach_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Day             string  `json:"day"`
	StartTime       *string `json:"start_time"`
	EndTime         *string `json:"end_time"`
	MeetingLocation string  `json:"meeting_location,omitempty"`
}

func ToResponse(ministry *domain.Ministry) *Ministry {
	var formattedStartTime *string

	if ministry.StartTime != nil {
		formatted := ministry.StartTime.Format("15:04")
		formattedStartTime = &formatted
	}

	var formattedEndTime *string

	if ministry.EndTime != nil {
		formatted := ministry.EndTime.Format("15:04")
		formattedEndTime = &formatted
	}

	return &Ministry{
		ID:              ministry.ID,
		OutreachID:      ministry.OutreachID,
		Name:            ministry.Name,
		Description:     ministry.Description,
		Day:             ministry.Day,
		StartTime:       formattedStartTime,
		EndTime:         formattedEndTime,
		MeetingLocation: ministry.MeetingLocation,
	}
}

type GetAllMinistriesResponse struct {
	Ministries []*Ministry `json:"ministries"`
}

func ToGetAllMinistriesResponse(ministries []*domain.Ministry) GetAllMinistriesResponse {
	var ministriesResp GetAllMinistriesResponse

	for _, ministry := range ministries {
		ministriesResp.Ministries = append(ministriesResp.Ministries, ToResponse(ministry))
	}

	return ministriesResp
}
