package hub

import (
	"testing"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_List(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedHubs := []onfleet.Hub{
		testingutil.GetSampleHub(),
	}

	mockClient.AddResponse("/hubs", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedHubs,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/hubs", mockClient.MockCaller)

	hubs, err := client.List()

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, hubs, 1)
	testingutil.AssertEqual(t, expectedHubs[0].ID, hubs[0].ID)
	testingutil.AssertEqual(t, expectedHubs[0].Name, hubs[0].Name)
	testingutil.AssertEqual(t, expectedHubs[0].Address.City, hubs[0].Address.City)

	mockClient.AssertRequestMade("GET", "/hubs")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_Create(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedHub := testingutil.GetSampleHub()
	mockClient.AddResponse("/hubs", testingutil.MockResponse{
		StatusCode: 201,
		Body:       expectedHub,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/hubs", mockClient.MockCaller)

	params := onfleet.HubCreateParams{
		Name: "New Distribution Center",
		Address: onfleet.DestinationAddress{
			Street:     "456 Industrial Way",
			City:       "Oakland",
			State:      "CA",
			PostalCode: "94607",
			Country:    "US",
		},
		Teams: []string{"team_789"},
	}

	hub, err := client.Create(params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedHub.ID, hub.ID)
	testingutil.AssertEqual(t, expectedHub.Name, hub.Name)

	mockClient.AssertRequestMade("POST", "/hubs")
}

func TestClient_Update(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedHub := testingutil.GetSampleHub()
	expectedHub.Name = "Updated Distribution Center"

	mockClient.AddResponse("/hubs/hub_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedHub,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/hubs", mockClient.MockCaller)

	params := onfleet.HubUpdateParams{
		Name: "Updated Distribution Center",
		Address: onfleet.DestinationAddress{
			Street:     "789 Updated Street",
			City:       "San Francisco",
			State:      "CA",
			PostalCode: "94105",
			Country:    "US",
		},
		Teams: []string{"team_123", "team_456", "team_789"},
	}

	hub, err := client.Update("hub_123", params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedHub.ID, hub.ID)
	testingutil.AssertEqual(t, "Updated Distribution Center", hub.Name)

	mockClient.AssertRequestMade("PUT", "/hubs/hub_123")
}

func TestClient_AddressValidation(t *testing.T) {
	tests := []struct {
		name    string
		address onfleet.DestinationAddress
	}{
		{
			name: "US address",
			address: onfleet.DestinationAddress{
				Street:     "123 Main Street",
				City:       "San Francisco",
				State:      "CA",
				PostalCode: "94105",
				Country:    "US",
			},
		},
		{
			name: "International address",
			address: onfleet.DestinationAddress{
				Street:     "10 Downing Street",
				City:       "London",
				State:      "",
				PostalCode: "SW1A 2AA",
				Country:    "UK",
			},
		},
		{
			name: "Address with apartment",
			address: onfleet.DestinationAddress{
				Apartment:  "Suite 100",
				Street:     "456 Business Blvd",
				City:       "New York",
				State:      "NY",
				PostalCode: "10001",
				Country:    "US",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedHub := testingutil.GetSampleHub()
			expectedHub.Address = tt.address

			mockClient.AddResponse("/hubs", testingutil.MockResponse{
				StatusCode: 201,
				Body:       expectedHub,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/hubs", mockClient.MockCaller)

			params := onfleet.HubCreateParams{
				Name:    "Test Hub",
				Address: tt.address,
			}

			hub, err := client.Create(params)

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.address.Street, hub.Address.Street)
			testingutil.AssertEqual(t, tt.address.City, hub.Address.City)
			testingutil.AssertEqual(t, tt.address.Country, hub.Address.Country)
		})
	}
}

func TestClient_TeamAssignment(t *testing.T) {
	tests := []struct {
		name  string
		teams []string
	}{
		{
			name:  "single team",
			teams: []string{"team_123"},
		},
		{
			name:  "multiple teams",
			teams: []string{"team_123", "team_456", "team_789"},
		},
		{
			name:  "no teams",
			teams: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedHub := testingutil.GetSampleHub()
			expectedHub.Teams = tt.teams

			mockClient.AddResponse("/hubs", testingutil.MockResponse{
				StatusCode: 201,
				Body:       expectedHub,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/hubs", mockClient.MockCaller)

			params := onfleet.HubCreateParams{
				Name: "Team Assignment Test Hub",
				Address: onfleet.DestinationAddress{
					Street:  "123 Test Street",
					City:    "Test City",
					State:   "CA",
					Country: "US",
				},
				Teams: tt.teams,
			}

			hub, err := client.Create(params)

			testingutil.AssertNoError(t, err)
			testingutil.AssertLen(t, hub.Teams, len(tt.teams))
			if len(tt.teams) > 0 {
				testingutil.AssertEqual(t, tt.teams[0], hub.Teams[0])
			}
		})
	}
}

func TestClient_ErrorScenarios(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
		operation  func(client *Client) error
	}{
		{
			name:       "create invalid address",
			method:     "POST",
			url:        "/hubs",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.HubCreateParams{
					Name: "Invalid Hub",
					Address: onfleet.DestinationAddress{
						Street: "", // Invalid empty street
					},
				})
				return err
			},
		},
		{
			name:       "create empty name",
			method:     "POST",
			url:        "/hubs",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.HubCreateParams{
					Name: "", // Invalid empty name
					Address: onfleet.DestinationAddress{
						Street:  "123 Test Street",
						City:    "Test City",
						Country: "US",
					},
				})
				return err
			},
		},
		{
			name:       "update not found",
			method:     "PUT",
			url:        "/hubs/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Update("nonexistent", onfleet.HubUpdateParams{
					Name: "Updated Hub",
					Address: onfleet.DestinationAddress{
						Street:  "123 Test Street",
						City:    "Test City",
						Country: "US",
					},
				})
				return err
			},
		},
		{
			name:       "list unauthorized",
			method:     "GET",
			url:        "/hubs",
			statusCode: 401,
			operation: func(client *Client) error {
				_, err := client.List()
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			mockClient.AddResponse(tt.url, testingutil.MockResponse{
				StatusCode: tt.statusCode,
				Body:       testingutil.GetSampleErrorResponse(),
			})

			client := Plug("test_api_key", nil, "https://api.example.com/hubs", mockClient.MockCaller)

			err := tt.operation(client)
			testingutil.AssertError(t, err)
		})
	}
}