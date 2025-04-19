package mappers

import (
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/utils"
)

func EntityToEventsDomain(eventEntities *entity.EventSlice, eventDays *entity.EventDaySlice) *[]domain.Events {
	if eventEntities == nil {
		return nil
	}

	// Map of eventID -> []EventDays
	daysByEventID := make(map[string][]domain.EventDays)

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

		weekDay, weekDayShortform, err := utils.GetWeekdayFromDate(day.Date.Format("2006-01-02"))
		if err != nil {
			return nil
		}

		daysByEventID[day.EventID] = append(daysByEventID[day.EventID], domain.EventDays{
			Date:               day.Date.Format("2006-01-02"),
			WeekDay:            weekDay,
			WeekDayShortFormat: weekDayShortform,
			StartTime:          start,
			EndTime:            end,
		})
	}

	var results []domain.Events

	for _, e := range *eventEntities {
		results = append(results, domain.Events{
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
