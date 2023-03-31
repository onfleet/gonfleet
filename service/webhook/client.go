package webhook

import (
    "net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netwrk"
)

type Client struct {
	apiKey       string
	rlHttpClient *netwrk.RlHttpClient
	url          string
	call         netwrk.Caller
}

func Plug(apiKey string, rlHttpClient *netwrk.RlHttpClient, url string, call netwrk.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/list-webhooks
func (c *Client) List() ([]onfleet.Webhook, error) {
	webhooks := []onfleet.Webhook{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		nil,
		nil,
		&webhooks,
	)
	return webhooks, err
}

// Reference https://docs.onfleet.com/reference/create-webhook
func (c *Client) Create(params onfleet.WebhookCreateParams) (onfleet.Webhook, error) {
	webhook := onfleet.Webhook{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&webhook,
	)
	return webhook, err
}

// Reference https://docs.onfleet.com/reference/delete-webhook
func (c *Client) Delete(webhookId string) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodDelete,
		c.url,
		[]string{webhookId},
		nil,
		nil,
		nil,
	)
	return err
}
