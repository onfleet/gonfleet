package worker

import (
	"encoding/json"
	"net/http"

	"github.com/onfleet/gonfleet/util"
)

// List fetches all workers in organization
func (c *Client) List() ([]Worker, error) {
	workers := []Worker{}
	req, err := util.NewHttpRequest(
		c.ApiKey,
		http.MethodGet,
		c.Url,
		nil,
	)
	if err != nil {
		return workers, err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return workers, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return workers, util.ReadRequestError(resp.Body)
	}
	err = json.NewDecoder(resp.Body).Decode(&workers)
	return workers, err
}
