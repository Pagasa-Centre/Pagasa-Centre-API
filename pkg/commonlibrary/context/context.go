package context

import (
	"context"
	"errors"
)

type contextKey string

const userIDKey contextKey = "userID"

// SetUserID sets the user ID in the context.
func SetUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserID retrieves the user ID from the context.
func GetUserID(ctx context.Context) (int, error) {
	val := ctx.Value(userIDKey)
	if val == nil {
		return 0, errors.New("user ID not found in context")
	}

	uid, ok := val.(int)
	if !ok {
		return 0, errors.New("user ID has wrong type")
	}

	return uid, nil
}
