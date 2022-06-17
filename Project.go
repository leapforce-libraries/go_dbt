package dbt

import (
	"fmt"
	"net/http"

	d_types "github.com/leapforce-libraries/go_dbt/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Project struct {
	Id                     int64                  `json:"id"`
	AccountId              int64                  `json:"account_id"`
	Connection             Connection             `json:"connection"`
	ConnectionId           int64                  `json:"connection_id"`
	DbtProjectSubdirectory string                 `json:"dbt_project_subdirectory"`
	Name                   string                 `json:"name"`
	Repository             Repository             `json:"repository"`
	RepositoryId           int64                  `json:"repository_id"`
	State                  int64                  `json:"state"`
	CreatedAt              d_types.DateTimeString `json:"created_at"`
	UpdatedAt              d_types.DateTimeString `json:"updated_at"`
}

type Connection struct {
	Id                      int64                  `json:"id"`
	AccountId               int64                  `json:"account_id"`
	Name                    string                 `json:"name"`
	Type                    string                 `json:"type"`
	CreatedById             int64                  `json:"created_by_id"`
	CreatedByServiceTokenId int64                  `json:"created_by_service_token_id"`
	State                   int64                  `json:"state"`
	CreatedAt               d_types.DateTimeString `json:"created_at"`
	UpdatedAt               d_types.DateTimeString `json:"updated_at"`
}

type Repository struct {
	Id                   int64                  `json:"id"`
	AccountId            int64                  `json:"account_id"`
	RemoteUrl            string                 `json:"remote_url"`
	RemoteBackend        string                 `json:"remote_backend"`
	GitCloneStrategy     string                 `json:"git_clone_strategy"`
	DeployKeyId          int64                  `json:"deploy_key_id"`
	GithubInstallationId int64                  `json:"github_installation_id"`
	GithubRepo           string                 `json:"github_repo"`
	Name                 string                 `json:"name"`
	FullName             string                 `json:"full_name"`
	State                int64                  `json:"state"`
	CreatedAt            d_types.DateTimeString `json:"created_at"`
	UpdatedAt            d_types.DateTimeString `json:"updated_at"`
}

type GetProjectsConfig struct {
	AccountId int64
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
		Url:           service.url(fmt.Sprintf("accounts/%d/projects", config.AccountId)),
		ResponseModel: &projects,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &projects, nil
}
