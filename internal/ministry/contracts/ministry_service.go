package contracts

import (
	"context"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
)

type MinistryService interface {
	AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error
	GetByID(ctx context.Context, ministryID string) (*domain.Ministry, error)
}
