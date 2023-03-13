package hub

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

type caller func(apiKey string, httpClient *http.Client, method string, url string, body any, result any) error

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
	call       caller
}

func New(apiKey string, httpClient *http.Client, url string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
	}
}

func (c *Client) List() ([]onfleet.Hub, error) {
	hubs := []onfleet.Hub{}
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil, &hubs)
	return hubs, err
}

func (c *Client) Create(params onfleet.HubCreateParams) (onfleet.Hub, error) {
	hub := onfleet.Hub{}
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, params, &hub)
	return hub, err
}

func (c *Client) Update(hubId string, params onfleet.HubUpdateParams) (onfleet.Hub, error) {
	hub := onfleet.Hub{}
	url := util.UrlAttachPath(c.url, hubId)
	err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, params, &hub)
	return hub, err
}
