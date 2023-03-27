package team

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

// Reference https://docs.onfleet.com/reference/get-single-team
func (c *Client) Get(teamId string) (onfleet.Team, error) {
	team := onfleet.Team{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{teamId},
		nil,
		nil,
		&team,
	)
	return team, err
}

// Reference https://docs.onfleet.com/reference/list-teams
func (c *Client) List() ([]onfleet.Team, error) {
	teams := []onfleet.Team{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		nil,
		nil,
		nil,
		&teams,
	)
	return teams, err
}

// Reference https://docs.onfleet.com/reference/create-team
func (c *Client) Create(params onfleet.TeamCreateParams) (onfleet.Team, error) {
	team := onfleet.Team{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&team,
	)
	return team, err
}

// Reference https://docs.onfleet.com/reference/update-team
func (c *Client) Update(teamId string, params onfleet.TeamUpdateParams) (onfleet.Team, error) {
	team := onfleet.Team{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{teamId},
		nil,
		params,
		&team,
	)
	return team, err
}

// Reference https://docs.onfleet.com/reference/delete-team
func (c *Client) Delete(teamId string) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodDelete,
		c.url,
		[]string{teamId},
		nil,
		nil,
		nil,
	)
	return err
}

// Reference https://docs.onfleet.com/reference/team-auto-dispatch
func (c *Client) AutoDispatch(teamId string, params *onfleet.TeamAutoDispatchParams) (onfleet.TeamAutoDispatch, error) {
	autoDispatch := onfleet.TeamAutoDispatch{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{teamId, "dispatch"},
		nil,
		params,
		&autoDispatch,
	)
	return autoDispatch, err
}

// Reference https://docs.onfleet.com/reference/delivery-estimate
func (c *Client) GetWorkerEta(teamId string, params onfleet.TeamWorkerEtaQueryParams) (onfleet.TeamWorkerEta, error) {
	teamWorkerEta := onfleet.TeamWorkerEta{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{teamId, "estimate"},
		params,
		nil,
		&teamWorkerEta,
	)
	return teamWorkerEta, err
}

// Reference https://docs.onfleet.com/reference/list-tasks-in-team
func (c *Client) ListTasks(teamId string, params onfleet.TeamTasksListQueryParams) (onfleet.TeamTasks, error) {
	teamTasks := onfleet.TeamTasks{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{teamId, "tasks"},
		nil,
		nil,
		&teamTasks,
	)
	return teamTasks, err
}
