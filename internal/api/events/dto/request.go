package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateEventDayRequest struct {
	Date      string    `json:"date" validate:"required"`       // e.g., "2025-06-01"
	StartTime time.Time `json:"start_time" validate:"required"` // ISO format
	EndTime   time.Time `json:"end_time" validate:"required"`   // ISO format
}

type CreateEventRequest struct {
	Title                 string                  `json:"title" validate:"required"`
	Description           string                  `json:"description,omitempty"`
	AdditionalInformation string                  `json:"additional_information,omitempty"`
	Location              string                  `json:"location,omitempty"`
	RegistrationLink      string                  `json:"registration_link,omitempty"`
	Days                  []CreateEventDayRequest `json:"days" validate:"required,dive"`
}

func (cer CreateEventRequest) Validate() error {
	return validator.New().Struct(cer)
}
