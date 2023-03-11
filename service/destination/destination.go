package destination

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

// Client for Workers resource
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

func (c *Client) Get(destinationId string) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	url := fmt.Sprintf("%s/%s", c.url, destinationId)
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return destination, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return destination, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&destination); err != nil {
		return destination, err
	}
	return destination, nil
}
