package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/onfleet/gonfleet/service/admin"
	"github.com/onfleet/gonfleet/service/destination"
	"github.com/onfleet/gonfleet/service/hub"
	"github.com/onfleet/gonfleet/service/organization"
	"github.com/onfleet/gonfleet/service/recipient"
	"github.com/onfleet/gonfleet/service/worker"
)

type API struct {
	Administrators *admin.Client
	Destinations   *destination.Client
	Hubs           *hub.Client
	Organizations  *organization.Client
	Recipients     *recipient.Client
	Workers        *worker.Client
}

// user overridable defaults
const (
	defaultUserTimeout int64 = 70000
	defaultBaseUrl           = "https://onfleet.com"
	defaultPath              = "/api"
	defaultApiVersion        = "/v2"
)

const (
	pkgName    = "onfleet/gonfleet"
	pkgVersion = "1.0.0"
)

// InitParams accepts user provided overrides to be set on Config
type InitParams struct {
	// timeout used for http client in milliseconds
	UserTimeout int64
	BaseUrl     string
	Path        string
	ApiVersion  string
}

type requestErrorMessage struct {
	Cause any `json:"cause,omitempty"`
	// Error is an internal error code.
	// It is different than the request status code.
	Error int `json:"error"`
	// Message is the error messages / description
	Message string `json:"message"`
	// RemoteAddress is remote ip
	RemoteAddress string `json:"remoteAddress"`
	// Request is uuid associated with the request
	Request string `json:"request"`
}

type requestError struct {
	// Code is error type e.g. "InvalidArgument"
	Code string `json:"code"`
	// Message contains futher details about the error.
	Message requestErrorMessage `json:"message"`
}

func (err requestError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message.Message)
}

func parseError(r io.Reader) error {
	var reqError requestError
	if err := json.NewDecoder(r).Decode(&reqError); err != nil {
		return err
	}
	return reqError
}

func call(apiKey string, httpClient *http.Client, method string, url string, body any, result any) error {
	var request *http.Request
	var err error
	switch method {
	case "GET", "DELETE":
		request, err = http.NewRequest(
			method,
			url,
			nil,
		)
		request.Header.Set("Accept", "application/json")
	case "POST", "PUT":
		body, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buffer := bytes.NewBuffer(body)
		request, err = http.NewRequest(
			method,
			url,
			buffer,
		)
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", fmt.Sprintf("%s-%s", pkgName, pkgVersion))
	request.SetBasicAuth(apiKey, "")
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return parseError(response.Body)
	}
	if result == nil {
		return nil
	}
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return err
	}
	return nil
}

func New(apiKey string, params *InitParams) (*API, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("Onfleet API key not found")
	}

	api := API{}
	baseUrl := defaultBaseUrl
	path := defaultPath
	apiVersion := defaultApiVersion
	timeout := defaultUserTimeout

	if params != nil {
		if params.BaseUrl != "" {
			baseUrl = params.BaseUrl
		}
		if params.Path != "" {
			path = params.Path
		}
		if params.ApiVersion != "" {
			apiVersion = params.ApiVersion
		}
		if params.UserTimeout > 0 && params.UserTimeout <= defaultUserTimeout {
			timeout = params.UserTimeout
		}
	}

	httpClient := &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	fullBaseUrl := baseUrl + path + apiVersion

	api.Administrators = admin.New(
		apiKey,
		httpClient,
		fullBaseUrl+"/admins",
		call,
	)
	api.Destinations = destination.New(
		apiKey,
		httpClient,
		fullBaseUrl+"/destinations",
		call,
	)
	api.Hubs = hub.New(
		apiKey,
		httpClient,
		fullBaseUrl+"/hubs",
		call,
	)
	api.Organizations = organization.New(
		apiKey,
		httpClient,
		fullBaseUrl+"/organization",
		fullBaseUrl+"/organizations",
		call,
	)
	api.Recipients = recipient.New(
		apiKey,
		httpClient,
		fullBaseUrl+"/recipients",
		call,
	)
	api.Workers = worker.New(
		apiKey,
		httpClient,
		fullBaseUrl+"/workers",
		call,
	)

	return &api, nil
}
