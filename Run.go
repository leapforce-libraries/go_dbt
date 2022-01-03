package dbt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	d_types "github.com/leapforce-libraries/go_dbt/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Run struct {
	ID                 int64                  `json:"id"`
	TriggerID          int64                  `json:"trigger_id"`
	AccountID          int64                  `json:"account_id"`
	ProjectID          int64                  `json:"project_id"`
	JobID              int64                  `json:"job_id"`
	JobDefinitionID    int64                  `json:"job_definition_id"`
	Status             int64                  `json:"status"`
	GitBranch          string                 `json:"git_branch"`
	StatusMessage      string                 `json:"status_message"`
	StatusHumanized    string                 `json:"status_humanized"`
	InProgress         bool                   `json:"in_progress"`
	IsComplete         bool                   `json:"is_complete"`
	IsSuccess          bool                   `json:"is_success"`
	IsError            bool                   `json:"is_error"`
	IsCancelled        bool                   `json:"is_cancelled"`
	DBTVersion         string                 `json:"dbt_version"`
	CreatedAt          d_types.DateTimeString `json:"created_at"`
	UpdatedAt          d_types.DateTimeString `json:"updated_at"`
	DequeudAt          d_types.DateTimeString `json:"dequeued_at"`
	StartedAt          d_types.DateTimeString `json:"started_at"`
	FinishedAt         d_types.DateTimeString `json:"finished_at"`
	LastCheckedAt      d_types.DateTimeString `json:"last_checked_at"`
	LastHeartbeatAt    d_types.DateTimeString `json:"last_heartbeat_at"`
	OwnerThreadID      string                 `json:"owner_thread_id"`
	ExecutedByThreadID string                 `json:"executed_by_thread_id"`
	ArtifactsSaved     bool                   `json:"artifacts_saved"`
	ArtifactsS3Saved   string                 `json:"artifacts_s3_path"`
	HasDocsGenerated   bool                   `json:"has_docs_generated"`
	Trigger            json.RawMessage        `json:"triggers"`
	Job                json.RawMessage        `json:"job"`
	Environment        json.RawMessage        `json:"environment"`
	Duration           go_types.TimeString    `json:"duration"`
	QueuedDuration     go_types.TimeString    `json:"queued_duration"`
	RunDuration        go_types.TimeString    `json:"run_duration"`
}

type GetRunsConfig struct {
	AccountID          int64
	IncludeTrigger     bool
	IncludeJob         bool
	IncludeRepository  bool
	IncludeEnvironment bool
	JobDefinitionID    *int64
	OrderBy            *string
	Offset             *int64
	Limit              *int64
}

// GetRuns returns all runs
//
func (service *Service) GetRuns(config *GetRunsConfig) (*[]Run, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}
	runs := []Run{}

	includeRelated := []string{}
	if config.IncludeTrigger {
		includeRelated = append(includeRelated, "trigger")
	}
	if config.IncludeJob {
		includeRelated = append(includeRelated, "job")
	}
	if config.IncludeRepository {
		includeRelated = append(includeRelated, "repository")
	}
	if config.IncludeEnvironment {
		includeRelated = append(includeRelated, "environment")
	}

	params := url.Values{}
	if len(includeRelated) > 0 {
		params.Add("include_related", fmt.Sprintf("[%s]", strings.Join(includeRelated, ",")))
	}
	if config.JobDefinitionID != nil {
		params.Add("job_defintion_id", fmt.Sprintf("%d", *config.JobDefinitionID))
	}
	if config.OrderBy != nil {
		params.Add("order_by", *config.OrderBy)
	}

	limit := int64(100)
	offset := int64(0)

	if config.Limit != nil {
		limit = *config.Limit
	}

	if config.Offset != nil {
		offset = *config.Offset
	}

	params.Add("limit", fmt.Sprintf("%d", limit))

	for {
		params.Set("offset", fmt.Sprintf("%d", offset))

		_runs := []Run{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.url(fmt.Sprintf("accounts/%d/runs?%s", config.AccountID, params.Encode())),
			ResponseModel: &_runs,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		runs = append(runs, _runs...)

		if int64(len(_runs)) < limit {
			break
		}

		offset += limit
	}

	return &runs, nil
}
