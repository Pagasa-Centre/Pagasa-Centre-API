package outreach

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	domain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/storage"
)

type OutreachService interface {
	All(ctx context.Context) ([]*domain.Outreach, error)
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

func (os *service) All(ctx context.Context) ([]*domain.Outreach, error) {
	outreachesEntities, err := os.outreachRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all outreaches: %s", err)
	}

	outreaches := domain.EntitiesToDomain(outreachesEntities)

	return outreaches, nil
}
