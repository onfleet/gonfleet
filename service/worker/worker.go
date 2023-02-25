package worker

import (
	"encoding/json"
	"net/http"

	"github.com/onfleet/gonfleet/types"
	"github.com/onfleet/gonfleet/util"
)

// Client for Workers resource
type Client struct {
	ApiKey     string
	HttpClient *http.Client
	Url        string
}

// List fetches all workers in organization
func (c *Client) List() ([]types.Worker, error) {
	workers := []types.Worker{}
	resp, err := util.Call(c.HttpClient, c.ApiKey, http.MethodGet, c.Url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return workers, util.ReadRequestError(resp.Body)
	}
	err = json.NewDecoder(resp.Body).Decode(&workers)
	return workers, err
}
