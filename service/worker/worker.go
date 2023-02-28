package worker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onfleet/gonfleet/resource"
	"github.com/onfleet/gonfleet/util"
)

// Client for Workers resource
type Client struct {
	ApiKey     string
	HttpClient *http.Client
	Url        string
}

// List fetches all workers
func (c *Client) List() ([]resource.Worker, error) {
	workers := []resource.Worker{}
	resp, err := util.Call(c.HttpClient, c.ApiKey, http.MethodGet, c.Url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return nil, util.ReadRequestError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&workers); err != nil {
		return nil, err
	}
	return workers, nil
}

// GetSchedule gets the specified worker's schedule
func (c *Client) GetSchedule(workerId string) (resource.WorkerScheduleEntries, error) {
	var scheduleEntries resource.WorkerScheduleEntries
	url := fmt.Sprintf("%s/%s/schedule", c.Url, workerId)
	resp, err := util.Call(c.HttpClient, c.ApiKey, http.MethodGet, url, nil)
	if err != nil {
		return scheduleEntries, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return scheduleEntries, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&scheduleEntries); err != nil {
		return scheduleEntries, err
	}
	return scheduleEntries, nil
}
