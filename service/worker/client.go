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
}

func Plug(apiKey string, rlHttpClient *netw.RlHttpClient, url string) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
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
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, c.url, nil, &workers)
	return workers, err
}

// GetSchedule gets the specified worker's schedule.
func (c *Client) GetSchedule(workerId string) (onfleet.WorkerScheduleEntries, error) {
	scheduleEntries := onfleet.WorkerScheduleEntries{}
	url := util.UrlAttachPath(c.url, workerId, "schedule")
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &scheduleEntries)
	return scheduleEntries, err
}
