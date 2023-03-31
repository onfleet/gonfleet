package admin

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netwrk"
)

type Client struct {
	apiKey       string
	rlHttpClient *netwrk.RlHttpClient
	url          string
	call         netwrk.Caller
}

func Plug(apiKey string, rlHttpClient *netwrk.RlHttpClient, url string, call netwrk.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/list-administrators
func (c *Client) List() ([]onfleet.Admin, error) {
	admins := []onfleet.Admin{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		nil,
		nil,
		&admins,
	)
	return admins, err
}

// Reference https://docs.onfleet.com/reference/querying-by-metadata
func (c *Client) ListWithMetadataQuery(metadata []onfleet.Metadata) ([]onfleet.Admin, error) {
	admins := []onfleet.Admin{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{"metadata"},
		nil,
		metadata,
		&admins,
	)
	return admins, err
}

// Reference https://docs.onfleet.com/reference/create-administrator
func (c *Client) Create(params onfleet.AdminCreateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&admin,
	)
	return admin, err
}

// Reference https://docs.onfleet.com/reference/update-administrator
func (c *Client) Update(adminId string, params onfleet.AdminUpdateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		nil,
		nil,
		params,
		&admin,
	)
	return admin, err
}

// Reference https://docs.onfleet.com/reference/delete-administrator
func (c *Client) Delete(adminId string) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodDelete,
		c.url,
		[]string{adminId},
		nil,
		nil,
		nil,
	)
	return err
}
