package destination

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

// Client for Workers resource
type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
}

func Register(apiKey string, httpClient *http.Client, url string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
	}
}

func (c *Client) Get(destinationId string) (onfleet.Destination, error) {
	destination := onfleet.Destination{}
	url := fmt.Sprintf("%s/%s", c.url, destinationId)
	resp, err := util.Call(c.httpClient, c.apiKey, http.MethodGet, url, nil)
	if err != nil {
		return destination, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return destination, util.ReadRequestError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&destination); err != nil {
		return destination, err
	}
	return destination, nil
}
