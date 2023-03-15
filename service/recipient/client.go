package recipient

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

func Plug(apiKey string, httpClient *http.Client, url string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
	}
}

func (c *Client) Get(recipientId string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := util.UrlAttachPath(c.url, recipientId)
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &recipient)
	return recipient, err
}

// Find searches for recipient based on provided name of phone value.
// Key options are RecipientQueryKeyName and RecipientQueryKeyPhone
//
// e.g.
//
// Find("jane doe", onfleet.RecipientQueryKeyName)
// Find("3105550100", onfleet.RecipientQueryKeyPhone)
func (c *Client) Find(value string, key onfleet.RecipientQueryKey) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := util.UrlAttachPath(c.url, string(key), value)
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &recipient)
	return recipient, err
}

// Update accepts recipientId and RecipientUpdateParams and updates the recipient.
func (c *Client) Update(recipientId string, params onfleet.RecipientUpdateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := util.UrlAttachPath(c.url, recipientId)
	err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, params, &recipient)
	return recipient, err
}

func (c *Client) Create(params onfleet.RecipientCreateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, params, &recipient)
	return recipient, err
}
