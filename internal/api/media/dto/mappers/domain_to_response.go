package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/media/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/domain"
)

func ToMediaResponse(mediaDomain *domain.Media) dto.Media {
	return dto.Media{
		ID:             mediaDomain.ID,
		Title:          mediaDomain.Title,
		Description:    mediaDomain.Description,
		YoutubeVideoID: mediaDomain.YouTubeVideoID,
		Category:       mediaDomain.Category,
		PublishedAt:    mediaDomain.PublishedAt,
		ThumbnailURL:   mediaDomain.ThumbnailURL,
	}
}

func ToGetAllMediaResponse(domainMedias []*domain.Media, message string) dto.GetAllMediaResponse {
	if domainMedias == nil {
		return dto.GetAllMediaResponse{
			Message: "No media found.",
			Media:   nil,
		}
	}
	var mediaResponses []dto.Media
	for _, m := range domainMedias {
		mediaResponses = append(mediaResponses, ToMediaResponse(m))
	}

	return dto.GetAllMediaResponse{
		Message: message,
		Media:   mediaResponses,
	}
}
