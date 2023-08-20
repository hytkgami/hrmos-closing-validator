package usecase

import (
	"context"

	"github.com/hytkgami/hrmos-closing-validator/domain"
)

type WorkReportRepository interface {
	GetMonthlyWorkReportsByUserID(ctx context.Context, token, month string, userID int) ([]*domain.WorkReport, error)
}
