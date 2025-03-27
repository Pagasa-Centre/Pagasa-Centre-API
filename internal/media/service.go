package media

import (
	"context"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/storage"
)

type MediaService interface {
	All(ctx context.Context) ([]*entity.Medium, error)
	BulkInsert(ctx context.Context, videos []*entity.Medium) error
}

type service struct {
	logger zap.SugaredLogger
	repo   storage.MediaRepository
}

func NewMediaService(logger zap.SugaredLogger, repo storage.MediaRepository) MediaService {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) All(ctx context.Context) ([]*entity.Medium, error) {
	s.logger.Info("Fetching All Media")
	return s.repo.GetAll(ctx)
}

func (s *service) BulkInsert(ctx context.Context, videos []*entity.Medium) error {
	s.logger.Info("Bulk inserting Media")
	return s.repo.BulkInsert(ctx, videos)
}
