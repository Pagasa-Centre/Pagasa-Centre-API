package media

import (
	"context"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/storage"
)

type MediaService interface {
	All(ctx context.Context) ([]*domain.Media, error)
	BulkInsert(ctx context.Context, videos []*entity.Medium) error
}

type service struct {
	logger *zap.Logger
	repo   storage.MediaRepository
}

func NewMediaService(logger *zap.Logger, repo storage.MediaRepository) MediaService {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) All(ctx context.Context) ([]*domain.Media, error) {
	s.logger.Info("Fetching All Media")

	mediaEntities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	mediaDomain := mappers.EntitySliceToDomainMediaSlice(mediaEntities)

	return mediaDomain, nil
}

func (s *service) BulkInsert(ctx context.Context, videos []*entity.Medium) error {
	s.logger.Info("Bulk inserting Media")
	return s.repo.BulkInsert(ctx, videos)
}
