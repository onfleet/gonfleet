package organization

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netwrk"
)

type Client struct {
	apiKey       string
	rlHttpClient *netwrk.RlHttpClient
	url          string
	altUrl       string
	call         netwrk.Caller
}

func Plug(apiKey string, rlHttpClient *netwrk.RlHttpClient, url string, altUrl string, call netwrk.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		altUrl:       altUrl,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/get-details
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

// Reference https://docs.onfleet.com/reference/get-delegatee-details
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
