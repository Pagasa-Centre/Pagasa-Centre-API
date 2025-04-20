package contracts

import (
	"context"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

type UserService interface {
	GetUserById(ctx context.Context, id string) (*domain.User, error)
}
