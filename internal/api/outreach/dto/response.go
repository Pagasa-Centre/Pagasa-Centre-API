package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
	"time"
)

type Outreach struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	Postcode     string    `json:"postcode"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	VenueName    string    `json:"venue_name"`
	Region       string    `json:"region"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Services     []Service `json:"services,omitempty"`
}

type Service struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Day       string    `json:"day"`
}

func ToResponse(outreach *domain.Outreach) *Outreach {
	return &Outreach{
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

type GetAllOutreachesResponse struct {
	Message    string      `json:"message"`
	Outreaches []*Outreach `json:"outreaches,omitempty"`
}

func ToGetAllOutreachesResponse(outreaches []*domain.Outreach, services []*domain.Service, message string) GetAllOutreachesResponse {
	// Map outreachID -> []Service
	serviceMap := make(map[string][]Service)

	for _, svc := range services {
		serviceMap[svc.OutreachID] = append(serviceMap[svc.OutreachID], Service{
			StartTime: svc.StartTime,
			EndTime:   svc.EndTime,
			Day:       svc.Day,
		})
	}

	var outreachesResp GetAllOutreachesResponse
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

func ToErrorOutreachesResponse(message string) GetAllOutreachesResponse {
	return GetAllOutreachesResponse{
		Message: message,
	}
}
