package domain

import "time"

type Media struct {
	ID             int
	Title          string
	Description    string
	YouTubeVideoID string
	Category       string
	PublishedAt    time.Time
	ThumbnailURL   string
}
