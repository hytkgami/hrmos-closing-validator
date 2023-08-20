package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hytkgami/hrmos-closing-validator/domain"
)

type WorkReportRepository struct{}

func (wr *WorkReportRepository) GetMonthlyWorkReportsByUserID(ctx context.Context, token, month string, userID int) ([]*domain.WorkReport, error) {
	b, err := get(ctx, "work_outputs/monthly/"+month, map[string]string{
		"user_id": fmt.Sprintf("%d", userID),
		"limit":   fmt.Sprintf("%d", 31),
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	if err != nil {
		return nil, err
	}
	var result []*domain.WorkReport
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
