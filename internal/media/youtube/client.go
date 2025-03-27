package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/volatiletech/null/v8"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type YouTubeClient struct {
	APIKey     string
	ChannelID  string
	HTTPClient *http.Client
}

type VideoItem struct {
	ID struct {
		VideoID string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		PublishedAt time.Time `json:"publishedAt"`
		Thumbnails  struct {
			High struct {
				URL string `json:"url"`
			} `json:"high"`
		} `json:"thumbnails"`
	} `json:"snippet"`
}

type YouTubeResponse struct {
	Items []VideoItem `json:"items"`
}

func NewYouTubeClient(apiKey, channelID string) *YouTubeClient {
	return &YouTubeClient{
		APIKey:     apiKey,
		ChannelID:  channelID,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *YouTubeClient) FetchVideosFromPlaylist(ctx context.Context, playlistID, category string) ([]*entity.Medium, error) {
	url := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=25&playlistId=%s&key=%s",
		playlistID, c.APIKey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ytRes struct {
		Items []struct {
			Snippet struct {
				ResourceID struct {
					VideoID string `json:"videoId"`
				} `json:"resourceId"`
				Title       string    `json:"title"`
				Description string    `json:"description"`
				PublishedAt time.Time `json:"publishedAt"`
				Thumbnails  struct {
					High struct {
						URL string `json:"url"`
					} `json:"high"`
				} `json:"thumbnails"`
			} `json:"snippet"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&ytRes); err != nil {
		return nil, err
	}

	var mediaList []*entity.Medium

	for _, item := range ytRes.Items {
		media := &entity.Medium{
			YoutubeVideoID: item.Snippet.ResourceID.VideoID,
			Title:          item.Snippet.Title,
			Description:    null.StringFrom(item.Snippet.Description),
			Category:       category,
			PublishedAt:    item.Snippet.PublishedAt,
			ThumbnailURL:   item.Snippet.Thumbnails.High.URL,
		}
		mediaList = append(mediaList, media)
	}

	return mediaList, nil
}
