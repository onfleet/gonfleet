package webhook

import (
	"net/http"

	onfleet "github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
	"github.com/onfleet/gonfleet/util"
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

func (c *Client) List() ([]onfleet.Webhook, error) {
	webhooks := []onfleet.Webhook{}
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &webhooks)
	return webhooks, err
}

func (c *Client) Create(params onfleet.WebhookCreateParams) (onfleet.Webhook, error) {
	webhook := onfleet.Webhook{}
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodPost, c.url, params, &webhook)
	return webhook, err
}

func (c *Client) Delete(webhookId string) error {
	url := util.UrlAttachPath(c.url, webhookId)
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodDelete, url, nil, nil)
	return err
}
