package recipient

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

// Reference https://docs.onfleet.com/reference/get-single-recipient
func (c *Client) Get(recipientId string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{recipientId},
		nil,
		nil,
		&recipient,
	)
	return recipient, err
}

// Reference https://docs.onfleet.com/reference/find-recipient
func (c *Client) Find(value string, key onfleet.RecipientQueryKey) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{string(key), value},
		nil,
		nil,
		&recipient,
	)
	return recipient, err
}

// Reference https://docs.onfleet.com/reference/update-recipient
func (c *Client) Update(recipientId string, params onfleet.RecipientUpdateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{recipientId},
		nil,
		params,
		&recipient,
	)
	return recipient, err
}

// Reference https://docs.onfleet.com/reference/create-recipient
func (c *Client) Create(params onfleet.RecipientCreateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&recipient,
	)
	return recipient, err
}
