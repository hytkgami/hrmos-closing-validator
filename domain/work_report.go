package domain

import (
	"fmt"
	"time"
)

type WorkReport struct {
	Day                string `json:"day"`
	SegmentTitle       string `json:"segment_title"`
	StartAt            string `json:"start_at"`
	EndAt              string `json:"end_at"`
	TotalBreakTime     string `json:"total_break_time"`
	ActualWorkingHours string `json:"actual_working_hours"`
}

func (wr *WorkReport) validateOmittedStamp() error {
	if wr.SegmentTitle != "出勤" {
		return nil
	}
	if wr.StartAt == "" {
		return fmt.Errorf("%s: start_at is empty", wr.Day)
	}
	if wr.EndAt == "" {
		return fmt.Errorf("%s: end_at is empty", wr.Day)
	}
	return nil
}

const (
	hourMinuteFormat = "15:04"

	BREAK_SECONDS_THRESHOLD_1   = 60 * 60
	WORK_TIME_HOURS_THRESHOLD_1 = 8
	BREAK_SECONDS_THRESHOLD_2   = 45 * 60
	WORK_TIME_HOURS_THRESHOLD_2 = 6
)

func (wr *WorkReport) validateBreakTime() error {
	if wr.SegmentTitle != "出勤" {
		return nil
	}
	if wr.TotalBreakTime == "" {
		return fmt.Errorf("%s: total_break_time is empty", wr.Day)
	}
	startAt, err := time.Parse(hourMinuteFormat, wr.StartAt)
	if err != nil {
		return fmt.Errorf("%s: start_at is invalid format", wr.Day)
	}
	endAt, err := time.Parse(hourMinuteFormat, wr.EndAt)
	if err != nil {
		return fmt.Errorf("%s: end_at is invalid format", wr.Day)
	}
	totalBreakTime, err := time.Parse(hourMinuteFormat, wr.TotalBreakTime)
	if err != nil {
		return fmt.Errorf("%s: total_break_time is invalid format", wr.Day)
	}
	h, m, s := totalBreakTime.Clock()
	totalBreakTimeSeconds := h*60*60 + m*60 + s

	workTime := endAt.Sub(startAt)
	switch {
	case workTime.Hours() > WORK_TIME_HOURS_THRESHOLD_1:
		if totalBreakTimeSeconds < BREAK_SECONDS_THRESHOLD_1 {
			return fmt.Errorf("%s: total_break_time is less than 60 minutes", wr.Day)
		}
	case workTime.Hours() > WORK_TIME_HOURS_THRESHOLD_2:
		if totalBreakTimeSeconds < BREAK_SECONDS_THRESHOLD_2 {
			return fmt.Errorf("%s: total_break_time is less than 45 minutes", wr.Day)
		}
	default:
		return nil
	}
	return nil
}
