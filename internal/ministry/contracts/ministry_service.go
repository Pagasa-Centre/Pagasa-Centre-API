package contracts

import "context"

type MinistryService interface {
	AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error
}
