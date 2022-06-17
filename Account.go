package dbt

import (
	"net/http"

	d_types "github.com/leapforce-libraries/go_dbt/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Account struct {
	Id             int64                  `json:"id"`
	Name           string                 `json:"name"`
	Plan           string                 `json:"plan"`
	PendingCancel  bool                   `json:"pending_cancel"`
	State          int64                  `json:"state"`
	DeveloperSeats int64                  `json:"developer_seats"`
	ReadOnlySeats  int64                  `json:"read_only_seats"`
	RunSlots       int64                  `json:"run_slots"`
	CreatedAt      d_types.DateTimeString `json:"created_at"`
	UpdatedAt      d_types.DateTimeString `json:"updated_at"`
}

type GetAccountsConfig struct {
}

// GetAccounts returns all accounts
//
func (service *Service) GetAccounts(config *GetAccountsConfig) (*[]Account, *errortools.Error) {
	accounts := []Account{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("accounts"),
		ResponseModel: &accounts,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &accounts, nil
}
