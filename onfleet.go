package onfleet

import (
	"fmt"

	"github.com/onfleet/gonfleet/service/worker"
	"github.com/onfleet/gonfleet/util"
)

// user overridable defaults
const (
	defaultUserTimeout int64 = 70000
	defaultBaseUrl           = "https://onfleet.com"
	defaultPath              = "/api"
	defaultApiVersion        = "/v2"
)

type Onfleet struct {
	Workers *worker.Client
}

// InitParams accepts user provided overrides to be set on Config
type InitParams struct {
	// timeout used for http client in milliseconds
	UserTimeout int64
	BaseUrl     string
	Path        string
	ApiVersion  string
}

func New(apiKey string, params *InitParams) (*Onfleet, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("Onfleet API key not found")
	}

	o := Onfleet{}
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

	o.Workers = &worker.Client{
		ApiKey:     apiKey,
		HttpClient: httpClient,
		Url:        fullBaseUrl + "/workers",
	}
	return &o, nil
}
