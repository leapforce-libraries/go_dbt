package dbt

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	d_types "github.com/leapforce-libraries/go_dbt/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

const limitDefault int64 = 100

type Run struct {
	Id                 int64                  `json:"id"`
	TriggerId          int64                  `json:"trigger_id"`
	AccountId          int64                  `json:"account_id"`
	ProjectId          int64                  `json:"project_id"`
	JobId              int64                  `json:"job_id"`
	JobDefinitionId    int64                  `json:"job_definition_id"`
	Status             int64                  `json:"status"`
	GitBranch          string                 `json:"git_branch"`
	StatusMessage      string                 `json:"status_message"`
	StatusHumanized    string                 `json:"status_humanized"`
	InProgress         bool                   `json:"in_progress"`
	IsComplete         bool                   `json:"is_complete"`
	IsSuccess          bool                   `json:"is_success"`
	IsError            bool                   `json:"is_error"`
	IsCancelled        bool                   `json:"is_cancelled"`
	DbtVersion         string                 `json:"dbt_version"`
	CreatedAt          d_types.DateTimeString `json:"created_at"`
	UpdatedAt          d_types.DateTimeString `json:"updated_at"`
	DequeudAt          d_types.DateTimeString `json:"dequeued_at"`
	StartedAt          d_types.DateTimeString `json:"started_at"`
	FinishedAt         d_types.DateTimeString `json:"finished_at"`
	LastCheckedAt      d_types.DateTimeString `json:"last_checked_at"`
	LastHeartbeatAt    d_types.DateTimeString `json:"last_heartbeat_at"`
	OwnerThreadId      string                 `json:"owner_thread_id"`
	ExecutedByThreadId string                 `json:"executed_by_thread_id"`
	ArtifactsSaved     bool                   `json:"artifacts_saved"`
	ArtifactsS3Saved   string                 `json:"artifacts_s3_path"`
	HasDocsGenerated   bool                   `json:"has_docs_generated"`
	Trigger            *Trigger               `json:"triggers"`
	Job                *Job                   `json:"job"`
	RunSteps           *[]Step                `json:"run_steps"`
	Environment        *Environment           `json:"environment"`
	Duration           go_types.TimeString    `json:"duration"`
	QueuedDuration     go_types.TimeString    `json:"queued_duration"`
	RunDuration        go_types.TimeString    `json:"run_duration"`
}

type ListRunsConfig struct {
	AccountId          int64
	IncludeTrigger     bool
	IncludeJob         bool
	IncludeRepository  bool
	IncludeEnvironment bool
	IncludeDebugLogs   bool
	IncludeRunSteps    bool
	JobDefinitionId    *int64
	OrderBy            *string
	Limit              *int64
}

// ListRuns returns all runs
func (service *Service) ListRuns(config *ListRunsConfig) (*[]Run, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	var runs []Run

	offset := int64(0)

	if config.Limit == nil {
		limit := limitDefault
		config.Limit = &limit
	}

	for {
		runs_, e := service.listRuns(config, offset)
		if e != nil {
			return nil, e
		}

		runs = append(runs, *runs_...)

		if int64(len(*runs_)) < *config.Limit {
			break
		}

		offset += *config.Limit
	}

	return &runs, nil
}

func (service *Service) listRuns(config *ListRunsConfig, offset int64) (*[]Run, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	var runs []Run

	var includeRelated []string
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
	if config.IncludeDebugLogs {
		includeRelated = append(includeRelated, "debug_logs")
	}
	if config.IncludeRunSteps {
		includeRelated = append(includeRelated, "run_steps")
	}

	params := url.Values{}
	if len(includeRelated) > 0 {
		params.Add("include_related", fmt.Sprintf("[\"%s\"]", strings.Join(includeRelated, "\",\"")))
	}
	if config.JobDefinitionId != nil {
		params.Add("job_definition_id", fmt.Sprintf("%d", *config.JobDefinitionId))
	}
	if config.OrderBy != nil {
		params.Add("order_by", *config.OrderBy)
	}

	params.Add("limit", fmt.Sprintf("%d", *config.Limit))
	params.Add("offset", fmt.Sprintf("%d", offset))

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("accounts/%d/runs?%s", config.AccountId, params.Encode())),
		ResponseModel: &runs,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &runs, nil
}

type ListRunsAfterConfig struct {
	AccountId          int64
	IncludeTrigger     bool
	IncludeJob         bool
	IncludeRepository  bool
	IncludeEnvironment bool
	IncludeDebugLogs   bool
	IncludeRunSteps    bool
	JobDefinitionId    *int64
	Limit              *int64
	After              time.Time
}

// ListRunsAfter returns all runs
func (service *Service) ListRunsAfter(config *ListRunsAfterConfig) (*[]Run, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	var runs []Run

	var orderBy = "-created_at"
	var config_ = ListRunsConfig{
		AccountId:          config.AccountId,
		IncludeTrigger:     config.IncludeTrigger,
		IncludeJob:         config.IncludeJob,
		IncludeRepository:  config.IncludeRepository,
		IncludeEnvironment: config.IncludeEnvironment,
		IncludeDebugLogs:   config.IncludeDebugLogs,
		IncludeRunSteps:    config.IncludeRunSteps,
		JobDefinitionId:    config.JobDefinitionId,
		OrderBy:            &orderBy,
		Limit:              config.Limit,
	}

	if config_.Limit == nil {
		limit := limitDefault
		config_.Limit = &limit
	}

	offset := int64(0)

	for {
		runs_, e := service.listRuns(&config_, offset)
		if e != nil {
			return nil, e
		}

		for _, run := range *runs_ {
			if !run.CreatedAt.Value().After(config.After) {
				goto ready
			}

			runs = append(runs, run)
		}

		if int64(len(*runs_)) < *config_.Limit {
			break
		}

		offset += *config_.Limit
	}
ready:

	return &runs, nil
}
