package container

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

func (c *Client) Get(id string, key onfleet.ContainerQueryKey) (onfleet.Container, error) {
	container := onfleet.Container{}
	url := util.UrlAttachPath(c.url, string(key), id)
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &container)
	return container, err
}
