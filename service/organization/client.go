package organization

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
	altUrl       string
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string, altUrl string) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		altUrl:       altUrl,
	}
}

func (c *Client) Get() (onfleet.Organization, error) {
	organization := onfleet.Organization{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &organization)
	return organization, err
}

func (c *Client) GetDelegate(orgId string) (onfleet.OrganizationDelegate, error) {
	delegate := onfleet.OrganizationDelegate{}
	url := util.UrlAttachPath(c.altUrl, orgId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &delegate)
	return delegate, err
}
