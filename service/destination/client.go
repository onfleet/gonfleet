package destination

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

func (c *Client) Get(destinationId string) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	url := util.UrlAttachPath(c.url, destinationId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &destination)
	return destination, err
}

func (c *Client) Create(params onfleet.DestinationCreateParams) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPost, c.url, params, &destination)
	return destination, err
}
