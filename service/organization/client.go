package organization

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
)

type Client struct {
	apiKey       string
	rlHttpClient *netw.RlHttpClient
	url          string
	altUrl       string
	call         netw.Caller
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string, altUrl string, call netw.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		altUrl:       altUrl,
		call:         call,
	}
}

func (c *Client) Get() (onfleet.Organization, error) {
	organization := onfleet.Organization{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		nil,
		nil,
		&organization,
	)
	return organization, err
}

func (c *Client) GetDelegate(orgId string) (onfleet.OrganizationDelegate, error) {
	delegate := onfleet.OrganizationDelegate{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.altUrl,
		[]string{orgId},
		nil,
		nil,
		&delegate,
	)
	return delegate, err
}
