package ministry

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
)

type MinistryService interface {
	All(ctx context.Context) ([]*domain.Ministry, error)
}

type service struct {
	logger       zap.SugaredLogger
	ministryRepo storage.MinistryRepository
}

func NewMinistryService(
	logger zap.SugaredLogger,
	ministryRepo storage.MinistryRepository,
) MinistryService {
	return &service{
		logger:       logger,
		ministryRepo: ministryRepo,
	}
}

func (ms *service) All(ctx context.Context) ([]*domain.Ministry, error) {
	ministryEntities, err := ms.ministryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ministries: %s", err)
	}

	var ministries []*domain.Ministry
	for _, entity := range ministryEntities {
		ministries = append(ministries, domain.ToDomain(entity))
	}

	return ministries, nil
}
