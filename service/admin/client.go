package admin

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

// List fetches all admins.
func (c *Client) List() ([]onfleet.Admin, error) {
	admins := []onfleet.Admin{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &admins)
	return admins, err
}

// Creates a new administrator.
func (c *Client) Create(params onfleet.AdminCreateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPost, c.url, params, &admin)
	return admin, err
}

// Updates an administrator.
// If updating email address further email verification will be required before change is processed.
func (c *Client) Update(adminId string, params onfleet.AdminUpdateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	url := util.UrlAttachPath(c.url, adminId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodPut, url, params, &admin)
	return admin, err
}

// Deletes an admin
func (c *Client) Delete(adminId string) error {
	url := util.UrlAttachPath(c.url, adminId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodDelete, url, nil, nil)
	return err
}
