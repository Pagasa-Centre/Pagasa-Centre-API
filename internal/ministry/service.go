package ministry

import (
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
)

type MinistryService interface {
	All() error
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

func (ms *ministryService) All() error {

	return nil
}
