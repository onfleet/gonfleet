package hub

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
)

type Client struct {
	apiKey       string
	rlHttpClient *netw.RlHttpClient
	url          string
	call         netw.Caller
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string, call netw.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/list-hubs
func (c *Client) List() ([]onfleet.Hub, error) {
	hubs := []onfleet.Hub{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		nil,
		nil,
		&hubs,
	)
	return hubs, err
}

// Reference https://docs.onfleet.com/reference/create-hub
func (c *Client) Create(params onfleet.HubCreateParams) (onfleet.Hub, error) {
	hub := onfleet.Hub{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&hub,
	)
	return hub, err
}

// Reference https://docs.onfleet.com/reference/update-hub
func (c *Client) Update(hubId string, params onfleet.HubUpdateParams) (onfleet.Hub, error) {
	hub := onfleet.Hub{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{hubId},
		nil,
		params,
		&hub,
	)
	return hub, err
}
