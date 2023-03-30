package container

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netwrk"
)

type Client struct {
	apiKey       string
	rlHttpClient *netwrk.RlHttpClient
	url          string
	call         netwrk.Caller
}

func Plug(apiKey string, rlHttpClient *netwrk.RlHttpClient, url string, call netwrk.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/get-container
func (c *Client) Get(id string, key onfleet.ContainerQueryKey) (onfleet.Container, error) {
	container := onfleet.Container{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{string(key), id},
		nil,
		nil,
		&container,
	)
	return container, err
}

// Reference https://docs.onfleet.com/reference/insert-tasks-at-index-or-append
//
// Reference https://docs.onfleet.com/reference/update-tasks
func (c *Client) InsertTasks(id string, key onfleet.ContainerQueryKey, params onfleet.ContainerTaskInsertParams) (onfleet.Container, error) {
	container := onfleet.Container{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{string(key), id},
		nil,
		params,
		&container,
	)
	return container, err
}
