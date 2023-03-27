package destination

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

// Reference https://docs.onfleet.com/reference/get-single-destination
func (c *Client) Get(destinationId string) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{destinationId},
		nil,
		nil,
		&destination,
	)
	return destination, err
}

// Reference https://docs.onfleet.com/reference/create-destination
func (c *Client) Create(params onfleet.DestinationCreateParams) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&destination,
	)
	return destination, err
}
