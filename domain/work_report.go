package domain

type WorkReport struct {
	Day                string `json:"day"`
	SegmentTitle       string `json:"segment_title"`
	StartAt            string `json:"start_at"`
	EndAt              string `json:"end_at"`
	TotalBreakTime     string `json:"total_break_time"`
	ActualWorkingHours string `json:"actual_working_hours"`
}
