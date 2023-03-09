package worker

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

// List fetches all workers
func (c *Client) List() ([]onfleet.Worker, error) {
	workers := []onfleet.Worker{}
	resp, err := util.Call(c.httpClient, c.apiKey, http.MethodGet, c.url, nil)
	if err != nil {
		return workers, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return workers, util.ReadRequestError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&workers); err != nil {
		return workers, err
	}
	return workers, nil
}

// GetSchedule gets the specified worker's schedule
func (c *Client) GetSchedule(workerId string) (onfleet.WorkerScheduleEntries, error) {
	scheduleEntries := onfleet.WorkerScheduleEntries{}
	url := fmt.Sprintf("%s/%s/schedule", c.url, workerId)
	resp, err := util.Call(c.httpClient, c.apiKey, http.MethodGet, url, nil)
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
