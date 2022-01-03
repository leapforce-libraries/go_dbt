package dbt

import (
	"encoding/json"
	"fmt"
	"net/http"

	a_types "github.com/leapforce-libraries/go_airtable/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
	Offset   string    `json:"offset"`
}

type Account struct {
	ID          string                     `json:"id"`
	Fields      map[string]json.RawMessage `json:"fields"`
	CreatedTime a_types.DateTimeString     `json:"createdTime"`
}

type GetAccountsConfig struct {
	Fields          *[]string
	FilterByFormula *string
	MaxAccounts     *int64
	PageSize        *int64
	Sort            *[]struct {
		Field     string
		Direction string
	}
	View       *string
	CellFormat *string
	TimeZone   *string
	UserLocale *string
}

// GetAccounts returns all accounts
//
func (service *Service) GetAccounts(config *GetAccountsConfig) (*[]Account, *errortools.Error) {
	accounts := []Account{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url("accounts"),
		ResponseModel: &accounts,
	}
	fmt.Println(requestConfig.URL)
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &accounts, nil
}
