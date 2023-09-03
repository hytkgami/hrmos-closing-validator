package domain_test

import (
	"errors"
	"testing"

	"github.com/hytkgami/hrmos-closing-validator/domain"
)

func TestValidateOmittedStamp(t *testing.T) {
	cases := []struct {
		name string
		wr   *domain.WorkReport
		want error
	}{
		{
			name: "start_at is empty",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "",
				EndAt:              "19:00",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: errors.New("2021-08-01: start_at is empty"),
		},
		{
			name: "end_at is empty",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: errors.New("2021-08-01: end_at is empty"),
		},
		{
			name: "start_at and end_at are not empty",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:00",
				TotalBreakTime:     "01:00",
				ActualWorkingHours: "08:00",
			},
			want: nil,
		},
		{
			name: "on holiday",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "公休",
				StartAt:            "",
				EndAt:              "",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: nil,
		},
		{
			name: "on paid holiday",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "有給休暇",
				StartAt:            "",
				EndAt:              "",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: nil,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := domain.ValidateOmittedStamp(c.wr)
			if err != nil && err.Error() != c.want.Error() {
				t.Errorf("got: %v, want: %v", err, c.want)
			}
		})
	}
}

func TestValidateBreakTime(t *testing.T) {
	cases := []struct {
		name string
		wr   *domain.WorkReport
		want error
	}{
		{
			name: "total_break_time is empty",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:00",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: errors.New("2021-08-01: total_break_time is empty"),
		},
		{
			name: "total_break_time is invalid format",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:00",
				TotalBreakTime:     "01:00:00",
				ActualWorkingHours: "",
			},
			want: errors.New("2021-08-01: total_break_time is invalid format"),
		},
		{
			name: "start_at is invalid format",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00:00",
				EndAt:              "19:00",
				TotalBreakTime:     "01:00",
				ActualWorkingHours: "",
			},
			want: errors.New("2021-08-01: start_at is invalid format"),
		},
		{
			name: "end_at is invalid format",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:00:00",
				TotalBreakTime:     "01:00",
				ActualWorkingHours: "",
			},
			want: errors.New("2021-08-01: end_at is invalid format"),
		},
		{
			name: "total_break_time is less than 45 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "16:30",
				TotalBreakTime:     "00:30",
				ActualWorkingHours: "06:00",
			},
			want: errors.New("2021-08-01: total_break_time is less than 45 minutes"),
		},
		{
			name: "total_break_time is less than 60 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:00",
				TotalBreakTime:     "00:45",
				ActualWorkingHours: "08:15",
			},
			want: errors.New("2021-08-01: total_break_time is less than 60 minutes"),
		},
		{
			name: "total_break_time is more than 60 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:00",
				TotalBreakTime:     "01:00",
				ActualWorkingHours: "08:00",
			},
			want: nil,
		},
		{
			name: "total_break_time is more than 45 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "16:45",
				TotalBreakTime:     "00:45",
				ActualWorkingHours: "06:00",
			},
			want: nil,
		},
		{
			name: "work_time is more than 8 hours and total_break_time is less than 60 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "18:45",
				TotalBreakTime:     "00:45",
				ActualWorkingHours: "08:00",
			},
			want: errors.New("2021-08-01: total_break_time is less than 60 minutes"),
		},
		{
			name: "work_time is more than 8 hours and total_break_time is less than 45 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "19:30",
				TotalBreakTime:     "00:30",
				ActualWorkingHours: "08:00",
			},
			want: errors.New("2021-08-01: total_break_time is less than 60 minutes"),
		},
		{
			name: "work_time is less than 6 hours and total_break_time is less than 45 minutes",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "出勤",
				StartAt:            "10:00",
				EndAt:              "16:00",
				TotalBreakTime:     "00:30",
				ActualWorkingHours: "05:30",
			},
			want: nil,
		},
		{
			name: "on holiday",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "公休",
				StartAt:            "",
				EndAt:              "",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: nil,
		},
		{
			name: "on paid holiday",
			wr: &domain.WorkReport{
				Day:                "2021-08-01",
				SegmentTitle:       "有給休暇",
				StartAt:            "",
				EndAt:              "",
				TotalBreakTime:     "",
				ActualWorkingHours: "",
			},
			want: nil,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := domain.ValidateBreakTime(c.wr)
			if err != nil && err.Error() != c.want.Error() {
				t.Errorf("got: %v, want: %v", err, c.want)
			}
		})
	}
}
