package mapper

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/domain"
)

func EntityToDomainMedia(m *entity.Medium) *domain.Media {
	if m == nil {
		return nil
	}

	return &domain.Media{
		ID:             m.ID,
		Title:          m.Title,
		Description:    m.Description.String,
		YouTubeVideoID: m.YoutubeVideoID,
		Category:       m.Category,
		PublishedAt:    m.PublishedAt,
		ThumbnailURL:   m.ThumbnailURL,
	}
}

func EntitySliceToDomainMediaSlice(mediaEntities []*entity.Medium) []*domain.Media {
	var domainMedias []*domain.Media

	for _, m := range mediaEntities {
		domainMedias = append(domainMedias, EntityToDomainMedia(m))
	}

	return domainMedias
}
