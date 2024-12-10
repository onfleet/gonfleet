package client

import (
	"fmt"
	"time"

	"github.com/onfleet/gonfleet/netwrk"
	"github.com/onfleet/gonfleet/service/admin"
	"github.com/onfleet/gonfleet/service/container"
	"github.com/onfleet/gonfleet/service/destination"
	"github.com/onfleet/gonfleet/service/hub"
	"github.com/onfleet/gonfleet/service/organization"
	"github.com/onfleet/gonfleet/service/providers/manifest"
	"github.com/onfleet/gonfleet/service/recipient"
	"github.com/onfleet/gonfleet/service/task"
	"github.com/onfleet/gonfleet/service/team"
	"github.com/onfleet/gonfleet/service/webhook"
	"github.com/onfleet/gonfleet/service/worker"

	"golang.org/x/time/rate"
)

// user overridable defaults
const (
	defaultUserTimeout       int64 = 70000
	defaultBaseUrl                 = "https://onfleet.com"
	defaultPath                    = "/api"
	defaultApiVersion              = "/v2"
	defaultMaxCallsPerSecond       = 18
)

type API struct {
	Administrators   *admin.Client
	Containers       *container.Client
	Destinations     *destination.Client
	Hubs             *hub.Client
	Organizations    *organization.Client
	Recipients       *recipient.Client
	Tasks            *task.Client
	Teams            *team.Client
	Webhooks         *webhook.Client
	Workers          *worker.Client
	ManifestProvider *manifest.Client
}

// InitParams accepts user provided overrides to be set on Config
type InitParams struct {
	// timeout used for http client in milliseconds
	UserTimeout       int64
	BaseUrl           string
	Path              string
	ApiVersion        string
	MaxCallsPerSecond int
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
	maxCallsPerSecond := defaultMaxCallsPerSecond
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
		if params.MaxCallsPerSecond > 0 && params.MaxCallsPerSecond <= defaultMaxCallsPerSecond {
			maxCallsPerSecond = params.MaxCallsPerSecond
		}
	}

	rlHttpClient := netwrk.NewRlHttpClient(
		rate.NewLimiter(rate.Every(1*time.Second), maxCallsPerSecond),
		timeout,
	)

	fullBaseUrl := baseUrl + path + apiVersion

	api.Administrators = admin.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/admins",
		netwrk.Call,
	)
	api.Containers = container.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/containers",
		netwrk.Call,
	)
	api.Destinations = destination.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/destinations",
		netwrk.Call,
	)
	api.Hubs = hub.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/hubs",
		netwrk.Call,
	)
	api.Organizations = organization.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/organization",
		fullBaseUrl+"/organizations",
		netwrk.Call,
	)
	api.Recipients = recipient.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/recipients",
		netwrk.Call,
	)
	api.Tasks = task.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/tasks",
		netwrk.Call,
	)
	api.Teams = team.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/teams",
		netwrk.Call,
	)
	api.Webhooks = webhook.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/webhooks",
		netwrk.Call,
	)
	api.Workers = worker.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/workers",
		netwrk.Call,
	)

	// Integration Marketplace Providers
	api.ManifestProvider = manifest.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/integrations/marketplace",
		netwrk.Call,
	)

	return &api, nil
}
