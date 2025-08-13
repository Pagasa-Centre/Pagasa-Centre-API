package domain

import "time"

type Events struct {
	Title                 string
	Description           string
	AdditionalInformation string
	Location              string
	RegistrationLink      string
	Days                  []EventDays
}

type EventDays struct {
	Date               string
	WeekDay            string
	WeekDayShortFormat string
	StartTime          time.Time
	EndTime            time.Time
}
