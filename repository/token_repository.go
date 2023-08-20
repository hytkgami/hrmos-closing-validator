package repository

import (
	"context"
	"encoding/json"

	"github.com/hytkgami/hrmos-closing-validator/usecase"
)

type TokenRepository struct{}

func (tr *TokenRepository) Get(ctx context.Context, apiKey string) (*usecase.TokenResult, error) {
	b, err := get(ctx, "authentication/token", nil, map[string]string{
		"Authorization": "Basic " + apiKey,
	})
	if err != nil {
		return nil, err
	}
	var result usecase.TokenResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (tr *TokenRepository) Delete(ctx context.Context, token string) (*usecase.TokenResult, error) {
	b, err := delete(ctx, "authentication/destroy", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	if err != nil {
		return nil, err
	}
	var result usecase.TokenResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
