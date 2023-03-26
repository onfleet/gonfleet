package container

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

// Get retrieves a container by id and query key.
//
// e.g.
//
// Get("2Fwp6wS5wLNjDn36r1LJPscA", "workers")
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

// InsertTasks adds tasks to the container specified by "id" and "key".
// Tasks are inserted based on position (provided as first element of the Tasks field on ContainerTaskInsertParams),
// where 0 prepends, -1 appends, and any number greater than 0 indicates an insertion index.
// Optionally, ConsiderDependencies may be set on ContainerTaskInsertParams (defaults to false). If
// ConsiderDependencies is set to true all child and/or parent tasks will be moved if one or more child / parent
// task ids are included in Tasks on ContainerTaskInsertParams.
//
// *Full task replacement only available on "worker" container*
// If "tasks" does not include an int as it's first element, all tasks on the specified container
// will be replaced with the tasks provided in "tasks".
//
// e.g. append tasks to existing tasks on container.
// InsertTasks(
//
//	"2Fwp6wS5wLNjDn36r1LJPscA",
//	"workers",
//	onfleet.ContainerTaskInsertParams{
//	    Tasks: []any{-1, "b3F~z2sU7H*auNKkM6LoiXzP", "1ry863mrjoQaqMNxnrD5YvxH"},
//	    ConsiderDependencies: true,
//	}
//
// )
//
// e.g. replace all tasks on existing container with provided "tasks".
// InsertTasks(
//
//	"2Fwp6wS5wLNjDn36r1LJPscA",
//	"workers",
//	onfleet.ContainerTaskInsertParams{
//	    Tasks: []any{"b3F~z2sU7H*auNKkM6LoiXzP", "1ry863mrjoQaqMNxnrD5YvxH"},
//	    ConsiderDependencies: true,
//	}
//
// )
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
