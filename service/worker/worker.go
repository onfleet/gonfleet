package worker

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

// List fetches all workers
func (c *Client) List() ([]onfleet.Worker, error) {
	workers := []onfleet.Worker{}
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil)
	if err != nil {
		return workers, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return workers, c.parseError(resp.Body)
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
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return scheduleEntries, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return scheduleEntries, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&scheduleEntries); err != nil {
		return scheduleEntries, err
	}
	return scheduleEntries, nil
}
