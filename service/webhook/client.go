package webhook

import (
	"net/http"

	onfleet "github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
	"github.com/onfleet/gonfleet/util"
)


type Client struct {
	apiKey     string
	rlHttpClient *netw.RlHttpClient
	url        string
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string) *Client {
	return &Client{
		apiKey:     apiKey,
		rlHttpClient: rlHttpClient,
		url:        url,
	}
}

func (c *Client) List() ([]onfleet.Webhook, error) {
	webhooks := []onfleet.Webhook{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &webhooks)
	return webhooks, err
}

func (c *Client) Create(params onfleet.WebhookCreateParams) (onfleet.Webhook, error) {
	webhook := onfleet.Webhook{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPost, c.url, params, &webhook)
	return webhook, err
}

func (c *Client) Delete(webhookId string) error {
	url := util.UrlAttachPath(c.url, webhookId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodDelete, url, nil, nil)
	return err
}
