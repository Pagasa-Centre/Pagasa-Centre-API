package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/domain"
)

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
