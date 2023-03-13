package organization

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

type caller func(apiKey string, httpClient *http.Client, method string, url string, body any, result any) error

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
	altUrl     string
	call       caller
}

func New(apiKey string, httpClient *http.Client, url string, altUrl string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
		altUrl:     altUrl,
	}
}

func (c *Client) Get() (onfleet.Organization, error) {
	organization := onfleet.Organization{}
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil, &organization)
	return organization, err
}

func (c *Client) GetDelegate(orgId string) (onfleet.OrganizationDelegate, error) {
	delegate := onfleet.OrganizationDelegate{}
	url := util.UrlAttachPath(c.altUrl, orgId)
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &delegate)
	return delegate, err
}
