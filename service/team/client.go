package team

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

type caller func(apiKey string, httpClient *http.Client, method string, url string, body any, result any) error

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
	call       caller
}

func New(apiKey string, httpClient *http.Client, url string, call caller) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
	}
}

func (c *Client) List() ([]onfleet.Team, error) {
	teams := []onfleet.Team{}
	err := c.call(c.apiKey, c.httpClient, http.MethodGet, c.url, nil, &teams)
	return teams, err
}

func (c *Client) Create(params onfleet.TeamCreateParams) (onfleet.Team, error) {
	team := onfleet.Team{}
	err := c.call(c.apiKey, c.httpClient, http.MethodPost, c.url, params, &team)
	return team, err
}

func (c *Client) Update(teamId string, params onfleet.TeamUpdateParams) (onfleet.Team, error) {
	team := onfleet.Team{}
	url := util.UrlAttachPath(c.url, teamId)
	err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, params, &team)
	return team, err
}

func (c *Client) Delete(teamId string) error {
	url := util.UrlAttachPath(c.url, teamId)
	err := c.call(c.apiKey, c.httpClient, http.MethodDelete, url, nil, nil)
	return err
}
