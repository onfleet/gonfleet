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

func (c *Client) ListWorkersByLocation(params onfleet.WorkersByLocationListQueryParams) (onfleet.WorkersByLocation, error) {
	workersByLocation := onfleet.WorkersByLocation{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{"location"},
		params,
		nil,
		&workersByLocation,
	)
	return workersByLocation, err
}

func (c *Client) SetSchedule(workerId string, entries onfleet.WorkerScheduleEntries) (onfleet.WorkerScheduleEntries, error) {
	scheduleEntries := onfleet.WorkerScheduleEntries{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{workerId, "schedule"},
		nil,
		entries,
		&scheduleEntries,
	)
	return scheduleEntries, err
}

// Lists tasks assigned to specified worker.
// Params must contain a valid "From" millisecond timestamp.
func (c *Client) ListTasks(workerId string, params onfleet.WorkerTasksListQueryParams) (onfleet.WorkerTasks, error) {
	workerTasks := onfleet.WorkerTasks{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{workerId, "tasks"},
		params,
		nil,
		&workerTasks,
	)
	return workerTasks, err
}

// Creates new worker.
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

// Updates worker with specified id.
func (c *Client) Update(workerId string, params onfleet.WorkerUpdateParams) (onfleet.Worker, error) {
	worker := onfleet.Worker{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{workerId},
		nil,
		params,
		&worker,
	)
	return worker, err
}

// Deletes worker with specified id.
func (c *Client) Delete(workerId string) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodDelete,
		c.url,
		[]string{workerId},
		nil,
		nil,
		nil,
	)
	return err
}
