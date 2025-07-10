package destination

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_Get(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedDestination := testingutil.GetSampleDestination()
	mockClient.AddResponse("/destinations/destination_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedDestination,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/destinations", mockClient.MockCaller)

	destination, err := client.Get("destination_123")

	assert.NoError(t, err)
	assert.Equal(t, expectedDestination.ID, destination.ID)
	assert.Equal(t, expectedDestination.Address.Street, destination.Address.Street)

	mockClient.AssertRequestMade("GET", "/destinations/destination_123")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_Create(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedDestination := testingutil.GetSampleDestination()
	mockClient.AddResponse("/destinations", testingutil.MockResponse{
		StatusCode: 201,
		Body:       expectedDestination,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/destinations", mockClient.MockCaller)

	params := onfleet.DestinationCreateParams{
		Address: onfleet.DestinationAddress{
			Number:     "456",
			Street:     "Test Street",
			City:       "Test City",
			State:      "CA",
			PostalCode: "12345",
			Country:    "US",
		},
		Notes: "Test destination",
	}

	destination, err := client.Create(params)

	assert.NoError(t, err)
	assert.Equal(t, expectedDestination.ID, destination.ID)

	mockClient.AssertRequestMade("POST", "/destinations")
}

func TestClient_ListWithMetadataQuery(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedDestinations := []onfleet.Destination{
		testingutil.GetSampleDestination(),
	}

	mockClient.AddResponse("/destinations/metadata", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedDestinations,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/destinations", mockClient.MockCaller)

	metadata := []onfleet.Metadata{
		{
			Name:  "location_type",
			Type:  "string",
			Value: "warehouse",
		},
	}

	destinations, err := client.ListWithMetadataQuery(metadata)

	assert.NoError(t, err)
	assert.Len(t, destinations, 1)
	assert.Equal(t, expectedDestinations[0].ID, destinations[0].ID)

	mockClient.AssertRequestMade("POST", "/destinations/metadata")
}

func TestClient_ErrorScenarios(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{"get not found", "GET", "/destinations/nonexistent", 404},
		{"create invalid", "POST", "/destinations", 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			mockClient.AddResponse(tt.url, testingutil.MockResponse{
				StatusCode: tt.statusCode,
				Body:       testingutil.GetSampleErrorResponse(),
			})

			client := Plug("test_api_key", nil, "https://api.example.com/destinations", mockClient.MockCaller)

			var err error
			switch tt.method {
			case "GET":
				_, err = client.Get("nonexistent")
			case "POST":
				_, err = client.Create(onfleet.DestinationCreateParams{})
			}

			assert.Error(t, err)
		})
	}
}