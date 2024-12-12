package routePlan

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

// Reference https://docs.onfleet.com/reference/post-create-route-plan
func (c *Client) Create(params onfleet.RoutePlanParams) (onfleet.RoutePlan, error) {
	routePlan := onfleet.RoutePlan{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&routePlan,
	)
	return routePlan, err
}

// Reference https://docs.onfleet.com/reference/update-route-plan
func (c *Client) Update(routePlanId string, params onfleet.RoutePlanParams) (onfleet.RoutePlan, error) {
	routePlan := onfleet.RoutePlan{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{routePlanId},
		nil,
		params,
		&routePlan,
	)
	return routePlan, err
}

// Reference https://docs.onfleet.com/reference/add-tasks-to-route-plan
func (c *Client) AddTasks(routePlanId string, params onfleet.RoutePlanAddTasksParams) (onfleet.RoutePlan, error) {
	routePlan := onfleet.RoutePlan{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{routePlanId},
		nil,
		params,
		&routePlan,
	)
	return routePlan, err
}

// Reference https://docs.onfleet.com/reference/get-routeplan-by-id
func (c *Client) Get(routePlanId string) (onfleet.RoutePlan, error) {
	routePlan := onfleet.RoutePlan{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{routePlanId},
		nil,
		nil,
		&routePlan,
	)
	return routePlan, err
}

// Reference https://docs.onfleet.com/reference/get-route-plan
func (c *Client) List(params onfleet.RoutePlanListQueryParams) (onfleet.RoutePlansPaginated, error) {
	paginatedRoutePlans := onfleet.RoutePlansPaginated{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{"all"},
		params,
		nil,
		&paginatedRoutePlans,
	)
	return paginatedRoutePlans, err
}

// Reference https://docs.onfleet.com/reference/delete-routePlan
func (c *Client) Delete(routePlanId string) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodDelete,
		c.url,
		[]string{routePlanId},
		nil,
		nil,
		nil,
	)
	return err
}
