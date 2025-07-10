package manifest

import (
	"testing"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_Generate(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedManifest := testingutil.GetSampleDeliveryManifest()
	mockClient.AddResponse("/providers/manifest", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedManifest,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/providers/manifest", mockClient.MockCaller)

	params := &onfleet.ManifestGenerateParams{
		HubId:    "hub_123",
		WorkerId: "worker_456",
	}

	manifest, err := client.Generate(params, "")

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedManifest.DepartureTime, manifest.DepartureTime)
	testingutil.AssertEqual(t, expectedManifest.Driver.Name, manifest.Driver.Name)
	testingutil.AssertEqual(t, expectedManifest.Driver.Phone, manifest.Driver.Phone)
	testingutil.AssertEqual(t, expectedManifest.HubAddress, manifest.HubAddress)
	testingutil.AssertEqual(t, expectedManifest.TotalDistance, manifest.TotalDistance)
	testingutil.AssertLen(t, manifest.Tasks, 1)
	testingutil.AssertLen(t, manifest.TurnByTurn, 1)

	mockClient.AssertRequestMade("POST", "/providers/manifest")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_GenerateWithGoogleAPIKey(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedManifest := testingutil.GetSampleDeliveryManifest()
	// Enhanced manifest with Google API integration features
	expectedManifest.TurnByTurn = []onfleet.TurnByTurn{
		{
			DrivingDistance: "2.5 miles",
			EndAddress:      "456 Oak Avenue, San Francisco, CA 94103",
			ETA:             1640995800,
			StartAddress:    "123 Main Street, San Francisco, CA 94105",
			Steps: []string{
				"Head north on Main St for 0.5 miles",
				"Turn right onto Market St for 1.2 miles",
				"Turn left onto Oak Ave for 0.8 miles",
				"Destination will be on the right",
			},
		},
		{
			DrivingDistance: "3.1 miles",
			EndAddress:      "789 Pine Street, San Francisco, CA 94108",
			ETA:             1640996400,
			StartAddress:    "456 Oak Avenue, San Francisco, CA 94103",
			Steps: []string{
				"Head east on Oak Ave for 0.3 miles",
				"Turn left onto Van Ness Ave for 2.1 miles",
				"Turn right onto Pine St for 0.7 miles",
				"Destination will be on the left",
			},
		},
	}

	mockClient.AddResponse("/providers/manifest", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedManifest,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/providers/manifest", mockClient.MockCaller)

	params := &onfleet.ManifestGenerateParams{
		HubId:    "hub_123",
		WorkerId: "worker_456",
	}

	googleAPIKey := "AIzaSyABCDEF123456789"
	manifest, err := client.Generate(params, googleAPIKey)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedManifest.DepartureTime, manifest.DepartureTime)
	testingutil.AssertLen(t, manifest.TurnByTurn, 2)
	testingutil.AssertEqual(t, "Head north on Main St for 0.5 miles", manifest.TurnByTurn[0].Steps[0])

	mockClient.AssertRequestMade("POST", "/providers/manifest")

	// Verify that the Google API key was included in headers
	lastRequest := mockClient.GetLastRequest()
	expectedHeaderValue := "Google " + googleAPIKey
	actualHeaderValue := lastRequest.Header.Get("X-API-Key")
	testingutil.AssertEqual(t, expectedHeaderValue, actualHeaderValue)
}

func TestClient_ManifestWithDifferentVehicleTypes(t *testing.T) {
	tests := []struct {
		name        string
		vehicleType onfleet.WorkerVehicleType
		description string
		capacity    string
	}{
		{
			name:        "car manifest",
			vehicleType: onfleet.WorkerVehicleTypeCar,
			description: "Red Toyota Camry",
			capacity:    "4 packages",
		},
		{
			name:        "truck manifest",
			vehicleType: onfleet.WorkerVehicleTypeTruck,
			description: "Blue Freightliner",
			capacity:    "100 packages",
		},
		{
			name:        "motorcycle manifest",
			vehicleType: onfleet.WorkerVehicleTypeMotorcycle,
			description: "Black Yamaha",
			capacity:    "2 packages",
		},
		{
			name:        "bicycle manifest",
			vehicleType: onfleet.WorkerVehicleTypeBicycle,
			description: "Green Trek Bike",
			capacity:    "1 package",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedManifest := testingutil.GetSampleDeliveryManifest()
			expectedManifest.Vehicle = onfleet.WorkerVehicleParam{
				Type:         tt.vehicleType,
				Description:  tt.description,
				LicensePlate: "VEH123",
			}

			mockClient.AddResponse("/providers/manifest", testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedManifest,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/providers/manifest", mockClient.MockCaller)

			params := &onfleet.ManifestGenerateParams{
				HubId:    "hub_123",
				WorkerId: "worker_456",
			}

			manifest, err := client.Generate(params, "")

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.vehicleType, manifest.Vehicle.Type)
			testingutil.AssertEqual(t, tt.description, manifest.Vehicle.Description)
		})
	}
}

func TestClient_ManifestWithMultipleTasks(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedManifest := testingutil.GetSampleDeliveryManifest()
	// Add multiple tasks to simulate a full delivery route
	sampleTask := testingutil.GetSampleTask()
	expectedManifest.Tasks = []onfleet.Task{
		sampleTask,
		{
			ID:      "task_456",
			ShortId: "def456",
			Destination: onfleet.Destination{
				ID:       "dest_456",
				Location: onfleet.DestinationLocation{-122.4194, 37.7849},
				Address: onfleet.DestinationAddress{
					Street:     "456 Market Street",
					City:       "San Francisco",
					State:      "CA",
					PostalCode: "94105",
				},
			},
			State: onfleet.TaskStateUnassigned,
		},
		{
			ID:      "task_789",
			ShortId: "ghi789",
			Destination: onfleet.Destination{
				ID:       "dest_789",
				Location: onfleet.DestinationLocation{-122.4294, 37.7949},
				Address: onfleet.DestinationAddress{
					Street:     "789 Mission Street",
					City:       "San Francisco",
					State:      "CA",
					PostalCode: "94103",
				},
			},
			State: onfleet.TaskStateUnassigned,
		},
	}

	expectedManifest.TurnByTurn = []onfleet.TurnByTurn{
		{
			DrivingDistance: "2.5 miles",
			EndAddress:      "456 Market Street, San Francisco, CA 94105",
			ETA:             1640995800,
			StartAddress:    "123 Main Street, San Francisco, CA 94105",
			Steps:           []string{"Head north", "Turn right", "Arrive at destination"},
		},
		{
			DrivingDistance: "1.8 miles",
			EndAddress:      "789 Mission Street, San Francisco, CA 94103",
			ETA:             1640996400,
			StartAddress:    "456 Market Street, San Francisco, CA 94105",
			Steps:           []string{"Head south", "Turn left", "Arrive at destination"},
		},
	}

	mockClient.AddResponse("/providers/manifest", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedManifest,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/providers/manifest", mockClient.MockCaller)

	params := &onfleet.ManifestGenerateParams{
		HubId:    "hub_123",
		WorkerId: "worker_456",
	}

	manifest, err := client.Generate(params, "")

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, manifest.Tasks, 3)
	testingutil.AssertLen(t, manifest.TurnByTurn, 2)
	testingutil.AssertEqual(t, "task_456", manifest.Tasks[1].ID)
	testingutil.AssertEqual(t, "task_789", manifest.Tasks[2].ID)

	mockClient.AssertRequestMade("POST", "/providers/manifest")
}

func TestClient_ManifestDriverInformation(t *testing.T) {
	tests := []struct {
		name        string
		driverName  string
		driverPhone string
	}{
		{
			name:        "driver with name and phone",
			driverName:  "Alice Johnson",
			driverPhone: "+15559876543",
		},
		{
			name:        "driver with international phone",
			driverName:  "Roberto Martinez",
			driverPhone: "+34123456789",
		},
		{
			name:        "driver with extension",
			driverName:  "David Chen",
			driverPhone: "+15551234567x101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedManifest := testingutil.GetSampleDeliveryManifest()
			expectedManifest.Driver = onfleet.Driver{
				Name:  tt.driverName,
				Phone: tt.driverPhone,
			}

			mockClient.AddResponse("/providers/manifest", testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedManifest,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/providers/manifest", mockClient.MockCaller)

			params := &onfleet.ManifestGenerateParams{
				HubId:    "hub_123",
				WorkerId: "worker_456",
			}

			manifest, err := client.Generate(params, "")

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.driverName, manifest.Driver.Name)
			testingutil.AssertEqual(t, tt.driverPhone, manifest.Driver.Phone)
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
			name:       "generate manifest worker not found",
			method:     "POST",
			url:        "/providers/manifest",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Generate(&onfleet.ManifestGenerateParams{
					HubId:    "hub_123",
					WorkerId: "nonexistent_worker",
				}, "")
				return err
			},
		},
		{
			name:       "generate manifest hub not found",
			method:     "POST",
			url:        "/providers/manifest",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Generate(&onfleet.ManifestGenerateParams{
					HubId:    "nonexistent_hub",
					WorkerId: "worker_456",
				}, "")
				return err
			},
		},
		{
			name:       "generate manifest missing parameters",
			method:     "POST",
			url:        "/providers/manifest",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Generate(&onfleet.ManifestGenerateParams{
					// Missing required parameters
				}, "")
				return err
			},
		},
		{
			name:       "generate manifest invalid google api key",
			method:     "POST",
			url:        "/providers/manifest",
			statusCode: 401,
			operation: func(client *Client) error {
				_, err := client.Generate(&onfleet.ManifestGenerateParams{
					HubId:    "hub_123",
					WorkerId: "worker_456",
				}, "invalid_google_key")
				return err
			},
		},
		{
			name:       "generate manifest no tasks assigned",
			method:     "POST",
			url:        "/providers/manifest",
			statusCode: 422,
			operation: func(client *Client) error {
				_, err := client.Generate(&onfleet.ManifestGenerateParams{
					HubId:    "hub_123",
					WorkerId: "worker_without_tasks",
				}, "")
				return err
			},
		},
		{
			name:       "generate manifest unauthorized",
			method:     "POST",
			url:        "/providers/manifest",
			statusCode: 403,
			operation: func(client *Client) error {
				_, err := client.Generate(&onfleet.ManifestGenerateParams{
					HubId:    "hub_123",
					WorkerId: "worker_456",
				}, "")
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

			client := Plug("test_api_key", nil, "https://api.example.com/providers/manifest", mockClient.MockCaller)

			err := tt.operation(client)
			testingutil.AssertError(t, err)
		})
	}
}