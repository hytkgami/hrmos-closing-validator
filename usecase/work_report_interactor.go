package usecase

import (
	"context"

	"github.com/hytkgami/hrmos-closing-validator/domain"
)

type WorkReportInteractor struct {
	WorkReportRepository
}

func (itr *WorkReportInteractor) GetMonthlyWorkReportsByUserID(ctx context.Context, token, month string, userID int) ([]*domain.WorkReport, error) {
	return itr.WorkReportRepository.GetMonthlyWorkReportsByUserID(ctx, token, month, userID)
}
