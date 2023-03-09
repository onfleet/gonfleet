package client

import (
	"fmt"

	"github.com/onfleet/gonfleet/service/destination"
	"github.com/onfleet/gonfleet/service/worker"

	"github.com/onfleet/gonfleet/util"
)

type API struct {
	Destinations *destination.Client
	Workers      *worker.Client
}

// user overridable defaults
const (
	defaultUserTimeout int64 = 70000
	defaultBaseUrl           = "https://onfleet.com"
	defaultPath              = "/api"
	defaultApiVersion        = "/v2"
)

// InitParams accepts user provided overrides to be set on Config
type InitParams struct {
	// timeout used for http client in milliseconds
	UserTimeout int64
	BaseUrl     string
	Path        string
	ApiVersion  string
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

	httpClient := util.NewHttpClient(timeout)
	fullBaseUrl := baseUrl + path + apiVersion

	api.Workers = worker.Register(apiKey, httpClient, fullBaseUrl+"/workers")
	api.Destinations = destination.Register(apiKey, httpClient, fullBaseUrl+"/destinations")

	return &api, nil
}
