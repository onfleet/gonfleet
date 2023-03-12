package destination

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

func Register(apiKey string, httpClient *http.Client, url string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
	}
}

func (c *Client) Get(destinationId string) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	url := util.UrlAttachPath(c.url, destinationId)
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &destination)
	return destination, err
}

func (c *Client) Create(params onfleet.DestinationCreateParams) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, params, &destination)
	return destination, err
}
