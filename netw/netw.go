package netw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/onfleet/gonfleet/pkg"
)

type RlHttpClient struct {
	Client      *http.Client
	RateLimiter *rate.Limiter
}

func NewRlHttpClient(rl *rate.Limiter, timeout int64) *RlHttpClient {
	return &RlHttpClient{
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Millisecond,
		},
		RateLimiter: rl,
	}
}

type requestErrorMessage struct {
	Cause any `json:"cause,omitempty"`
	// Error is an internal error code.
	// It is different than the request status code.
	Error int `json:"error"`
	// Message is the error messages / description
	Message string `json:"message"`
	// RemoteAddress is remote ip
	RemoteAddress string `json:"remoteAddress"`
	// Request is uuid associated with the request
	Request string `json:"request"`
}

type requestError struct {
	// Code is error type e.g. "InvalidArgument"
	Code string `json:"code"`
	// Message contains futher details about the error.
	Message requestErrorMessage `json:"message"`
}

func (err requestError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message.Message)
}

func parseError(r io.Reader) error {
	var reqError requestError
	if err := json.NewDecoder(r).Decode(&reqError); err != nil {
		return err
	}
	return reqError
}

type Caller func(apiKey string, rlHttpClient *RlHttpClient, method string, url string, body any, v any) error

func Call(apiKey string, rlHttpClient *RlHttpClient, method string, url string, body any, v any) error {
	var request *http.Request
	var err error
	switch method {
	case "GET", "DELETE":
		request, err = http.NewRequest(
			method,
			url,
			nil,
		)
		request.Header.Set("Accept", "application/json")
	case "POST", "PUT":
		body, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buffer := bytes.NewBuffer(body)
		request, err = http.NewRequest(
			method,
			url,
			buffer,
		)
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", fmt.Sprintf("%s-%s", pkg.Name, pkg.Version))
	request.SetBasicAuth(apiKey, "")

	ctx := context.Background()
	err = rlHttpClient.RateLimiter.Wait(ctx)
	if err != nil {
		return err
	}
	response, err := rlHttpClient.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return parseError(response.Body)
	}
	if v == nil {
		return nil
	}
	if err := json.NewDecoder(response.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
