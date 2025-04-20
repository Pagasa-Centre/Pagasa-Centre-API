package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
)

func ToGetAllMinistriesResponse(ministries []*domain.Ministry, message string) dto.GetAllMinistriesResponse {
	if ministries == nil {
		return dto.GetAllMinistriesResponse{
			Message: message,
		}
	}
	var ministriesResp dto.GetAllMinistriesResponse

	for _, ministry := range ministries {
		ministriesResp.Ministries = append(ministriesResp.Ministries, toMinistries(ministry))
	}

	ministriesResp.Message = message

	return ministriesResp
}

func toMinistries(ministry *domain.Ministry) *dto.Ministry {
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

	return &dto.Ministry{
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

func ToMinistryApplicationResponse(message string) dto.MinistryApplicationResponse {
	return dto.MinistryApplicationResponse{
		Message: message,
	}
}
