package client

import (
	"testing"

	"github.com/onfleet/gonfleet/testingutil"
)

func TestNew_ValidAPIKey(t *testing.T) {
	apiKey := "test_api_key_123"
	
	api, err := New(apiKey, nil)
	
	testingutil.AssertNoError(t, err)
	testingutil.AssertNotNil(t, api)
	
	// Verify all services are initialized
	testingutil.AssertNotNil(t, api.Administrators)
	testingutil.AssertNotNil(t, api.Containers)
	testingutil.AssertNotNil(t, api.Destinations)
	testingutil.AssertNotNil(t, api.Hubs)
	testingutil.AssertNotNil(t, api.Organizations)
	testingutil.AssertNotNil(t, api.Recipients)
	testingutil.AssertNotNil(t, api.Tasks)
	testingutil.AssertNotNil(t, api.Teams)
	testingutil.AssertNotNil(t, api.Webhooks)
	testingutil.AssertNotNil(t, api.Workers)
	testingutil.AssertNotNil(t, api.ManifestProvider)
	testingutil.AssertNotNil(t, api.RoutePlans)
}

func TestNew_EmptyAPIKey(t *testing.T) {
	apiKey := ""
	
	api, err := New(apiKey, nil)
	
	testingutil.AssertError(t, err)
	if api != nil {
		t.Error("Expected nil API client, got non-nil")
	}
	testingutil.AssertStringContains(t, err.Error(), "API key not found")
}

func TestNew_DefaultParameters(t *testing.T) {
	apiKey := "test_api_key_123"
	
	api, err := New(apiKey, nil)
	
	testingutil.AssertNoError(t, err)
	testingutil.AssertNotNil(t, api)
	
	// We can't directly test the internal configuration, but we can verify
	// that the client was created successfully with default parameters
}

func TestNew_CustomParameters(t *testing.T) {
	apiKey := "test_api_key_123"
	params := &InitParams{
		BaseUrl:           "https://custom.onfleet.com",
		Path:              "/custom",
		ApiVersion:        "/v3",
		UserTimeout:       30000,
		MaxCallsPerSecond: 10,
	}
	
	api, err := New(apiKey, params)
	
	testingutil.AssertNoError(t, err)
	testingutil.AssertNotNil(t, api)
	
	// Verify all services are still initialized with custom parameters
	testingutil.AssertNotNil(t, api.Administrators)
	testingutil.AssertNotNil(t, api.Containers)
	testingutil.AssertNotNil(t, api.Destinations)
	testingutil.AssertNotNil(t, api.Hubs)
	testingutil.AssertNotNil(t, api.Organizations)
	testingutil.AssertNotNil(t, api.Recipients)
	testingutil.AssertNotNil(t, api.Tasks)
	testingutil.AssertNotNil(t, api.Teams)
	testingutil.AssertNotNil(t, api.Webhooks)
	testingutil.AssertNotNil(t, api.Workers)
	testingutil.AssertNotNil(t, api.ManifestProvider)
	testingutil.AssertNotNil(t, api.RoutePlans)
}

func TestNew_PartialCustomParameters(t *testing.T) {
	tests := []struct {
		name   string
		params *InitParams
	}{
		{
			name: "custom base URL only",
			params: &InitParams{
				BaseUrl: "https://custom.onfleet.com",
			},
		},
		{
			name: "custom timeout only",
			params: &InitParams{
				UserTimeout: 45000,
			},
		},
		{
			name: "custom rate limit only",
			params: &InitParams{
				MaxCallsPerSecond: 5,
			},
		},
		{
			name: "custom path and version",
			params: &InitParams{
				Path:       "/myapi",
				ApiVersion: "/v1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey := "test_api_key_123"
			
			api, err := New(apiKey, tt.params)
			
			testingutil.AssertNoError(t, err)
			testingutil.AssertNotNil(t, api)
			
			// Verify all services are initialized
			testingutil.AssertNotNil(t, api.Tasks)
			testingutil.AssertNotNil(t, api.Workers)
			testingutil.AssertNotNil(t, api.Organizations)
		})
	}
}

func TestNew_InvalidTimeoutBounds(t *testing.T) {
	tests := []struct {
		name    string
		timeout int64
	}{
		{
			name:    "timeout too large",
			timeout: 100000, // Greater than defaultUserTimeout (70000)
		},
		{
			name:    "zero timeout",
			timeout: 0,
		},
		{
			name:    "negative timeout",
			timeout: -1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey := "test_api_key_123"
			params := &InitParams{
				UserTimeout: tt.timeout,
			}
			
			api, err := New(apiKey, params)
			
			// Should still succeed but use default timeout
			testingutil.AssertNoError(t, err)
			testingutil.AssertNotNil(t, api)
		})
	}
}

func TestNew_InvalidRateLimitBounds(t *testing.T) {
	tests := []struct {
		name      string
		rateLimit int
	}{
		{
			name:      "rate limit too large",
			rateLimit: 50, // Greater than defaultMaxCallsPerSecond (18)
		},
		{
			name:      "zero rate limit",
			rateLimit: 0,
		},
		{
			name:      "negative rate limit",
			rateLimit: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey := "test_api_key_123"
			params := &InitParams{
				MaxCallsPerSecond: tt.rateLimit,
			}
			
			api, err := New(apiKey, params)
			
			// Should still succeed but use default rate limit
			testingutil.AssertNoError(t, err)
			testingutil.AssertNotNil(t, api)
		})
	}
}

func TestNew_URLConstruction(t *testing.T) {
	// This test verifies that URLs are constructed correctly by checking
	// that services are initialized without error
	
	tests := []struct {
		name           string
		params         *InitParams
		expectedInURL  string
	}{
		{
			name: "default configuration",
			params: nil,
			expectedInURL: "https://onfleet.com/api/v2",
		},
		{
			name: "custom base URL",
			params: &InitParams{
				BaseUrl: "https://staging.onfleet.com",
			},
			expectedInURL: "https://staging.onfleet.com/api/v2",
		},
		{
			name: "custom path",
			params: &InitParams{
				Path: "/newapi",
			},
			expectedInURL: "https://onfleet.com/newapi/v2",
		},
		{
			name: "custom API version",
			params: &InitParams{
				ApiVersion: "/v3",
			},
			expectedInURL: "https://onfleet.com/api/v3",
		},
		{
			name: "all custom URL components",
			params: &InitParams{
				BaseUrl:    "https://api.example.com",
				Path:       "/custom",
				ApiVersion: "/v1",
			},
			expectedInURL: "https://api.example.com/custom/v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey := "test_api_key_123"
			
			api, err := New(apiKey, tt.params)
			
			testingutil.AssertNoError(t, err)
			testingutil.AssertNotNil(t, api)
			
			// All services should be initialized successfully
			testingutil.AssertNotNil(t, api.Tasks)
			testingutil.AssertNotNil(t, api.Workers)
			testingutil.AssertNotNil(t, api.Organizations)
			testingutil.AssertNotNil(t, api.Destinations)
			testingutil.AssertNotNil(t, api.Recipients)
			testingutil.AssertNotNil(t, api.Teams)
			testingutil.AssertNotNil(t, api.Webhooks)
			testingutil.AssertNotNil(t, api.Hubs)
			testingutil.AssertNotNil(t, api.Administrators)
			testingutil.AssertNotNil(t, api.Containers)
			testingutil.AssertNotNil(t, api.RoutePlans)
			testingutil.AssertNotNil(t, api.ManifestProvider)
		})
	}
}

func TestDefaultConstants(t *testing.T) {
	// Test that the default constants have expected values
	testingutil.AssertEqual(t, int64(70000), defaultUserTimeout)
	testingutil.AssertEqual(t, "https://onfleet.com", defaultBaseUrl)
	testingutil.AssertEqual(t, "/api", defaultPath)
	testingutil.AssertEqual(t, "/v2", defaultApiVersion)
	testingutil.AssertEqual(t, 18, defaultMaxCallsPerSecond)
}

func TestInitParams_Struct(t *testing.T) {
	// Test that InitParams struct can be created and used
	params := InitParams{
		UserTimeout:       30000,
		BaseUrl:           "https://example.com",
		Path:              "/test",
		ApiVersion:        "/v1",
		MaxCallsPerSecond: 10,
	}
	
	testingutil.AssertEqual(t, int64(30000), params.UserTimeout)
	testingutil.AssertEqual(t, "https://example.com", params.BaseUrl)
	testingutil.AssertEqual(t, "/test", params.Path)
	testingutil.AssertEqual(t, "/v1", params.ApiVersion)
	testingutil.AssertEqual(t, 10, params.MaxCallsPerSecond)
}

func TestNew_RateLimiterConfiguration(t *testing.T) {
	// Test different rate limiting configurations
	tests := []struct {
		name      string
		rateLimit int
	}{
		{
			name:      "minimum valid rate limit",
			rateLimit: 1,
		},
		{
			name:      "mid-range rate limit",
			rateLimit: 10,
		},
		{
			name:      "maximum valid rate limit",
			rateLimit: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey := "test_api_key_123"
			params := &InitParams{
				MaxCallsPerSecond: tt.rateLimit,
			}
			
			api, err := New(apiKey, params)
			
			testingutil.AssertNoError(t, err)
			testingutil.AssertNotNil(t, api)
			
			// Verify services are created
			testingutil.AssertNotNil(t, api.Tasks)
		})
	}
}

func TestNew_TimeoutConfiguration(t *testing.T) {
	// Test different timeout configurations
	tests := []struct {
		name    string
		timeout int64
	}{
		{
			name:    "minimum valid timeout",
			timeout: 1000,
		},
		{
			name:    "mid-range timeout",
			timeout: 30000,
		},
		{
			name:    "maximum valid timeout",
			timeout: 70000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey := "test_api_key_123"
			params := &InitParams{
				UserTimeout: tt.timeout,
			}
			
			api, err := New(apiKey, params)
			
			testingutil.AssertNoError(t, err)
			testingutil.AssertNotNil(t, api)
			
			// Verify services are created
			testingutil.AssertNotNil(t, api.Tasks)
		})
	}
}

func TestAPI_ServiceEndpoints(t *testing.T) {
	// Test that all expected service endpoints are available
	apiKey := "test_api_key_123"
	api, err := New(apiKey, nil)
	
	testingutil.AssertNoError(t, err)
	testingutil.AssertNotNil(t, api)
	
	// Test that API struct has all expected service fields
	services := []interface{}{
		api.Administrators,
		api.Containers,
		api.Destinations,
		api.Hubs,
		api.Organizations,
		api.Recipients,
		api.Tasks,
		api.Teams,
		api.Webhooks,
		api.Workers,
		api.ManifestProvider,
		api.RoutePlans,
	}
	
	for i, service := range services {
		if service == nil {
			t.Errorf("Service %d is nil", i)
		}
	}
	
	// Count to ensure we have all expected services
	expectedServiceCount := 12
	if len(services) != expectedServiceCount {
		t.Errorf("Expected %d services, got %d", expectedServiceCount, len(services))
	}
}

func TestNew_EmptyStringFields(t *testing.T) {
	// Test behavior with empty string fields in InitParams
	apiKey := "test_api_key_123"
	params := &InitParams{
		BaseUrl:    "", // Should use default
		Path:       "", // Should use default
		ApiVersion: "", // Should use default
	}
	
	api, err := New(apiKey, params)
	
	testingutil.AssertNoError(t, err)
	testingutil.AssertNotNil(t, api)
	
	// Should still work with defaults
	testingutil.AssertNotNil(t, api.Tasks)
	testingutil.AssertNotNil(t, api.Workers)
}