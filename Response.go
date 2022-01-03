package dbt

import "encoding/json"

// Response stores general Asana error response
//
type Response struct {
	Status struct {
		Code             int64   `json:"code"`
		IsSuccess        bool    `json:"is_success"`
		UserMessage      string  `json:"user_message"`
		DeveloperMessage *string `json:"developer_message"`
	} `json:"status"`
	Data  json.RawMessage `json:"data"`
	Extra struct {
		Filters struct {
			PkIn []int64 `json:"pk__in"`
		} `json:"filters"`
	} `json:"extra"`
	OrderBy    string `json:"order_by"`
	Pagination struct {
		Count      int64 `json:"count"`
		TotalCount int64 `json:"total_count"`
	} `json:"pagination"`
}
