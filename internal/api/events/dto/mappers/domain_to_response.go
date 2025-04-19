package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/events/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/domain"
)

func ToGetAllEventsResponse(events []domain.Events, message string) dto.GetAllEventsResponse {
	var dtoEvents []dto.Event

	for _, e := range events {
		var dtoDays []dto.EventDay
		for _, d := range e.Days {
			dtoDays = append(dtoDays, dto.EventDay{
				Date:               d.Date,
				StartTime:          d.StartTime,
				EndTime:            d.EndTime,
				Weekday:            d.WeekDay,
				WeekDayShortFormat: d.WeekDayShortFormat,
			})
		}

		dtoEvents = append(dtoEvents, dto.Event{
			Title:                 e.Title,
			Description:           e.Description,
			AdditionalInformation: e.AdditionalInformation,
			Location:              e.Location,
			RegistrationLink:      e.RegistrationLink,
			Days:                  dtoDays,
		})
	}

	return dto.GetAllEventsResponse{
		Message: message,
		Events:  &dtoEvents,
	}
}

func ToCreateEventResponse(message string) dto.CreateEventsResponse {
	return dto.CreateEventsResponse{
		Message: message,
	}
}
