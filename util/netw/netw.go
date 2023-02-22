package netw

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/onfleet/gonfleet/config"
	"github.com/onfleet/gonfleet/util/conv"
)

// NewHttpClient returns a pointer to a new http client with specified timeout.
// This http client can be shared across all resources.
func NewHttpClient(timeoutMillis int64) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeoutMillis) * time.Millisecond,
	}
}

// NewRequestParams contains Method, Url, and Body to be used in created a new http request via NewHttpRequest
type NewRequestParams struct {
	Method string // "GET", "POST", "PUT"
	Url    string
	Body   []byte
}

// NewHttpRequest reuturn an http request with package config parameters applied.
func NewHttpRequest(c config.Config, params NewRequestParams) (*http.Request, error) {
	var request *http.Request
	var err error
	switch params.Method {
	case "GET":
		request, err = http.NewRequest(
			params.Method,
			params.Url,
			nil,
		)
		request.Header.Set("Accept", "application/json")
	case "POST", "PUT":
		body := bytes.NewBuffer(params.Body)
		request, err = http.NewRequest(
			params.Method,
			params.Url,
			body,
		)
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", fmt.Sprintf("%s-%s", c.Name, c.Version))
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", conv.EncodeBase64(c.ApiKey)))
	return request, err
}
