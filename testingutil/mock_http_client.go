package testingutil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/onfleet/gonfleet/netwrk"
)

// MockResponse represents a canned HTTP response for testing
type MockResponse struct {
	StatusCode int
	Body       interface{}
	Headers    map[string]string
}

// MockHTTPClient is a mock implementation of netwrk.RlHttpClient for testing
type MockHTTPClient struct {
	// Responses is a map of URL patterns to mock responses
	Responses map[string]MockResponse
	// RequestHistory stores all requests made during testing
	RequestHistory []*http.Request
	// T is the testing context for assertions
	T *testing.T
}

// NewMockHTTPClient creates a new mock HTTP client for testing
func NewMockHTTPClient(t *testing.T) *MockHTTPClient {
	return &MockHTTPClient{
		Responses:      make(map[string]MockResponse),
		RequestHistory: make([]*http.Request, 0),
		T:              t,
	}
}

// AddResponse adds a mock response for a given URL pattern
func (m *MockHTTPClient) AddResponse(urlPattern string, response MockResponse) {
	m.Responses[urlPattern] = response
}

// MockCaller is a test implementation of netwrk.Caller that uses the mock HTTP client
func (m *MockHTTPClient) MockCaller(
	apiKey string,
	rlHttpClient *netwrk.RlHttpClient,
	method string,
	baseUrl string,
	pathSegments []string,
	queryParams any,
	body any,
	v any,
	additionalHeaders ...[2]string,
) error {
	// Build the full URL
	fullURL := baseUrl
	if pathSegments != nil {
		fullURL = fullURL + "/" + strings.Join(pathSegments, "/")
	}

	// For testing, we'll use a simpler approach to avoid URL parsing issues
	// We'll create a basic request object for history without complex URL parsing
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		// If URL parsing fails, create a simple URL object
		parsedURL = &url.URL{Path: fullURL}
	}
	
	req := &http.Request{
		Method: method,
		URL:    parsedURL,
		Header: make(http.Header),
	}
	req.SetBasicAuth(apiKey, "")
	
	// Add additional headers
	for _, h := range additionalHeaders {
		if h != ([2]string{}) {
			req.Header.Set(h[0], h[1])
		}
	}

	// Store request in history
	m.RequestHistory = append(m.RequestHistory, req)

	// Find matching response
	var response MockResponse
	var found bool
	
	// Try exact URL match first
	if resp, ok := m.Responses[fullURL]; ok {
		response = resp
		found = true
	} else {
		// Try pattern matching
		for pattern, resp := range m.Responses {
			if strings.Contains(fullURL, pattern) {
				response = resp
				found = true
				break
			}
		}
	}

	if !found {
		return fmt.Errorf("no mock response found for URL: %s", fullURL)
	}

	// Simulate HTTP status code errors
	if response.StatusCode >= 400 {
		return fmt.Errorf("HTTP %d error", response.StatusCode)
	}

	// Marshal response body and unmarshal into target
	if v != nil && response.Body != nil {
		bodyBytes, err := json.Marshal(response.Body)
		if err != nil {
			return err
		}
		
		if err := json.Unmarshal(bodyBytes, v); err != nil {
			return err
		}
	}

	return nil
}

// GetLastRequest returns the most recent request made
func (m *MockHTTPClient) GetLastRequest() *http.Request {
	if len(m.RequestHistory) == 0 {
		return nil
	}
	return m.RequestHistory[len(m.RequestHistory)-1]
}

// GetRequestCount returns the total number of requests made
func (m *MockHTTPClient) GetRequestCount() int {
	return len(m.RequestHistory)
}

// AssertRequestMade checks if a request was made to the expected URL with expected method
func (m *MockHTTPClient) AssertRequestMade(expectedMethod, expectedURL string) {
	for _, req := range m.RequestHistory {
		if req.Method == expectedMethod && strings.Contains(req.URL.String(), expectedURL) {
			return // Found matching request
		}
	}
	m.T.Errorf("Expected request %s %s was not made", expectedMethod, expectedURL)
}

// AssertBasicAuth checks if the last request used correct basic auth
func (m *MockHTTPClient) AssertBasicAuth(expectedAPIKey string) {
	lastReq := m.GetLastRequest()
	if lastReq == nil {
		m.T.Error("No requests made")
		return
	}
	
	username, password, ok := lastReq.BasicAuth()
	if !ok {
		m.T.Error("Basic auth not found in request")
		return
	}
	
	if username != expectedAPIKey {
		m.T.Errorf("Expected API key %s, got %s", expectedAPIKey, username)
	}
	
	if password != "" {
		m.T.Errorf("Expected empty password, got %s", password)
	}
}

// Reset clears all request history and responses
func (m *MockHTTPClient) Reset() {
	m.RequestHistory = make([]*http.Request, 0)
	m.Responses = make(map[string]MockResponse)
}