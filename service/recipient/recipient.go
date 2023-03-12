package recipient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/util"
)

type caller func(apiKey string, httpClient *http.Client, method string, url string, body []byte) (*http.Response, error)
type errorParser func(r io.Reader) error

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
	call       caller
	parseError errorParser
}

func Register(apiKey string, httpClient *http.Client, url string, call caller, parseError errorParser) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
		call:       call,
		parseError: parseError,
	}
}

func (c *Client) Get(recipientId string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := fmt.Sprintf("%s/%s", c.url, recipientId)
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return recipient, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return recipient, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&recipient); err != nil {
		return recipient, err
	}
	return recipient, nil
}

// Find searches for recipient based on provided name of phone value.
// Key options are "name" and "phone".
//
// e.g.
//
// Find("jane doe", "name")
// Find("3105550100", "phone")
func (c *Client) Find(value string, key string) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	allowed := []string{"name", "phone"}
	if !util.Contains(allowed, key) {
		return recipient, fmt.Errorf("invalid query key %s", key)
	}
	url := fmt.Sprintf("%s/%s/%s", c.url, key, url.PathEscape(value))
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return recipient, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return recipient, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&recipient); err != nil {
		return recipient, err
	}
	return recipient, nil
}

func (c *Client) Update(recipientId string, params onfleet.RecipientUpdateParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := fmt.Sprintf("%s/%s", c.url, recipientId)
	body, err := json.Marshal(params)
	if err != nil {
		return recipient, err
	}
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodPut, url, body)
	if err != nil {
		return recipient, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return recipient, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&recipient); err != nil {
		return recipient, err
	}
	return recipient, nil
}

func (c *Client) Create(params onfleet.RecipientCreationParams) (onfleet.Recipient, error) {
	recipient := onfleet.Recipient{}
	url := c.url
	body, err := json.Marshal(params)
	if err != nil {
		return recipient, err
	}
	resp, err := c.call(c.apiKey, c.httpClient, http.MethodPost, url, body)
	if err != nil {
		return recipient, err
	}
	defer resp.Body.Close()
	if util.IsErrorStatus(resp.StatusCode) {
		return recipient, c.parseError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&recipient); err != nil {
		return recipient, err
	}
	return recipient, nil
}
