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

func (c *Client) Get(taskId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	url := util.UrlAttachPath(c.url, taskId)
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &task)
	return task, err
}

func (c *Client) GetByShortId(taskShortId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	url := util.UrlAttachPath(c.url, "shortId", taskShortId)
	err := c.call(c.apiKey, c.rlHttpClient, http.MethodGet, url, nil, &task)
	return task, err
}

func (c *Client) List() {
}
