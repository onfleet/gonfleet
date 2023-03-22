package task

import (
	"net/http"

	onfleet "github.com/onfleet/gonfleet"
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

func (c *Client) Get(taskId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	url := util.UrlAttachPath(c.url, taskId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &task)
	return task, err
}

func (c *Client) GetByShortId(taskShortId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	url := util.UrlAttachPath(c.url, "shortId", taskShortId)
	err := netw.Call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &task)
	return task, err
}

func (c *Client) List() {
}
