package netwrk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/time/rate"

	onfleet "github.com/onfleet/gonfleet"
)

func TestNewRlHttpClient(t *testing.T) {
	tests := []struct {
		name    string
		timeout int64
	}{
		{
			name:    "default timeout",
			timeout: 70000,
		},
		{
			name:    "custom timeout",
			timeout: 30000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := rate.NewLimiter(rate.Every(1*time.Second), 10)
			client := NewRlHttpClient(rl, tt.timeout)

			if client == nil {
				t.Error("Expected non-nil client")
			}
			if client.RateLimiter != rl {
				t.Error("Rate limiter not set correctly")
			}
			if client.Client == nil {
				t.Error("HTTP client not initialized")
			}

			expectedTimeout := time.Duration(tt.timeout) * time.Millisecond
			if client.Client.Timeout != expectedTimeout {
				t.Errorf("Expected timeout %v, got %v", expectedTimeout, client.Client.Timeout)
			}
		})
	}
}

func TestUrlAttachPath(t *testing.T) {
	tests := []struct {
		name         string
		baseUrl      string
		pathSegments []string
		expected     string
	}{
		{
			name:         "single path segment",
			baseUrl:      "https://api.example.com",
			pathSegments: []string{"users"},
			expected:     "https://api.example.com/users",
		},
		{
			name:         "multiple path segments",
			baseUrl:      "https://api.example.com",
			pathSegments: []string{"api", "v1", "users"},
			expected:     "https://api.example.com/api/v1/users",
		},
		{
			name:         "empty path segments",
			baseUrl:      "https://api.example.com",
			pathSegments: []string{},
			expected:     "https://api.example.com/",
		},
		{
			name:         "nil path segments",
			baseUrl:      "https://api.example.com",
			pathSegments: nil,
			expected:     "https://api.example.com/",
		},
		{
			name:         "base url with trailing slash",
			baseUrl:      "https://api.example.com/",
			pathSegments: []string{"users"},
			expected:     "https://api.example.com/users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := urlAttachPath(tt.baseUrl, tt.pathSegments...)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestStomp(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]any
		hasError bool
	}{
		{
			name: "simple struct",
			input: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{
				Name: "John",
				Age:  30,
			},
			expected: map[string]any{
				"name": "John",
				"age":  float64(30), // JSON unmarshals numbers as float64
			},
			hasError: false,
		},
		{
			name:     "nil input",
			input:    nil,
			expected: map[string]any{},
			hasError: false,
		},
		{
			name: "nested struct",
			input: struct {
				User struct {
					Name string `json:"name"`
				} `json:"user"`
				Count int `json:"count"`
			}{
				User: struct {
					Name string `json:"name"`
				}{Name: "Alice"},
				Count: 5,
			},
			expected: map[string]any{
				"user": map[string]any{
					"name": "Alice",
				},
				"count": float64(5),
			},
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := stomp(tt.input)

			if tt.hasError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !equalMaps(result, tt.expected) {
				t.Errorf("Expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

func TestUrlAttachQuery(t *testing.T) {
	tests := []struct {
		name     string
		baseUrl  string
		params   interface{}
		expected string
	}{
		{
			name:    "simple query params",
			baseUrl: "https://api.example.com/users",
			params: struct {
				Page  int    `json:"page"`
				Limit int    `json:"limit"`
				Name  string `json:"name"`
			}{
				Page:  1,
				Limit: 10,
				Name:  "John",
			},
			expected: "https://api.example.com/users?limit=10&name=John&page=1",
		},
		{
			name:     "nil params",
			baseUrl:  "https://api.example.com/users",
			params:   nil,
			expected: "https://api.example.com/users",
		},
		{
			name:    "empty struct",
			baseUrl: "https://api.example.com/users",
			params:  struct{}{},
			expected: "https://api.example.com/users",
		},
		{
			name:    "url with existing query",
			baseUrl: "https://api.example.com/users?existing=true",
			params: struct {
				Page int `json:"page"`
			}{Page: 2},
			expected: "https://api.example.com/users?existing=true&page=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := urlAttachQuery(tt.baseUrl, tt.params)
			
			// For URL query parameters, order might vary, so we need to check components
			if tt.params == nil || (fmt.Sprintf("%T", tt.params) == "struct {}" && 
				fmt.Sprintf("%+v", tt.params) == "{}") {
				if result != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected, result)
				}
			} else {
				// Check that all expected parameters are present
				if !strings.Contains(result, "api.example.com/users") {
					t.Errorf("Base URL not preserved in %s", result)
				}
				if tt.name == "simple query params" {
					if !strings.Contains(result, "page=1") || 
					   !strings.Contains(result, "limit=10") || 
					   !strings.Contains(result, "name=John") {
						t.Errorf("Missing expected query parameters in %s", result)
					}
				}
			}
		})
	}
}

func TestCallInternal_GET(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		// Verify basic auth
		username, password, ok := r.BasicAuth()
		if !ok {
			t.Error("Basic auth not found")
		}
		if username != "test_api_key" {
			t.Errorf("Expected username 'test_api_key', got '%s'", username)
		}
		if password != "" {
			t.Errorf("Expected empty password, got '%s'", password)
		}

		// Verify headers
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header 'application/json', got '%s'", r.Header.Get("Accept"))
		}

		// Verify User-Agent
		userAgent := r.Header.Get("User-Agent")
		if !strings.Contains(userAgent, "onfleet/gonfleet") {
			t.Errorf("Expected User-Agent to contain 'onfleet/gonfleet', got '%s'", userAgent)
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]any{
			"id":   "test_123",
			"name": "Test Item",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create rate limiter and HTTP client
	rl := rate.NewLimiter(rate.Every(1*time.Second), 10)
	rlHttpClient := NewRlHttpClient(rl, 5000)

	// Test successful GET request
	var result map[string]any
	err := callInternal(
		context.Background(),
		"test_api_key",
		rlHttpClient,
		"GET",
		server.URL+"/test",
		nil,
		nil,
		nil,
		&result,
		[][2]string{},
	)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedID := "test_123"
	if result["id"] != expectedID {
		t.Errorf("Expected id '%s', got '%v'", expectedID, result["id"])
	}

	expectedName := "Test Item"
	if result["name"] != expectedName {
		t.Errorf("Expected name '%s', got '%v'", expectedName, result["name"])
	}
}

func TestCallInternal_POST(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		// Verify Content-Type header
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		// Read and verify request body
		var requestBody map[string]any
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}

		if requestBody["name"] != "New Item" {
			t.Errorf("Expected name 'New Item', got '%v'", requestBody["name"])
		}

		// Return created response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := map[string]any{
			"id":   "created_123",
			"name": requestBody["name"],
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create rate limiter and HTTP client
	rl := rate.NewLimiter(rate.Every(1*time.Second), 10)
	rlHttpClient := NewRlHttpClient(rl, 5000)

	// Test successful POST request
	requestBody := map[string]any{
		"name": "New Item",
	}

	var result map[string]any
	err := callInternal(
		context.Background(),
		"test_api_key",
		rlHttpClient,
		"POST",
		server.URL+"/test",
		nil,
		nil,
		requestBody,
		&result,
		[][2]string{},
	)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result["id"] != "created_123" {
		t.Errorf("Expected id 'created_123', got '%v'", result["id"])
	}
}

func TestCallInternal_ErrorResponses(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   map[string]any
		expectedError  bool
		expectedTooMany bool
	}{
		{
			name:            "400 Bad Request",
			statusCode:      http.StatusBadRequest,
			responseBody:    map[string]any{"error": "Invalid request"},
			expectedError:   true,
			expectedTooMany: false,
		},
		{
			name:            "404 Not Found",
			statusCode:      http.StatusNotFound,
			responseBody:    map[string]any{"error": "Not found"},
			expectedError:   true,
			expectedTooMany: false,
		},
		{
			name:            "429 Too Many Requests",
			statusCode:      http.StatusTooManyRequests,
			responseBody:    map[string]any{"error": "Rate limited"},
			expectedError:   true,
			expectedTooMany: true,
		},
		{
			name:            "412 Precondition Failed",
			statusCode:      http.StatusPreconditionFailed,
			responseBody:    map[string]any{"error": "Precondition failed"},
			expectedError:   true,
			expectedTooMany: true,
		},
		{
			name:            "500 Internal Server Error",
			statusCode:      http.StatusInternalServerError,
			responseBody:    map[string]any{"error": "Server error"},
			expectedError:   true,
			expectedTooMany: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				json.NewEncoder(w).Encode(tt.responseBody)
			}))
			defer server.Close()

			// Create rate limiter and HTTP client
			rl := rate.NewLimiter(rate.Every(1*time.Second), 10)
			rlHttpClient := NewRlHttpClient(rl, 5000)

			var result map[string]any
			err := callInternal(
				context.Background(),
				"test_api_key",
				rlHttpClient,
				"GET",
				server.URL+"/test",
				nil,
				nil,
				nil,
				&result,
				[][2]string{},
			)

			if tt.expectedError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if tt.expectedTooMany {
				var tooManyErr onfleet.TooManyRequestsError
				if !strings.Contains(fmt.Sprintf("%T", err), "TooManyRequestsError") {
					t.Errorf("Expected TooManyRequestsError, got %T", err)
				}
				_ = tooManyErr // Use the variable to avoid unused error
			}
		})
	}
}

func TestCallInternal_AdditionalHeaders(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify custom headers
		if r.Header.Get("X-Custom-Header") != "custom-value" {
			t.Errorf("Expected X-Custom-Header 'custom-value', got '%s'", r.Header.Get("X-Custom-Header"))
		}
		if r.Header.Get("X-Another-Header") != "another-value" {
			t.Errorf("Expected X-Another-Header 'another-value', got '%s'", r.Header.Get("X-Another-Header"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	// Create rate limiter and HTTP client
	rl := rate.NewLimiter(rate.Every(1*time.Second), 10)
	rlHttpClient := NewRlHttpClient(rl, 5000)

	err := callInternal(
		context.Background(),
		"test_api_key",
		rlHttpClient,
		"GET",
		server.URL+"/test",
		nil,
		nil,
		nil,
		nil,
		[][2]string{
			{"X-Custom-Header", "custom-value"},
			{"X-Another-Header", "another-value"},
		},
	)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCallInternal_UnsupportedMethod(t *testing.T) {
	rl := rate.NewLimiter(rate.Every(1*time.Second), 10)
	rlHttpClient := NewRlHttpClient(rl, 5000)

	err := callInternal(
		context.Background(),
		"test_api_key",
		rlHttpClient,
		"PATCH", // Unsupported method
		"https://example.com/test",
		nil,
		nil,
		nil,
		nil,
		[][2]string{},
	)

	if err == nil {
		t.Error("Expected error for unsupported method")
	}

	if !strings.Contains(err.Error(), "unsupported method") {
		t.Errorf("Expected 'unsupported method' error, got: %v", err)
	}
}

func TestCall_WithRetry(t *testing.T) {
	requestCount := 0
	
	// Create test server that fails first time, succeeds second time
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		
		if requestCount == 1 {
			// First request: return 429 (should retry)
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		
		// Second request: succeed
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]any{"success": true}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create rate limiter and HTTP client
	rl := rate.NewLimiter(rate.Every(1*time.Millisecond), 100) // Very permissive for testing
	rlHttpClient := NewRlHttpClient(rl, 5000)

	var result map[string]any
	err := Call(
		"test_api_key",
		rlHttpClient,
		"GET",
		server.URL+"/test",
		nil,
		nil,
		nil,
		&result,
		[2]string{},
	)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests (1 retry), got %d", requestCount)
	}

	if result["success"] != true {
		t.Errorf("Expected success=true, got %v", result["success"])
	}
}

func TestCall_PermanentError(t *testing.T) {
	requestCount := 0
	
	// Create test server that always returns 400 (permanent error)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	// Create rate limiter and HTTP client
	rl := rate.NewLimiter(rate.Every(1*time.Millisecond), 100)
	rlHttpClient := NewRlHttpClient(rl, 5000)

	err := Call(
		"test_api_key",
		rlHttpClient,
		"GET",
		server.URL+"/test",
		nil,
		nil,
		nil,
		nil,
		[2]string{},
	)

	if err == nil {
		t.Error("Expected error for 400 status")
	}

	if requestCount != 1 {
		t.Errorf("Expected 1 request (no retry for permanent error), got %d", requestCount)
	}
}

// Helper function to compare maps
func equalMaps(a, b map[string]any) bool {
	if len(a) != len(b) {
		return false
	}
	
	for k, v := range a {
		if bv, exists := b[k]; !exists {
			return false
		} else if !equalValues(v, bv) {
			return false
		}
	}
	
	return true
}

// Helper function to compare interface{} values
func equalValues(a, b any) bool {
	// Handle nested maps
	if aMap, aOk := a.(map[string]any); aOk {
		if bMap, bOk := b.(map[string]any); bOk {
			return equalMaps(aMap, bMap)
		}
		return false
	}
	
	return a == b
}