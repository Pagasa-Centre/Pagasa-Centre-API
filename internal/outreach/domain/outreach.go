package domain

import (
	"time"
)

type Outreach struct {
	ID           string
	Name         string
	AddressLine1 string
	AddressLine2 string
	Postcode     string
	City         string
	Country      string
	VenueName    string
	Region       string
	ThumbnailURL string
}

type Service struct {
	OutreachID string
	StartTime  time.Time
	EndTime    time.Time
	Day        string
}
