package dbt

import (
	d_types "github.com/leapforce-libraries/go_dbt/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type Step struct {
	Id                 int64                  `json:"id"`
	RunId              int64                  `json:"run_id"`
	AccountId          int64                  `json:"account_id"`
	Index              int64                  `json:"index"`
	Status             int64                  `json:"status"`
	Name               string                 `json:"name"`
	Logs               string                 `json:"logs"`
	DebugLogs          string                 `json:"debug_logs"`
	LogLocation        string                 `json:"log_location"`
	LogPath            string                 `json:"log_path"`
	DebugLogPath       string                 `json:"debug_log_path"`
	LogArchiveType     string                 `json:"log_archive_type"`
	TruncatedDebugLogs string                 `json:"truncated_debug_logs"`
	CreatedAt          d_types.DateTimeString `json:"created_at"`
	UpdatedAt          d_types.DateTimeString `json:"updated_at"`
	StartedAt          d_types.DateTimeString `json:"started_at"`
	FinishedAt         d_types.DateTimeString `json:"finished_at"`
	StatusColor        string                 `json:"status_color"`
	StatusHumanized    string                 `json:"status_humanized"`
	Duration           go_types.TimeString    `json:"duration"`
	DurationHumanized  string                 `json:"duration_humanized"`
}
