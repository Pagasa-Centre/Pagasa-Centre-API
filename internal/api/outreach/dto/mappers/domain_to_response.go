package mappers

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
)

func ToGetAllOutreachesResponse(result *outreach.GetAllOutreachesResult, message string) dto.GetAllOutreachesResponse {
	if result == nil {
		return dto.GetAllOutreachesResponse{
			Message: message,
		}
	}
	// Map outreachID -> []Service
	serviceMap := make(map[string][]dto.Service)

	for _, svc := range result.Services {
		serviceMap[svc.OutreachID] = append(serviceMap[svc.OutreachID], dto.Service{
			StartTime: formatTime(svc.StartTime),
			EndTime:   formatTime(svc.EndTime),
			Day:       svc.Day,
		})
	}

	var outreachesResp dto.GetAllOutreachesResponse
	outreachesResp.Message = message

	for _, o := range result.Outreaches {
		outreachDto := toResponse(o)

		// Attach services if present
		if svc, ok := serviceMap[o.ID]; ok {
			outreachDto.Services = svc
		}

		outreachesResp.Outreaches = append(outreachesResp.Outreaches, outreachDto)
	}

	return outreachesResp
}

func toResponse(outreach *domain.Outreach) *dto.Outreach {
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

func formatTime(t time.Time) string {
	return t.Format("15:04")
}
