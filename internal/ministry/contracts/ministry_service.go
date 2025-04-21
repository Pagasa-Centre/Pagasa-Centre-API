package contracts

import (
	"context"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
)

type MinistryService interface {
	GetByID(ctx context.Context, ministryID string) (*domain.Ministry, error)
}
