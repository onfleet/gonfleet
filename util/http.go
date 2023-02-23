package util

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/onfleet/gonfleet/constant"
)

// NewHttpClient returns a pointer to a new http client with specified timeout.
// This http client can be shared across all resources.
func NewHttpClient(timeoutMillis int64) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeoutMillis) * time.Millisecond,
	}
}

// Call creates and runs an http request.
// It returns the response.
func Call(httpClient *http.Client, apiKey string, method string, url string, body []byte) (*http.Response, error) {
	var request *http.Request
	var err error
	switch method {
	case "GET":
		request, err = http.NewRequest(
			method,
			url,
			nil,
		)
		request.Header.Set("Accept", "application/json")
	case "POST", "PUT":
		body := bytes.NewBuffer(body)
		request, err = http.NewRequest(
			method,
			url,
			body,
		)
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", fmt.Sprintf("%s-%s", constant.PkgName, constant.PkgVersion))
	request.SetBasicAuth(apiKey, "")
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func IsErrorStatus(status int) bool {
	return status < 200 && status > 299
}
