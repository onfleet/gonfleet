package worker

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netw"
	"github.com/onfleet/gonfleet/util"
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

func (c *Client) Get(workerId string, params *onfleet.WorkerGetParams) (onfleet.Worker, error) {
	worker := onfleet.Worker{}
	url := util.UrlAttachPath(c.url, workerId)
	url = util.UrlAttachQuery(url, params)
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &worker)
	return worker, err
}

// func (c *Client) GetWithQuery(workerId string, params onfleet.WorkerGetParams) (onfleet.Worker, error) {
// }

// List fetches all workers.
func (c *Client) List() ([]onfleet.Worker, error) {
	workers := []onfleet.Worker{}
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &workers)
	return workers, err
}

// func (c *Client) ListWithQuery() {
//
// }

// GetSchedule gets the specified worker's schedule.
func (c *Client) GetSchedule(workerId string) (onfleet.WorkerScheduleEntries, error) {
	scheduleEntries := onfleet.WorkerScheduleEntries{}
	url := util.UrlAttachPath(c.url, workerId, "schedule")
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &scheduleEntries)
	return scheduleEntries, err
}
