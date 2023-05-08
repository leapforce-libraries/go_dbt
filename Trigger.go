package dbt

type Trigger struct {
	Id                     int         `json:"id"`
	Cause                  string      `json:"cause"`
	JobDefinitionId        int         `json:"job_definition_id"`
	GitBranch              interface{} `json:"git_branch"`
	GitSha                 interface{} `json:"git_sha"`
	AzurePullRequestId     interface{} `json:"azure_pull_request_id"`
	GithubPullRequestId    interface{} `json:"github_pull_request_id"`
	GitlabMergeRequestId   interface{} `json:"gitlab_merge_request_id"`
	SchemaOverride         interface{} `json:"schema_override"`
	DbtVersionOverride     interface{} `json:"dbt_version_override"`
	ThreadsOverride        interface{} `json:"threads_override"`
	TargetNameOverride     interface{} `json:"target_name_override"`
	GenerateDocsOverride   interface{} `json:"generate_docs_override"`
	TimeoutSecondsOverride interface{} `json:"timeout_seconds_override"`
	StepsOverride          interface{} `json:"steps_override"`
	CreatedAt              string      `json:"created_at"`
	CauseHumanized         string      `json:"cause_humanized"`
	Job                    *Job        `json:"job"`
}
