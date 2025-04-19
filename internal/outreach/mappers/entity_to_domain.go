package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
)

func ServiceEntitiesToDomain(serviceEntities entity.OutreachServiceSlice) []*domain.Service {
	var services []*domain.Service

	for _, s := range serviceEntities {
		services = append(services, &domain.Service{
			OutreachID: s.OutreachID,
			StartTime:  s.StartTime,
			EndTime:    s.EndTime,
			Day:        s.Day,
		})
	}

	return services
}

func OutreachEntitiesToDomain(outreachEntities []*entity.Outreach) []*domain.Outreach {
	var outreaches []*domain.Outreach
	for _, ent := range outreachEntities {
		outreaches = append(outreaches, ToDomain(ent))
	}

	return outreaches
}

func ToDomain(outreachEntity *entity.Outreach) *domain.Outreach {
	var addressLine2 string

	var postcode string

	var venueName string

	var region string

	var thumbnailURL string

	if outreachEntity.AddressLine2.Valid {
		addressLine2 = outreachEntity.AddressLine2.String
	}

	if outreachEntity.PostCode.Valid {
		postcode = outreachEntity.PostCode.String
	}

	if outreachEntity.VenueName.Valid {
		venueName = outreachEntity.VenueName.String
	}

	if outreachEntity.Region.Valid {
		region = outreachEntity.Region.String
	}

	if outreachEntity.ThumbnailURL.Valid {
		thumbnailURL = outreachEntity.ThumbnailURL.String
	}

	outreach := domain.Outreach{
		ID:           outreachEntity.ID,
		Name:         outreachEntity.Name,
		AddressLine1: outreachEntity.AddressLine1,
		AddressLine2: addressLine2,
		Postcode:     postcode,
		City:         outreachEntity.City,
		Country:      outreachEntity.Country,
		VenueName:    venueName,
		Region:       region,
		ThumbnailURL: thumbnailURL,
	}

	return &outreach
}
