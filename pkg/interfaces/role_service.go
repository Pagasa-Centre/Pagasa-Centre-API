package interfaces

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type RoleService interface {
	Fetch(ctx context.Context, userID string) ([]string, error)
	AssignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, role string, ministryID *string) error
}
