package dbt

import (
	"encoding/json"
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	apiName string = "dbt"
	apiUrl  string = "https://cloud.getdbt.com/api/v2"
	//DateTimeLayout  string = "2006-01-02T15:04:05.000Z"
	//defaultPageSize int64  = 100
)

// type
//
type Service struct {
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	ApiKey string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ApiKey == "" {
		return nil, errortools.ErrorMessage("Service ApiKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiKey:      serviceConfig.ApiKey,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Token %s", service.apiKey))
	(*requestConfig).NonDefaultHeaders = &header

	errorResponse := Response{}
	if utilities.IsNil(requestConfig.ErrorModel) {
		// add error model
		(*requestConfig).ErrorModel = &errorResponse
	}

	responseModel := requestConfig.ResponseModel
	_response := Response{}
	requestConfig.ResponseModel = &_response

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if e != nil {
		if errorResponse.Status.UserMessage != "" {
			e.SetMessage(errorResponse.Status.UserMessage)
		}
	} else {
		err := json.Unmarshal(_response.Data, responseModel)
		if err != nil {
			return request, response, errortools.ErrorMessage(err)
		}
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
