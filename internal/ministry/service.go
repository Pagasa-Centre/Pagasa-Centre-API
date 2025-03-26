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

type ministryService struct {
	logger       zap.SugaredLogger
	ministryRepo storage.MinistryRepository
	jwtSecret    string
}

func NewMinistryService(
	logger zap.SugaredLogger,
	ministryRepo storage.MinistryRepository,
	jwtSecret string,
) MinistryService {
	return &ministryService{
		logger:       logger,
		ministryRepo: ministryRepo,
		jwtSecret:    jwtSecret,
	}
}

func (ms *ministryService) All(ctx context.Context) ([]*domain.Ministry, error) {
	ministries := []*domain.Ministry{}

	ministryEntities, err := ms.ministryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ministries: %s", err)
	}

	for _, entity := range ministryEntities {
		ministry := domain.ToDomain(entity)
		ministries = append(ministries, ministry)
	}

	return ministries, nil
}
