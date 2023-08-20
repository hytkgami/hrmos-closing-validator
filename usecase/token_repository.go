package usecase

import (
	"context"
	"time"
)

type TokenResult struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type TokenRepository interface {
	Get(ctx context.Context, apiKey string) (*TokenResult, error)
	Delete(ctx context.Context, token string) (*TokenResult, error)
}
