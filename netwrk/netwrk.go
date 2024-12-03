package netwrk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"

	onfleet "github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/version"
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

// urlAttachPath appends path segments onto provided baseUrl.
func urlAttachPath(baseUrl string, pathSegments ...string) string {
	newUrl, err := url.JoinPath(baseUrl, pathSegments...)
	if err != nil {
		return baseUrl
	}
	return newUrl
}

// stomp converts a struct to a map[string]any
func stomp(v any) (map[string]any, error) {
	m := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(b, &m)
	return m, err
}

// urlAttachQuery sets query parameters on the provided baseUrl.
func urlAttachQuery(baseUrl string, v any) string {
	URL, err := url.Parse(baseUrl)
	if err != nil {
		return baseUrl
	}
	q := URL.Query()
	params, err := stomp(v)
	if err != nil {
		return baseUrl
	}
	for k, v := range params {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	URL.RawQuery = q.Encode()
	return URL.String()
}

type Caller func(
	apiKey string,
	rlHttpClient *RlHttpClient,
	method string,
	baseUrl string,
	pathSegments []string,
	queryParams any,
	body any,
	v any,
	additionalHeaders ...[2]string,
) error

func Call(
	apiKey string,
	rlHttpClient *RlHttpClient,
	method string,
	baseUrl string,
	pathSegments []string,
	queryParams any,
	body any,
	v any,
	additionalHeaders ...[2]string,
) error {
	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 15 * time.Second
	ctx := context.Background()
	b := backoff.WithContext(exponentialBackOff, ctx)
	return backoff.Retry(func() error {
		err := callInternal(ctx, apiKey, rlHttpClient, method, baseUrl, pathSegments, queryParams, body, v, additionalHeaders)
		if errors.Is(err, onfleet.TooManyRequestsError{}) {
			return err
		}
		if err != nil {
			return backoff.Permanent(err)
		}
		return nil
	}, b)
}

func callInternal(ctx context.Context, apiKey string, rlHttpClient *RlHttpClient, method string, baseUrl string, pathSegments []string, queryParams any, body any, v any, additionalHeaders [][2]string) error {
	var request *http.Request
	var err error

	callUrl := baseUrl
	if pathSegments != nil {
		callUrl = urlAttachPath(callUrl, pathSegments...)
	}
	if queryParams != nil {
		callUrl = urlAttachQuery(callUrl, queryParams)
	}

	switch method {
	case "GET", "DELETE":
		request, err = http.NewRequest(
			method,
			callUrl,
			nil,
		)
		if err != nil {
			return err
		}
		request.Header.Set("Accept", "application/json")
	case "POST", "PUT":
		bodyMarshal, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}
		buffer := bytes.NewBuffer(bodyMarshal)
		request, err = http.NewRequest(
			method,
			callUrl,
			buffer,
		)
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json")
	default:
		return fmt.Errorf("unsupported method: %s", method)
	}

	for _, h := range additionalHeaders {
		if h != ([2]string{}) {
			request.Header.Set(h[0], h[1])
		}
	}
	request.Header.Set("User-Agent", fmt.Sprintf("%s-%s", version.Name, version.Value))
	request.SetBasicAuth(apiKey, "")

	err = rlHttpClient.RateLimiter.Wait(ctx)
	if err != nil {
		return err
	}
	response, err := rlHttpClient.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusTooManyRequests || response.StatusCode == http.StatusPreconditionFailed {
		return onfleet.TooManyRequestsError{}
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return onfleet.ParseError(response.Body)
	}
	if v == nil {
		return nil
	}
	if err = json.NewDecoder(response.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
