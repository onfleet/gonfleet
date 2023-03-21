package client

import (
	"fmt"
	"time"

	"github.com/onfleet/gonfleet/constants"
	"github.com/onfleet/gonfleet/netw"
	"github.com/onfleet/gonfleet/service/admin"
	"github.com/onfleet/gonfleet/service/container"
	"github.com/onfleet/gonfleet/service/destination"
	"github.com/onfleet/gonfleet/service/hub"
	"github.com/onfleet/gonfleet/service/organization"
	"github.com/onfleet/gonfleet/service/recipient"
	"github.com/onfleet/gonfleet/service/team"
	"github.com/onfleet/gonfleet/service/webhook"
	"github.com/onfleet/gonfleet/service/worker"

	"golang.org/x/time/rate"
)

type API struct {
	Administrators *admin.Client
	Containers     *container.Client
	Destinations   *destination.Client
	Hubs           *hub.Client
	Organizations  *organization.Client
	Recipients     *recipient.Client
	Teams          *team.Client
	Webhooks       *webhook.Client
	Workers        *worker.Client
}

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
	baseUrl := constants.DefaultBaseUrl
	path := constants.DefaultPath
	apiVersion := constants.DefaultApiVersion
	timeout := constants.DefaultUserTimeout

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
		if params.UserTimeout > 0 && params.UserTimeout <= constants.DefaultUserTimeout {
			timeout = params.UserTimeout
		}
	}

	rlHttpClient := netw.NewRlHttpClient(
		rate.NewLimiter(rate.Every(1*time.Second), 18),
		timeout,
	)

	fullBaseUrl := baseUrl + path + apiVersion

	api.Administrators = admin.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/admins",
	)
	api.Containers = container.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/containers",
	)
	api.Destinations = destination.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/destinations",
	)
	api.Hubs = hub.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/hubs",
	)
	api.Organizations = organization.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/organization",
		fullBaseUrl+"/organizations",
	)
	api.Recipients = recipient.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/recipients",
	)
	api.Teams = team.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/teams",
	)
	api.Webhooks = webhook.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/webhooks",
	)
	api.Workers = worker.Plug(
		apiKey,
		rlHttpClient,
		fullBaseUrl+"/workers",
	)

	return &api, nil
}
