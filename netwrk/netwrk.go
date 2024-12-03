package netwrk

import (
	"bytes"
	"context"
	"encoding/json"
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
	var request *http.Request
	var err error

	calledUrl := baseUrl
	if pathSegments != nil {
		calledUrl = urlAttachPath(calledUrl, pathSegments...)
	}
	if queryParams != nil {
		calledUrl = urlAttachQuery(calledUrl, queryParams)
	}

	switch method {
	case "GET", "DELETE":
		request, err = http.NewRequest(
			method,
			calledUrl,
			nil,
		)
		if err != nil {
			return err
		}
		request.Header.Set("Accept", "application/json")
	case "POST", "PUT":
		marshalledBody, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}
		buffer := bytes.NewBuffer(marshalledBody)
		request, err = http.NewRequest(
			method,
			calledUrl,
			buffer,
		)
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json")
	default:
		return fmt.Errorf("unsupported method")
	}

	for _, h := range additionalHeaders {
		if h != ([2]string{}) {
			request.Header.Set(h[0], h[1])
		}
	}
	request.Header.Set("User-Agent", fmt.Sprintf("%s-%s", version.Name, version.Value))
	request.SetBasicAuth(apiKey, "")

	ctx := context.Background()
	err = rlHttpClient.RateLimiter.Wait(ctx)
	if err != nil {
		return err
	}
	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 15 * time.Second
	b := backoff.WithContext(exponentialBackOff, ctx)
	var response *http.Response
	errRetryCancel := backoff.Retry(func() error {
		var errDo error
		response, errDo = rlHttpClient.Client.Do(request)
		if errDo != nil {
			return errDo
		}
		defer response.Body.Close()
		if response.StatusCode == http.StatusTooManyRequests {
			// Retry on rate limiting
			return fmt.Errorf("rate limited. Retry")
		}
		if response.StatusCode < 200 || response.StatusCode > 299 {
			return backoff.Permanent(onfleet.ParseError(response.Body))
		}
		return nil
	}, b)
	if errRetryCancel != nil {
		return errRetryCancel
	}
	if v == nil {
		return nil
	}
	if err = json.NewDecoder(response.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
