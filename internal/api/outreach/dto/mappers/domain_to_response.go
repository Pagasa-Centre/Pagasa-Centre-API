package mappers

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
)

func ToGetAllOutreachesResponse(outreaches []*domain.Outreach, services []*domain.Service, message string) dto.GetAllOutreachesResponse {
	// Map outreachID -> []Service
	serviceMap := make(map[string][]dto.Service)

	for _, svc := range services {
		serviceMap[svc.OutreachID] = append(serviceMap[svc.OutreachID], dto.Service{
			StartTime: formatTime(svc.StartTime),
			EndTime:   formatTime(svc.EndTime),
			Day:       svc.Day,
		})
	}

	var outreachesResp dto.GetAllOutreachesResponse
	outreachesResp.Message = message

	for _, outreach := range outreaches {
		outreachDto := ToResponse(outreach)

		// Attach services if present
		if svc, ok := serviceMap[outreach.ID]; ok {
			outreachDto.Services = svc
		}

		outreachesResp.Outreaches = append(outreachesResp.Outreaches, outreachDto)
	}

	return outreachesResp
}

func ToResponse(outreach *domain.Outreach) *dto.Outreach {
	if outreach == nil {
		return nil
	}
	return &dto.Outreach{
		ID:           outreach.ID,
		Name:         outreach.Name,
		AddressLine1: outreach.AddressLine1,
		AddressLine2: outreach.AddressLine2,
		Postcode:     outreach.Postcode,
		City:         outreach.City,
		Country:      outreach.Country,
		VenueName:    outreach.VenueName,
		Region:       outreach.Region,
		ThumbnailURL: outreach.ThumbnailURL,
	}
}

func ToErrorOutreachesResponse(message string) dto.GetAllOutreachesResponse {
	return dto.GetAllOutreachesResponse{
		Message: message,
	}
}

func formatTime(t time.Time) string {
	return t.Format("15:04")
}
