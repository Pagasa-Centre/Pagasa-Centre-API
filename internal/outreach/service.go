package outreach

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	domain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/storage"
)

type OutreachService interface {
	GetAllOutreaches(ctx context.Context) ([]*domain.Outreach, []*domain.Service, error)
}

type service struct {
	logger       *zap.Logger
	outreachRepo storage.OutreachRepository
}

func NewOutreachService(
	logger *zap.Logger,
	outreachRepo storage.OutreachRepository,
) OutreachService {
	return &service{
		logger:       logger,
		outreachRepo: outreachRepo,
	}
}

func (os *service) GetAllOutreaches(ctx context.Context) ([]*domain.Outreach, []*domain.Service, error) {
	outreachesEntities, err := os.outreachRepo.GetAll(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get all outreaches: %s", err)
	}

	var services entity.OutreachServiceSlice

	for _, ent := range outreachesEntities {
		// 1. Get services for each outreach
		outreachServices, err := os.outreachRepo.GetServicesByOutreachID(ctx, ent.ID)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get outreach services: %s", err)
		}

		if len(outreachServices) == 0 {
			continue
		}
		// 2. add to services slice
		services = append(services, outreachServices...)
	}

	outreaches := domain.OutreachEntitiesToDomain(outreachesEntities)
	serv := domain.ServiceEntitiesToDomain(services)

	return outreaches, serv, nil
}
