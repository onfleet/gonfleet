package task

import (
	"net/http"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        string
}

func Register(apiKey string, httpClient *http.Client, url string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
		url:        url,
	}
}

func (c *Client) GetById(taskId string) {
}

func (c *Client) GetByShortId(taskShortId string) {
}

func (c *Client) Get() {
}
