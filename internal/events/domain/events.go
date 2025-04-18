package domain

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type Events struct {
	Title                 string
	Description           string
	AdditionalInformation string
	Location              string
	RegistrationLink      string
	Days                  []EventDays
}

type EventDays struct {
	Date      string
	StartTime time.Time
	EndTime   time.Time
}

func EntityToEventsDomain(eventEntities *entity.EventSlice, eventDays *entity.EventDaySlice) *[]Events {
	if eventEntities == nil {
		return nil
	}

	// Map of eventID -> []EventDays
	daysByEventID := make(map[string][]EventDays)

	for _, day := range *eventDays {
		if day.EventID == "" {
			continue
		}

		// Convert null.Time to time.Time safely (zero if invalid)
		start := time.Time{}
		end := time.Time{}

		if day.StartTime.Valid {
			start = day.StartTime.Time
		}

		if day.EndTime.Valid {
			end = day.EndTime.Time
		}

		daysByEventID[day.EventID] = append(daysByEventID[day.EventID], EventDays{
			Date:      day.Date.Format("2006-01-02"),
			StartTime: start,
			EndTime:   end,
		})
	}

	var results []Events

	for _, e := range *eventEntities {
		results = append(results, Events{
			Title:                 e.Title,
			Description:           e.Description.String,
			AdditionalInformation: e.AdditionalInformation.String,
			Location:              e.Location.String,
			RegistrationLink:      e.RegistrationLink.String,
			Days:                  daysByEventID[e.ID], // may be nil and that's okay
		})
	}

	return &results
}
