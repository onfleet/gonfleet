package recipient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/onfleet/gonfleet"
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

func (c *Client) Get(recipientId string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := fmt.Sprintf("%s/%s", c.url, recipientId)
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &recipient)
	return recipient, err
}

// Find searches for recipient based on provided name of phone value.
// Key options are "name" and "phone".
//
// e.g.
//
// Find("jane doe", "name")
// Find("3105550100", "phone")
func (c *Client) Find(value string, key string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := fmt.Sprintf("%s/%s/%s", c.url, key, url.PathEscape(value))
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &recipient)
	return recipient, err
}

func (c *Client) Update(recipientId string, params onfleet.RecipientUpdateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := fmt.Sprintf("%s/%s", c.url, recipientId)
	err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, params, &recipient)
	return recipient, err
}

func (c *Client) Create(params onfleet.RecipientCreateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := c.url
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, url, params, &recipient)
	return recipient, err
}
