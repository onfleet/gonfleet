package hub

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
	"github.com/onfleet/gonfleet/util"
)

type Client struct {
	apiKey       string
	rlHttpClient *netw.RlHttpClient
	url          string
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
	}
}

func (c *Client) List() ([]onfleet.Hub, error) {
	hubs := []onfleet.Hub{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &hubs)
	return hubs, err
}

func (c *Client) Create(params onfleet.HubCreateParams) (onfleet.Hub, error) {
	hub := onfleet.Hub{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPost, c.url, params, &hub)
	return hub, err
}

func (c *Client) Update(hubId string, params onfleet.HubUpdateParams) (onfleet.Hub, error) {
	hub := onfleet.Hub{}
	url := util.UrlAttachPath(c.url, hubId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPut, url, params, &hub)
	return hub, err
}
