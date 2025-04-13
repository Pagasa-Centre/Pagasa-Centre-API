package contracts

import (
	"context"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type UserService interface {
	GetUserById(ctx context.Context, id string) (*entity.User, error)
}
