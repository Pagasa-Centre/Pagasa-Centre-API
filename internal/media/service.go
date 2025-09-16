package media

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/storage"
)

type Service interface {
	GetAllMedia(ctx context.Context) ([]*domain.Media, error)
	BulkInsert(ctx context.Context, videos []*entity.Medium) error
}

type service struct {
	logger *zap.Logger
	repo   storage.MediaRepository
}

func NewMediaService(logger *zap.Logger, repo storage.MediaRepository) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) GetAllMedia(ctx context.Context) ([]*domain.Media, error) {
	mediaEntities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all media: %w", err)
	}

	if len(mediaEntities) == 0 {
		return nil, nil
	}

	return mapper.EntitySliceToDomainMediaSlice(mediaEntities), nil
}

func (s *service) BulkInsert(ctx context.Context, videos []*entity.Medium) error {
	s.logger.Info("Bulk inserting Media")
	return s.repo.BulkInsert(ctx, videos)
}
