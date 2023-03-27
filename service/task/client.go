package task

import (
	"net/http"

	onfleet "github.com/onfleet/gonfleet"
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

func (c *Client) Get(taskId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{taskId},
		nil,
		nil,
		&task,
	)
	return task, err
}

func (c *Client) GetByShortId(taskShortId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{"shortId", taskShortId},
		nil,
		nil,
		&task,
	)
	return task, err
}

func (c *Client) ListTasks(teamId string) {
}
