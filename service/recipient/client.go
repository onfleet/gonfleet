package recipient

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
	"github.com/onfleet/gonfleet/util"
)

type Client struct {
	apiKey       string
	rlHttpClient *netw.RlHttpClient
	url          string
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
	}
}

func (c *Client) Get(recipientId string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := util.UrlAttachPath(c.url, recipientId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &recipient)
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
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &recipient)
	return recipient, err
}

// Update accepts recipientId and RecipientUpdateParams and updates the recipient.
func (c *Client) Update(recipientId string, params onfleet.RecipientUpdateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := util.UrlAttachPath(c.url, recipientId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPut, url, params, &recipient)
	return recipient, err
}

func (c *Client) Create(params onfleet.RecipientCreateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPost, c.url, params, &recipient)
	return recipient, err
}
