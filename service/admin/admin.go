package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

type caller func(apiKey string, httpClient *http.Client, method string, url string, body []byte) (*http.Response, error)
type errorParser func(r io.Reader) error

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
	call       caller
	parseError errorParser
}

func Register(apiKey string, httpClient *http.Client, url string, call caller, parseError errorParser) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
		parseError: parseError,
	}
}

// List fetches all admins
func (c *Client) List() ([]onfleet.Admin, error) {
	admins := []onfleet.Admin{}
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil)
	if err != nil {
		return admins, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return admins, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&admins); err != nil {
		return admins, err
	}
	return admins, nil
}

func (c *Client) Create(params onfleet.AdminCreateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	body, err := json.Marshal(params)
	if err != nil {
		return admin, err
	}
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, body)
	if err != nil {
		return admin, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return admin, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&admin); err != nil {
		return admin, err
	}
	return admin, nil
}

func (c *Client) Update(adminId string, params onfleet.AdminUpdateParams) (onfleet.Admin, error) {
	admin := onfleet.Admin{}
	body, err := json.Marshal(params)
	if err != nil {
		return admin, err
	}
	url := fmt.Sprintf("%s/%s", c.url, adminId)
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, body)
	if err != nil {
		return admin, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return admin, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&admin); err != nil {
		return admin, err
	}
	return admin, nil
}
