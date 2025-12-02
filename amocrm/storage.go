package amocrm

import "context"

// TokenStorage is an interface for storing and retrieving OAuth2 tokens
type TokenStorage interface {
	// Save saves a token for the given domain
	Save(ctx context.Context, domain string, token *Token) error

	// Load loads a token for the given domain
	Load(ctx context.Context, domain string) (*Token, error)

	// HasToken checks if a token exists for the given domain
	HasToken(ctx context.Context, domain string) (bool, error)
}
