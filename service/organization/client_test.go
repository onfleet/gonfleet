package organization

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_Get(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedOrg := testingutil.GetSampleOrganization()
	mockClient.AddResponse("/organization", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedOrg,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

	org, err := client.Get()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrg.ID, org.ID)
	assert.Equal(t, expectedOrg.Name, org.Name)
	assert.Equal(t, expectedOrg.Email, org.Email)

	mockClient.AssertRequestMade("GET", "/organization")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_Get_Error(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/organization", testingutil.MockResponse{
		StatusCode: 401,
		Body:       testingutil.GetSampleErrorResponse(),
	})

	client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

	org, err := client.Get()

	assert.Error(t, err)
	assert.Equal(t, "", org.ID)
}

func TestClient_GetDelegate(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedDelegate := onfleet.OrganizationDelegate{
		ID:                 "delegate_123",
		Name:               "Delegate Organization",
		Email:              "delegate@example.com",
		DriverSupportEmail: "support@delegate.com",
		Country:            "US",
		Timezone:           "America/Los_Angeles",
		IsFulfillment:      true,
	}

	mockClient.AddResponse("/organizations/delegate_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedDelegate,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

	delegate, err := client.GetDelegate("delegate_123")

	assert.NoError(t, err)
	assert.Equal(t, expectedDelegate.ID, delegate.ID)
	assert.Equal(t, expectedDelegate.Name, delegate.Name)
	assert.Equal(t, expectedDelegate.Email, delegate.Email)
	assert.True(t, delegate.IsFulfillment)

	mockClient.AssertRequestMade("GET", "/organizations/delegate_123")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_GetDelegate_NotFound(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/organizations/nonexistent", testingutil.MockResponse{
		StatusCode: 404,
		Body:       testingutil.GetSampleErrorResponse(),
	})

	client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

	delegate, err := client.GetDelegate("nonexistent")

	assert.Error(t, err)
	assert.Equal(t, "", delegate.ID)
}

func TestClient_GetDelegate_Forbidden(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/organizations/delegate_123", testingutil.MockResponse{
		StatusCode: 403,
		Body:       testingutil.GetSampleErrorResponse(),
	})

	client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

	delegate, err := client.GetDelegate("delegate_123")

	assert.Error(t, err)
	assert.Equal(t, "", delegate.ID)
}

// Test that the client uses the correct URLs for different operations
func TestClient_URLUsage(t *testing.T) {
	tests := []struct {
		name        string
		operation   string
		expectedURL string
	}{
		{
			name:        "get organization uses organization URL",
			operation:   "get",
			expectedURL: "/organization",
		},
		{
			name:        "get delegate uses organizations URL",
			operation:   "delegate",
			expectedURL: "/organizations/delegate_123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			if tt.operation == "get" {
				expectedOrg := testingutil.GetSampleOrganization()
				mockClient.AddResponse(tt.expectedURL, testingutil.MockResponse{
					StatusCode: 200,
					Body:       expectedOrg,
				})

				client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)
				_, err := client.Get()
				assert.NoError(t, err)

			} else if tt.operation == "delegate" {
				expectedDelegate := onfleet.OrganizationDelegate{
					ID:   "delegate_123",
					Name: "Test Delegate",
				}
				mockClient.AddResponse(tt.expectedURL, testingutil.MockResponse{
					StatusCode: 200,
					Body:       expectedDelegate,
				})

				client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)
				_, err := client.GetDelegate("delegate_123")
				assert.NoError(t, err)
			}

			mockClient.AssertRequestMade("GET", tt.expectedURL)
		})
	}
}

// Test different organization configurations
func TestClient_OrganizationConfigurations(t *testing.T) {
	tests := []struct {
		name    string
		orgData onfleet.Organization
	}{
		{
			name: "US organization",
			orgData: onfleet.Organization{
				ID:       "org_us",
				Name:     "US Company",
				Country:  "US",
				Timezone: "America/New_York",
				Email:    "us@company.com",
			},
		},
		{
			name: "CA organization",
			orgData: onfleet.Organization{
				ID:       "org_ca",
				Name:     "Canadian Company",
				Country:  "CA",
				Timezone: "America/Toronto",
				Email:    "ca@company.com",
			},
		},
		{
			name: "organization with delegatees",
			orgData: onfleet.Organization{
				ID:          "org_with_delegates",
				Name:        "Parent Company",
				Country:     "US",
				Timezone:    "America/Los_Angeles",
				Email:       "parent@company.com",
				Delegatees:  []string{"delegate_1", "delegate_2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			mockClient.AddResponse("/organization", testingutil.MockResponse{
				StatusCode: 200,
				Body:       tt.orgData,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

			org, err := client.Get()

			assert.NoError(t, err)
			assert.Equal(t, tt.orgData.ID, org.ID)
			assert.Equal(t, tt.orgData.Name, org.Name)
			assert.Equal(t, tt.orgData.Country, org.Country)
			assert.Equal(t, tt.orgData.Timezone, org.Timezone)
			assert.Equal(t, len(tt.orgData.Delegatees), len(org.Delegatees))
		})
	}
}

// Test different delegate configurations
func TestClient_DelegateConfigurations(t *testing.T) {
	tests := []struct {
		name         string
		delegateData onfleet.OrganizationDelegate
	}{
		{
			name: "fulfillment delegate",
			delegateData: onfleet.OrganizationDelegate{
				ID:            "fulfillment_delegate",
				Name:          "Fulfillment Partner",
				IsFulfillment: true,
				Country:       "US",
				Timezone:      "America/Chicago",
			},
		},
		{
			name: "non-fulfillment delegate",
			delegateData: onfleet.OrganizationDelegate{
				ID:            "partner_delegate",
				Name:          "Business Partner",
				IsFulfillment: false,
				Country:       "CA",
				Timezone:      "America/Vancouver",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			mockClient.AddResponse("/organizations/"+tt.delegateData.ID, testingutil.MockResponse{
				StatusCode: 200,
				Body:       tt.delegateData,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

			delegate, err := client.GetDelegate(tt.delegateData.ID)

			assert.NoError(t, err)
			assert.Equal(t, tt.delegateData.ID, delegate.ID)
			assert.Equal(t, tt.delegateData.Name, delegate.Name)
			assert.Equal(t, tt.delegateData.IsFulfillment, delegate.IsFulfillment)
			assert.Equal(t, tt.delegateData.Country, delegate.Country)
		})
	}
}

// Test error scenarios
func TestClient_ErrorScenarios(t *testing.T) {
	tests := []struct {
		name       string
		operation  string
		statusCode int
		expectErr  bool
	}{
		{"get org unauthorized", "get", 401, true},
		{"get org forbidden", "get", 403, true},
		{"get org server error", "get", 500, true},
		{"get delegate unauthorized", "delegate", 401, true},
		{"get delegate forbidden", "delegate", 403, true},
		{"get delegate not found", "delegate", 404, true},
		{"get delegate server error", "delegate", 500, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			var url string
			if tt.operation == "get" {
				url = "/organization"
			} else {
				url = "/organizations/test_delegate"
			}

			mockClient.AddResponse(url, testingutil.MockResponse{
				StatusCode: tt.statusCode,
				Body:       testingutil.GetSampleErrorResponse(),
			})

			client := Plug("test_api_key", nil, "https://api.example.com/organization", "https://api.example.com/organizations", mockClient.MockCaller)

			var err error
			if tt.operation == "get" {
				_, err = client.Get()
			} else {
				_, err = client.GetDelegate("test_delegate")
			}

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Test client initialization with different URL configurations
func TestClient_DifferentURLConfigurations(t *testing.T) {
	tests := []struct {
		name              string
		organizationURL   string
		organizationsURL  string
		expectedOrgURL    string
		expectedDelegURL  string
	}{
		{
			name:              "production URLs",
			organizationURL:   "https://onfleet.com/api/v2/organization",
			organizationsURL:  "https://onfleet.com/api/v2/organizations",
			expectedOrgURL:    "/organization",
			expectedDelegURL:  "/organizations/delegate_123",
		},
		{
			name:              "staging URLs",
			organizationURL:   "https://staging.onfleet.com/api/v2/organization",
			organizationsURL:  "https://staging.onfleet.com/api/v2/organizations",
			expectedOrgURL:    "/organization",
			expectedDelegURL:  "/organizations/delegate_123",
		},
		{
			name:              "custom URLs",
			organizationURL:   "https://custom.example.com/api/v1/organization",
			organizationsURL:  "https://custom.example.com/api/v1/organizations",
			expectedOrgURL:    "/organization",
			expectedDelegURL:  "/organizations/delegate_123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedOrg := testingutil.GetSampleOrganization()
			expectedDelegate := onfleet.OrganizationDelegate{
				ID:   "delegate_123",
				Name: "Test Delegate",
			}

			mockClient.AddResponse(tt.expectedOrgURL, testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedOrg,
			})
			mockClient.AddResponse(tt.expectedDelegURL, testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedDelegate,
			})

			client := Plug("test_api_key", nil, tt.organizationURL, tt.organizationsURL, mockClient.MockCaller)

			// Test Get
			org, err := client.Get()
			assert.NoError(t, err)
			assert.Equal(t, expectedOrg.ID, org.ID)

			// Reset mock client responses for second call
			mockClient.Reset()
			mockClient.AddResponse(tt.expectedDelegURL, testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedDelegate,
			})

			// Test GetDelegate
			delegate, err := client.GetDelegate("delegate_123")
			assert.NoError(t, err)
			assert.Equal(t, expectedDelegate.ID, delegate.ID)
		})
	}
}