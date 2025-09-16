package cron

import (
	"context"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
	youtube "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/youtube"
)

type MediaCronJob struct {
	Logger        *zap.Logger
	YouTubeClient *youtube.YouTubeClient
	MediaService  media.Service
}

func NewMediaCronJob(logger *zap.Logger, ytClient *youtube.YouTubeClient, mediaService media.Service) *MediaCronJob {
	return &MediaCronJob{
		Logger:        logger,
		YouTubeClient: ytClient,
		MediaService:  mediaService,
	}
}

var playlistMap = map[string]string{
	"Bible Study":         "PL_isAVg17p0Cy27I5EK7PN88f6dKhS5Bd",
	"Sunday Preachings":   "PL_isAVg17p0DiqigrZAEcl8qNdupG7Vao",
	"Evangelistic Nights": "PL_isAVg17p0BeZMYzxWNzB9OZI1e253eZ",
} // todo: move to domain

func (job *MediaCronJob) Start() {
	c := cron.New()

	// Every 3 days at 11 PM
	spec := "0 23 */3 * *"

	_, err := c.AddFunc(spec, func() {
		job.Logger.Info("Starting media fetch cron job")

		ctx := context.Background()

		for category, playlistID := range playlistMap {
			videos, err := job.YouTubeClient.FetchVideosFromPlaylist(ctx, playlistID, category)
			if err != nil {
				job.Logger.Sugar().Errorw("Error fetching YouTube videos", "category", category, "error", err)
				continue
			}

			if len(videos) == 0 {
				job.Logger.Sugar().Infow("No new videos found", "category", category)
				continue
			}

			err = job.MediaService.BulkInsert(ctx, videos)
			if err != nil {
				job.Logger.Sugar().Errorw("Error storing videos in DB", "error", err)
			}
		}
	})
	if err != nil {
		job.Logger.Sugar().Errorw("Failed to schedule media cron job", "error", err)
		return
	}

	job.Logger.Sugar().Infow("Media cron job scheduled")
	c.Start()
}

func (job *MediaCronJob) RunOnce() {
	job.Logger.Sugar().Infow("Manually running media fetch job")

	ctx := context.Background()

	for category, playlistID := range playlistMap {
		videos, err := job.YouTubeClient.FetchVideosFromPlaylist(ctx, playlistID, category)
		if err != nil {
			job.Logger.Sugar().Errorw("Error fetching YouTube videos", "category", category, "error", err)
			continue
		}

		err = job.MediaService.BulkInsert(ctx, videos)
		if err != nil {
			job.Logger.Sugar().Errorw("Error storing videos in DB", "error", err)
		}
	}
}
