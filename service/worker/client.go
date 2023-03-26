package worker

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
)

type Client struct {
	apiKey       string
	rlHttpClient *netw.RlHttpClient
	url          string
	call         netw.Caller
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string, call netw.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Get gets a single worker.
func (c *Client) Get(workerId string) (onfleet.Worker, error) {
	worker := onfleet.Worker{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{workerId},
		nil,
		nil,
		&worker,
	)
	return worker, err
}

// GetWithQuery gets a single worker with query params.
func (c *Client) GetWithQuery(workerId string, params onfleet.WorkerGetQueryParams) (map[string]any, error) {
	worker := map[string]any{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{workerId},
		params,
		nil,
		&worker,
	)
	return worker, err
}

// List fetches all workers.
func (c *Client) List() ([]onfleet.Worker, error) {
	workers := []onfleet.Worker{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		nil,
		nil,
		&workers,
	)
	return workers, err
}

// List fetches all workers with specified query param.
func (c *Client) ListWithQuery(params onfleet.WorkerListQueryParams) ([]map[string]any, error) {
	workers := []map[string]any{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		params,
		nil,
		&workers,
	)
	return workers, err
}

// GetSchedule gets the specified worker's schedule.
func (c *Client) GetSchedule(workerId string) (onfleet.WorkerScheduleEntries, error) {
	scheduleEntries := onfleet.WorkerScheduleEntries{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{workerId, "schedule"},
		nil,
		nil,
		&scheduleEntries,
	)
	return scheduleEntries, err
}

// Creates creates a new Onfleet worker.
func (c *Client) Create(params onfleet.WorkerCreateParams) (onfleet.Worker, error) {
	worker := onfleet.Worker{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&worker,
	)
	return worker, err
}
