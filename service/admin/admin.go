package admin

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
	call       caller
}

func Register(apiKey string, httpClient *http.Client, url string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
	}
}

// List fetches all admins
func (c *Client) List() ([]onfleet.Admin, error) {
	admins := []onfleet.Admin{}
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil, &admins)
	return admins, err
}

func (c *Client) Create(params onfleet.AdminCreateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, params, &admin)
	return admin, err
}

func (c *Client) Update(adminId string, params onfleet.AdminUpdateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	url := util.UrlAttachPath(c.url, adminId)
	err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, params, &admin)
	return admin, err
}

func (c *Client) Delete(adminId string) error {
	url := util.UrlAttachPath(c.url, adminId)
	err := c.call(c.apiKey, c.httpClient, http.MethodDelete, url, nil, nil)
	return err
}
