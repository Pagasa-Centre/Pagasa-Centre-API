package mappers

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/events/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/utils"
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

func EventDomainToEntity(event domain.Events) *entity.Event {
	eventID := uuid.New().String()
	now := time.Now()

	eventEntity := &entity.Event{
		ID:                    eventID,
		Title:                 event.Title,
		Description:           null.StringFrom(event.Description),
		AdditionalInformation: null.StringFrom(event.AdditionalInformation),
		Location:              null.StringFrom(event.Location),
		RegistrationLink:      null.StringFrom(event.RegistrationLink),
		CreatedAt:             null.TimeFrom(now),
		UpdatedAt:             null.TimeFrom(now),
	}

	return eventEntity
}

func EventDaysDomainToEntities(eventID string, domainEventDays []domain.EventDays) []entity.EventDay {
	now := time.Now()

	var eventDaysEntity []entity.EventDay
	for _, d := range domainEventDays {
		eventDaysEntity = append(
			eventDaysEntity,
			entity.EventDay{
				ID:        uuid.New().String(),
				EventID:   eventID,
				Date:      parseDate(d.Date),
				StartTime: null.TimeFrom(d.StartTime),
				EndTime:   null.TimeFrom(d.EndTime),
				CreatedAt: null.TimeFrom(now),
				UpdatedAt: null.TimeFrom(now),
			},
		)
	}

	return eventDaysEntity
}

// Helper: Convert string (YYYY-MM-DD) to time.Time
func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{} // or handle error as appropriate
	}

	return t
}

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
