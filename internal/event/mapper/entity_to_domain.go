package mapper

import (
	"fmt"
	"time"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/utils"
)

func EntityToEventsDomain(eventEntities *entity.EventSlice, eventDays *entity.EventDaySlice) ([]*domain.Events, error) {
	if eventEntities == nil {
		return []*domain.Events{}, nil
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

		weekDay, weekDayShortform, err := utils.GetWeekdayFromDate(day.Date.Format(time.DateOnly))
		if err != nil {
			return []*domain.Events{}, fmt.Errorf("failed to get weekday from date: %w", err)
		}

		daysByEventID[day.EventID] = append(daysByEventID[day.EventID], domain.EventDays{
			Date:               day.Date.Format(time.DateOnly),
			WeekDay:            weekDay,
			WeekDayShortFormat: weekDayShortform,
			StartTime:          start,
			EndTime:            end,
		})
	}

	var results []*domain.Events

	for _, e := range *eventEntities {
		results = append(results, &domain.Events{
			Title:                 e.Title,
			Description:           e.Description.String,
			AdditionalInformation: e.AdditionalInformation.String,
			Location:              e.Location.String,
			RegistrationLink:      e.RegistrationLink.String,
			Days:                  daysByEventID[e.ID], // may be nil and that's okay
		})
	}

	return results, nil
}
