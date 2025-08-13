package mapper

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/event/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/domain"
)

func CreateEventRequestToDomain(req dto.CreateEventRequest) domain.Events {
	var days []domain.EventDays
	for _, d := range req.Days {
		days = append(days, domain.EventDays{
			Date:      d.Date,
			StartTime: d.StartTime,
			EndTime:   d.EndTime,
		})
	}

	return domain.Events{
		Title:                 req.Title,
		Description:           req.Description,
		AdditionalInformation: req.AdditionalInformation,
		Location:              req.Location,
		RegistrationLink:      req.RegistrationLink,
		Days:                  days,
	}
}
