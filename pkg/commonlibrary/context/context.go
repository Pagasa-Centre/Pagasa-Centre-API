package context

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type contextKey string

const (
	userIDIntKey  contextKey = "userIDInt"
	userIDUUIDKey contextKey = "userIDUUID"
	userIDKey     contextKey = "userIDKey"
)

//
// --- INT VERSION ---
//

// SetUserIDInt sets an integer user ID in the context.
func SetUserIDInt(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDIntKey, userID)
}

// GetUserIDInt retrieves an integer user ID from the context.
func GetUserIDInt(ctx context.Context) (int, error) {
	val := ctx.Value(userIDIntKey)
	if val == nil {
		return 0, errors.New("int user ID not found in context")
	}

	uid, ok := val.(int)
	if !ok {
		return 0, errors.New("int user ID has wrong type")
	}

	return uid, nil
}

//
// --- UUID VERSION ---
//

// SetUserIDUUID sets a UUID user ID in the context.
func SetUserIDUUID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDUUIDKey, userID)
}

// GetUserIDUUID retrieves a UUID user ID from the context.
func GetUserIDUUID(ctx context.Context) (uuid.UUID, error) {
	val := ctx.Value(userIDUUIDKey)
	if val == nil {
		return uuid.Nil, errors.New("UUID user ID not found in context")
	}

	uid, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("UUID user ID has wrong type")
	}

	return uid, nil
}

// SetUserIDString sets the user ID (as a string) in context.
func SetUserIDString(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDString retrieves the user ID (as a string) from context.
func GetUserIDString(ctx context.Context) (string, error) {
	val := ctx.Value(userIDKey)
	if val == nil {
		return "", errors.New("user ID not found in context")
	}

	uid, ok := val.(string)
	if !ok {
		return "", errors.New("user ID has wrong type")
	}

	return uid, nil
}
