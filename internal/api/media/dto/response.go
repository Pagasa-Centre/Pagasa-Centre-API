package dto

import "time"

type Media struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	YoutubeVideoID string    `json:"youtube_video_id"`
	Category       string    `json:"category"`
	PublishedAt    time.Time `json:"published_at"`
	ThumbnailURL   string    `json:"thumbnail_url"`
}

type GetAllMediaResponse struct {
	Message string  `json:"message"`
	Media   []Media `json:"media"`
}
