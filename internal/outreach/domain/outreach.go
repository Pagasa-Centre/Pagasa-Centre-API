package domain

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"time"
)

type Outreach struct {
	ID           string
	Name         string
	AddressLine1 string
	AddressLine2 string
	Postcode     string
	City         string
	Country      string
	VenueName    string
	Region       string
	ThumbnailURL string
}

type Service struct {
	OutreachID string
	StartTime  time.Time
	EndTime    time.Time
	Day        string
}

func ServiceEntitiesToDomain(serviceEntities entity.OutreachServiceSlice) []*Service {
	var services []*Service

	for _, s := range serviceEntities {
		services = append(services, &Service{
			OutreachID: s.OutreachID,
			StartTime:  s.StartTime,
			EndTime:    s.EndTime,
			Day:        s.Day,
		})
	}

	return services
}
func OutreachEntitiesToDomain(outreachEntities []*entity.Outreach) []*Outreach {
	var outreaches []*Outreach
	for _, ent := range outreachEntities {
		outreaches = append(outreaches, ToDomain(ent))
	}

	return outreaches
}

func ToDomain(outreachEntity *entity.Outreach) *Outreach {
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

	outreach := Outreach{
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
