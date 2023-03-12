package worker

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

type caller func(apiKey string, httpClient *http.Client, method string, url string, body any, result any) error

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
	call       caller
}

func Register(apiKey string, httpClient *http.Client, url string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
	}
}

// Get retrieves a single worker by id.
// func (c *Client) Get(workerId string) (onfleet.Worker, error) {
// 	worker := onfleet.Worker{}
// 	return worker, err
// }

// List fetches all workers.
func (c *Client) List() ([]onfleet.Worker, error) {
	workers := []onfleet.Worker{}
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil, &workers)
	return workers, err
}

// GetSchedule gets the specified worker's schedule.
func (c *Client) GetSchedule(workerId string) (onfleet.WorkerScheduleEntries, error) {
	scheduleEntries := onfleet.WorkerScheduleEntries{}
	url := util.UrlAttachPath(c.url, workerId, "schedule")
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil, &scheduleEntries)
	return scheduleEntries, err
}
