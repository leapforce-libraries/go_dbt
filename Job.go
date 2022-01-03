package dbt

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Job struct {
	AccountID     int64    `json:"account_id"`
	ProjectID     int64    `json:"project_id"`
	EnvironmentID int64    `json:"environment_id"`
	Name          string   `json:"name"`
	DBTVersion    string   `json:"dbt_version"`
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
	AccountID int64
	ProjectID *int64
	OrderBy   *string
}

// GetJobs returns all jobs
//
func (service *Service) GetJobs(config *GetJobsConfig) (*[]Job, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}
	jobs := []Job{}

	params := url.Values{}
	if config.ProjectID != nil {
		params.Add("project_id", fmt.Sprintf("%d", *config.ProjectID))
	}
	if config.OrderBy != nil {
		params.Add("order_by", *config.OrderBy)
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(fmt.Sprintf("accounts/%d/jobs?%s", config.AccountID, params.Encode())),
		ResponseModel: &jobs,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &jobs, nil
}
