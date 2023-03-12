package admin

import (
	"encoding/json"
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
