package dto

import (
	"time"
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

type CreateEventsResponse struct {
	Message string `json:"message"`
}
