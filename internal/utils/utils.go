package utils

import (
	"fmt"
	"time"
)

// GetWeekdayFromDate returns both the full and short form of the weekday
// e.g., "Monday" and "Mon" for the date "2025-08-18"
func GetWeekdayFromDate(dateStr string) (fullDay string, shortDay string, err error) {
	parsedDate, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return "", "", fmt.Errorf("invalid date format: %w", err)
	}

	weekday := parsedDate.Weekday()

	return weekday.String(), weekday.String()[:3], nil
}
