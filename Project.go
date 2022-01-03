package dbt

import (
	"fmt"
	"net/http"

	d_types "github.com/leapforce-libraries/go_dbt/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Project struct {
	ID                     int64                  `json:"id"`
	AccountID              int64                  `json:"account_id"`
	Connection             Connection             `json:"connection"`
	ConnectionID           int64                  `json:"connection_id"`
	DBTProjectSubdirectory string                 `json:"dbt_project_subdirectory"`
	Name                   string                 `json:"name"`
	Repository             Repository             `json:"repository"`
	RepositoryID           int64                  `json:"repository_id"`
	State                  int64                  `json:"state"`
	CreatedAt              d_types.DateTimeString `json:"created_at"`
	UpdatedAt              d_types.DateTimeString `json:"updated_at"`
}

type Connection struct {
	ID                      int64                  `json:"id"`
	AccountID               int64                  `json:"account_id"`
	Name                    string                 `json:"name"`
	Type                    string                 `json:"type"`
	CreatedByID             int64                  `json:"created_by_id"`
	CreatedByServiceTokenID int64                  `json:"created_by_service_token_id"`
	State                   int64                  `json:"state"`
	CreatedAt               d_types.DateTimeString `json:"created_at"`
	UpdatedAt               d_types.DateTimeString `json:"updated_at"`
}

type Repository struct {
	ID                   int64                  `json:"id"`
	AccountID            int64                  `json:"account_id"`
	RemoteURL            string                 `json:"remote_url"`
	RemoteBackend        string                 `json:"remote_backend"`
	GitCloneStrategy     string                 `json:"git_clone_strategy"`
	DeployKeyID          int64                  `json:"deploy_key_id"`
	GithubInstallationID int64                  `json:"github_installation_id"`
	GithubRepo           string                 `json:"github_repo"`
	Name                 string                 `json:"name"`
	FullName             string                 `json:"full_name"`
	State                int64                  `json:"state"`
	CreatedAt            d_types.DateTimeString `json:"created_at"`
	UpdatedAt            d_types.DateTimeString `json:"updated_at"`
}

type GetProjectsConfig struct {
	AccountID int64
}

// GetProjects returns all projects
//
func (service *Service) GetProjects(config *GetProjectsConfig) (*[]Project, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}
	projects := []Project{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(fmt.Sprintf("accounts/%d/projects", config.AccountID)),
		ResponseModel: &projects,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &projects, nil
}
