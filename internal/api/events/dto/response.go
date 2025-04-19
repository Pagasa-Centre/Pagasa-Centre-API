package dto

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/domain"
)

type EventDay struct {
	Date               string    `json:"date"`
	StartTime          time.Time `json:"start_time"`
	EndTime            time.Time `json:"end_time"`
	Weekday            string    `json:"weekday"`
	WeekDayShortFormat string    `json:"week_day_short_format"`
}

type Event struct {
	Title                 string     `json:"title"`
	Description           string     `json:"description"`
	AdditionalInformation string     `json:"additional_information"`
	Location              string     `json:"location"`
	RegistrationLink      string     `json:"registration_link"`
	Days                  []EventDay `json:"days"`
}

type GetAllEventsResponse struct {
	Message string   `json:"message"`
	Events  *[]Event `json:"events"`
}

func ToGetAllEventsResponse(events []domain.Events, message string) GetAllEventsResponse {
	var dtoEvents []Event

	for _, e := range events {
		var dtoDays []EventDay
		for _, d := range e.Days {
			dtoDays = append(dtoDays, EventDay{
				Date:               d.Date,
				StartTime:          d.StartTime,
				EndTime:            d.EndTime,
				Weekday:            d.WeekDay,
				WeekDayShortFormat: d.WeekDayShortFormat,
			})
		}

		dtoEvents = append(dtoEvents, Event{
			Title:                 e.Title,
			Description:           e.Description,
			AdditionalInformation: e.AdditionalInformation,
			Location:              e.Location,
			RegistrationLink:      e.RegistrationLink,
			Days:                  dtoDays,
		})
	}

	return GetAllEventsResponse{
		Message: message,
		Events:  &dtoEvents,
	}
}

type CreateEventsResponse struct {
	Message string `json:"message"`
}

func ToCreateEventResponse(message string) CreateEventsResponse {
	return CreateEventsResponse{
		Message: message,
	}
}
