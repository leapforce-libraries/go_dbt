package dbt

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Job struct {
	Id            int64    `json:"id"`
	AccountId     int64    `json:"account_id"`
	ProjectId     int64    `json:"project_id"`
	EnvironmentId int64    `json:"environment_id"`
	Name          string   `json:"name"`
	DbtVersion    string   `json:"dbt_version"`
	Triggers      Triggers `json:"triggers"`
	ExecuteSteps  []string `json:"execute_steps"`
	Settings      Settings `json:"settings"`
	State         int64    `json:"state"`
	GenerateDocs  bool     `json:"generate_docs"`
	Schedule      Schedule `json:"schedule"`
}

type Triggers struct {
	GithubWebhook    bool `json:"github_webhook"`
	Schedule         bool `json:"schedule"`
	CustomBranchOnly bool `json:"custom_branch_only"`
}

type Settings struct {
	Threads    int64  `json:"threads"`
	TargetName string `json:"target_name"`
}

type Schedule struct {
	Cron string `json:"cron"`
	Date struct {
		Type string `json:"type"`
	} `json:"date"`
	Time struct {
		Type string `json:"type"`
	} `json:"time"`
}

type GetJobsConfig struct {
	AccountId int64
	ProjectId *int64
	OrderBy   *string
}

// GetJobs returns all jobs
func (service *Service) GetJobs(config *GetJobsConfig) (*[]Job, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}
	jobs := []Job{}

	params := url.Values{}
	if config.ProjectId != nil {
		params.Add("project_id", fmt.Sprintf("%d", *config.ProjectId))
	}
	if config.OrderBy != nil {
		params.Add("order_by", *config.OrderBy)
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("accounts/%d/jobs?%s", config.AccountId, params.Encode())),
		ResponseModel: &jobs,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &jobs, nil
}

type RunJobConfig struct {
	AccountId              int64     `json:"-"`
	JobId                  int64     `json:"-"`
	Cause                  string    `json:"cause"`
	GitSha                 *string   `json:"git_sha,omitempty"`
	GitBranch              *string   `json:"git_branch,omitempty"`
	SchemaOverride         *string   `json:"schema_override,omitempty"`
	DbtVersionOverride     *string   `json:"dbt_version_override,omitempty"`
	ThreadsOverride        *int64    `json:"threads_override,omitempty"`
	TargetNameOverride     *string   `json:"target_name_override,omitempty"`
	GenerateDocsOverride   *bool     `json:"generate_docs_override,omitempty"`
	TimeoutSecondsOverride *int64    `json:"timeout_seconds_override,omitempty"`
	StepsOverride          *[]string `json:"steps_override,omitempty"`
}

type RunJobResponse struct {
	Data struct {
		Id                 int       `json:"id"`
		TriggerId          int       `json:"trigger_id"`
		AccountId          int       `json:"account_id"`
		ProjectId          int       `json:"project_id"`
		JobDefinitionId    int       `json:"job_definition_id"`
		Status             int       `json:"status"`
		GitBranch          string    `json:"git_branch"`
		GitSha             string    `json:"git_sha"`
		StatusMessage      string    `json:"status_message"`
		DbtVersion         string    `json:"dbt_version"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		DequeuedAt         time.Time `json:"dequeued_at"`
		StartedAt          time.Time `json:"started_at"`
		FinishedAt         time.Time `json:"finished_at"`
		LastCheckedAt      time.Time `json:"last_checked_at"`
		LastHeartbeatAt    time.Time `json:"last_heartbeat_at"`
		OwnerThreadId      string    `json:"owner_thread_id"`
		ExecutedByThreadId string    `json:"executed_by_thread_id"`
		ArtifactsSaved     bool      `json:"artifacts_saved"`
		ArtifactS3Path     string    `json:"artifact_s3_path"`
		HasDocsGenerated   bool      `json:"has_docs_generated"`
		Trigger            struct {
			Id                     int       `json:"id"`
			Cause                  string    `json:"cause"`
			JobDefinitionId        int       `json:"job_definition_id"`
			GitBranch              string    `json:"git_branch"`
			GitSha                 string    `json:"git_sha"`
			GithubPullRequestId    int       `json:"github_pull_request_id"`
			SchemaOverride         string    `json:"schema_override"`
			DbtVersionOverride     string    `json:"dbt_version_override"`
			ThreadsOverride        int       `json:"threads_override"`
			TargetNameOverride     string    `json:"target_name_override"`
			GenerateDocsOverride   bool      `json:"generate_docs_override"`
			TimeoutSecondsOverride int       `json:"timeout_seconds_override"`
			StepsOverride          []string  `json:"steps_override"`
			CreatedAt              time.Time `json:"created_at"`
		} `json:"trigger"`
		Job struct {
			Id            int    `json:"id"`
			AccountId     int    `json:"account_id"`
			ProjectId     int    `json:"project_id"`
			EnvironmentId int    `json:"environment_id"`
			Name          string `json:"name"`
			DbtVersion    string `json:"dbt_version"`
			Triggers      struct {
				GithubWebhook      bool `json:"github_webhook"`
				GitProviderWebhook bool `json:"git_provider_webhook"`
				Schedule           bool `json:"schedule"`
				CustomBranchOnly   bool `json:"custom_branch_only"`
			} `json:"triggers"`
			ExecuteSteps []string `json:"execute_steps"`
			Settings     struct {
				Threads    int    `json:"threads"`
				TargetName string `json:"target_name"`
			} `json:"settings"`
			State        int  `json:"state"`
			GenerateDocs bool `json:"generate_docs"`
			Schedule     struct {
				Cron string `json:"cron"`
				Date struct {
					Type string `json:"type"`
					Days []int  `json:"days"`
					Cron string `json:"cron"`
				} `json:"date"`
				Time struct {
					Type     string `json:"type"`
					Interval int    `json:"interval"`
					Hours    []int  `json:"hours"`
				} `json:"time"`
			} `json:"schedule"`
		} `json:"job"`
		Duration                string `json:"duration"`
		QueuedDuration          string `json:"queued_duration"`
		RunDuration             string `json:"run_duration"`
		DurationHumanized       string `json:"duration_humanized"`
		QueuedDurationHumanized string `json:"queued_duration_humanized"`
		RunDurationHumanized    string `json:"run_duration_humanized"`
		FinishedAtHumanized     string `json:"finished_at_humanized"`
		StatusHumanized         string `json:"status_humanized"`
		CreatedAtHumanized      string `json:"created_at_humanized"`
	} `json:"data"`
	/*Status struct {
		Code             int    `json:"code"`
		IsSuccess        bool   `json:"is_success"`
		UserMessage      string `json:"user_message"`
		DeveloperMessage string `json:"developer_message"`
	} `json:"status"`*/
}

// RunJob runs a job
func (service *Service) RunJob(config *RunJobConfig) (*RunJobResponse, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	var runJobResponse RunJobResponse

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("accounts/%d/jobs/%d/run", config.AccountId, config.JobId)),
		BodyModel:     config,
		ResponseModel: &runJobResponse,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &runJobResponse, nil
}
