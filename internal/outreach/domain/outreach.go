package domain

import "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"

type Outreach struct {
	ID           int
	Name         string
	AddressLine1 string
	AddressLine2 string
	Postcode     string
	City         string
	Country      string
}

func EntitiesToDomain(outreachEntities []*entity.Outreach) []*Outreach {
	var outreaches []*Outreach
	for _, ent := range outreachEntities {
		outreaches = append(outreaches, ToDomain(ent))
	}

	return outreaches
}

func ToDomain(outreachEntity *entity.Outreach) *Outreach {
	var addressLine2 string

	if outreachEntity.AddressLine2.Valid {
		addressLine2 = outreachEntity.AddressLine2.String
	}

	outreach := Outreach{
		ID:           outreachEntity.ID,
		Name:         outreachEntity.Name,
		AddressLine1: outreachEntity.AddressLine1,
		AddressLine2: addressLine2,
		Postcode:     outreachEntity.PostCode,
		City:         outreachEntity.City,
		Country:      outreachEntity.Country,
	}

	return &outreach
}
