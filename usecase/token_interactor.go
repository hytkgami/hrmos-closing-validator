package usecase

import (
	"context"
	"errors"
	"os"
)

type TokenInteractor struct {
	TokenRepository
}

func (itr *TokenInteractor) Get(ctx context.Context) (*TokenResult, error) {
	apiKey := os.Getenv("HRMOS_API_KEY")
	if apiKey == "" {
		return nil, errors.New("HRMOS_API_KEY is not set")
	}
	return itr.TokenRepository.Get(ctx, apiKey)
}

func (itr *TokenInteractor) Delete(ctx context.Context, token string) (*TokenResult, error) {
	return itr.TokenRepository.Delete(ctx, token)
}
