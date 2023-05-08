package dbt

type Environment struct {
	DbtProjectSubdirectory string `json:"dbt_project_subdirectory"`
	ProjectId              int64  `json:"project_id"`
	Id                     int64  `json:"id"`
	AccountId              int64  `json:"account_id"`
	ConnectionId           int64  `json:"connection_id"`
	RepositoryId           int64  `json:"repository_id"`
	CredentialsId          int64  `json:"credentials_id"`
	CreatedById            int64  `json:"created_by_id"`
	Name                   string `json:"name"`
	UseCustomBranch        bool   `json:"use_custom_branch"`
	CustomBranch           string `json:"custom_branch"`
	DbtVersion             string `json:"dbt_version"`
	RawDbtVersion          string `json:"raw_dbt_version"`
	SupportsDocs           bool   `json:"supports_docs"`
	State                  int64  `json:"state"`
}
