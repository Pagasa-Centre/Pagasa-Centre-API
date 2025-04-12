package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
)

type Ministry struct {
	ID              string  `json:"id"`
	OutreachID      string  `json:"outreach_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Day             string  `json:"day"`
	StartTime       *string `json:"start_time,omitempty"`
	EndTime         *string `json:"end_time,omitempty"`
	MeetingLocation string  `json:"meeting_location,omitempty"`
}

type MinistryApplicationResponse struct {
	message string
}

func ToMinistryApplicationResponse(message string) MinistryApplicationResponse {
	return MinistryApplicationResponse{
		message: message,
	}
}

func ToMinistries(ministry *domain.Ministry) *Ministry {
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

func ToErrorMinistriesResponse(message string) GetAllMinistriesResponse {
	return GetAllMinistriesResponse{
		Message: message,
	}
}

type GetAllMinistriesResponse struct {
	Ministries []*Ministry `json:"ministries,omitempty"`
	Message    string      `json:"message,"`
}

func ToGetAllMinistriesResponse(ministries []*domain.Ministry, message string) GetAllMinistriesResponse {
	var ministriesResp GetAllMinistriesResponse

	for _, ministry := range ministries {
		ministriesResp.Ministries = append(ministriesResp.Ministries, ToMinistries(ministry))
	}

	ministriesResp.Message = message

	return ministriesResp
}
