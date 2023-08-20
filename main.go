package main

import (
	"context"
	"encoding/json"

	"github.com/hytkgami/hrmos-closing-validator/repository"
	"github.com/hytkgami/hrmos-closing-validator/usecase"
)

func main() {
	ctx := context.Background()
	tokenInteractor := &usecase.TokenInteractor{TokenRepository: &repository.TokenRepository{}}
	token, err := tokenInteractor.Get(ctx)
	if err != nil {
		panic(err)
	}
	workReportInteractor := &usecase.WorkReportInteractor{WorkReportRepository: &repository.WorkReportRepository{}}
	reports, err := workReportInteractor.GetMonthlyWorkReportsByUserID(ctx, token.Token, "2023-08", 55)
	if err != nil {
		panic(err)
	}
	_, err = tokenInteractor.Delete(ctx, token.Token)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(reports)
	if err != nil {
		panic(err)
	}
	println(string(b))
}
