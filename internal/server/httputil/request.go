package httputil

import (
	"context"

	"github.com/edmarfelipe/rss-scraper/internal/database"
)

// requestCtxKeyType is a custom type to avoid collisions with other context keys.
type requestCtxKeyType string

// userCtxKey is the key for the user in the request context.
const userCtxKey requestCtxKeyType = "user"

// WithUser adds the user to the request context.
func WithUser(ctx context.Context, user *database.User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

// GetUser returns the user from the request context.
func GetUser(ctx context.Context) *database.User {
	user, ok := ctx.Value(userCtxKey).(*database.User)
	if !ok {
		return nil
	}
	return user
}
