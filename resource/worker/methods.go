package worker

import (
	"encoding/json"
	"github.com/onfleet/gonfleet/util"
	"net/http"
)

// List fetches all workers in organization
func (c *Client) List() ([]Worker, error) {
	workers := []Worker{}
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
