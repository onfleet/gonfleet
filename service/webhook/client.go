package webhook

import (
	"net/http"

	onfleet "github.com/onfleet/gonfleet"
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

func (c *Client) List() ([]onfleet.Webhook, error) {
	webhooks := []onfleet.Webhook{}
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil, &webhooks)
	return webhooks, err
}

func (c *Client) Create(params onfleet.WebhookCreateParams) (onfleet.Webhook, error) {
	webhook := onfleet.Webhook{}
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, params, &webhook)
	return webhook, err
}

func (c *Client) Delete(webhookId string) error {
	url := util.UrlAttachPath(c.url, webhookId)
	err := c.call(c.apiKey, c.httpClient, http.MethodDelete, url, nil, nil)
	return err
}
