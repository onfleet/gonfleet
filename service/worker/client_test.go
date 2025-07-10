package worker

import (
	"testing"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_Get(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWorker := testingutil.GetSampleWorker()
	mockClient.AddResponse("/workers/worker_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedWorker,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	worker, err := client.Get("worker_123")

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedWorker.ID, worker.ID)
	testingutil.AssertEqual(t, expectedWorker.Name, worker.Name)
	testingutil.AssertEqual(t, expectedWorker.Phone, worker.Phone)

	mockClient.AssertRequestMade("GET", "/workers/worker_123")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_Get_NotFound(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/workers/nonexistent", testingutil.MockResponse{
		StatusCode: 404,
		Body:       testingutil.GetSampleErrorResponse(),
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	worker, err := client.Get("nonexistent")

	testingutil.AssertError(t, err)
	testingutil.AssertEqual(t, "", worker.ID)
}

func TestClient_GetWithQuery(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := map[string]any{
		"id":   "worker_123",
		"name": "John Doe",
		"analytics": map[string]any{
			"distances": map[string]any{
				"enroute": 25.5,
				"idle":    5.2,
			},
		},
	}

	mockClient.AddResponse("/workers/worker_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	params := onfleet.WorkerGetQueryParams{
		Analytics: true,
		From:      1640995200,
		To:        1672531199,
	}

	response, err := client.GetWithQuery("worker_123", params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, "worker_123", response["id"])
	testingutil.AssertNotNil(t, response["analytics"])

	mockClient.AssertRequestMade("GET", "/workers/worker_123")
}

func TestClient_List(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWorkers := []onfleet.Worker{
		testingutil.GetSampleWorker(),
	}

	mockClient.AddResponse("/workers", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedWorkers,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	workers, err := client.List()

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, workers, 1)
	testingutil.AssertEqual(t, expectedWorkers[0].ID, workers[0].ID)

	mockClient.AssertRequestMade("GET", "/workers")
}

func TestClient_ListWithQuery(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := []map[string]any{
		{
			"id":   "worker_123",
			"name": "John Doe",
		},
	}

	mockClient.AddResponse("/workers", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	params := onfleet.WorkerListQueryParams{
		Filter: "active",
		Teams:  "team_123,team_456",
	}

	response, err := client.ListWithQuery(params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, response, 1)
	testingutil.AssertEqual(t, "worker_123", response[0]["id"])

	mockClient.AssertRequestMade("GET", "/workers")
}

func TestClient_ListWithMetadataQuery(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWorkers := []onfleet.Worker{
		testingutil.GetSampleWorker(),
	}

	mockClient.AddResponse("/workers/metadata", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedWorkers,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	metadata := []onfleet.Metadata{
		{
			Name:  "employee_id",
			Type:  "string",
			Value: "EMP001",
		},
	}

	workers, err := client.ListWithMetadataQuery(metadata)

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, workers, 1)
	testingutil.AssertEqual(t, expectedWorkers[0].ID, workers[0].ID)

	mockClient.AssertRequestMade("POST", "/workers/metadata")
}

func TestClient_Create(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWorker := testingutil.GetSampleWorker()
	mockClient.AddResponse("/workers", testingutil.MockResponse{
		StatusCode: 201,
		Body:       expectedWorker,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	params := testingutil.GetSampleWorkerCreateParams()

	worker, err := client.Create(params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedWorker.ID, worker.ID)
	testingutil.AssertEqual(t, expectedWorker.Name, worker.Name)

	mockClient.AssertRequestMade("POST", "/workers")
}

func TestClient_Create_ValidationError(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/workers", testingutil.MockResponse{
		StatusCode: 400,
		Body:       testingutil.GetSampleValidationErrorResponse(),
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	// Invalid params - missing required fields
	params := onfleet.WorkerCreateParams{
		// Missing name, phone, teams
	}

	worker, err := client.Create(params)

	testingutil.AssertError(t, err)
	testingutil.AssertEqual(t, "", worker.ID)
}

func TestClient_Update(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWorker := testingutil.GetSampleWorker()
	expectedWorker.Name = "Updated Name"

	mockClient.AddResponse("/workers/worker_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedWorker,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	params := onfleet.WorkerUpdateParams{
		Name: "Updated Name",
	}

	worker, err := client.Update("worker_123", params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedWorker.ID, worker.ID)
	testingutil.AssertEqual(t, "Updated Name", worker.Name)

	mockClient.AssertRequestMade("PUT", "/workers/worker_123")
}

func TestClient_Delete(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/workers/worker_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       map[string]interface{}{"success": true},
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	err := client.Delete("worker_123")

	testingutil.AssertNoError(t, err)
	mockClient.AssertRequestMade("DELETE", "/workers/worker_123")
}

func TestClient_GetSchedule(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedSchedule := onfleet.WorkerScheduleEntries{
		Entries: []onfleet.WorkerSchedule{
			{
				Date:     "2023-01-01",
				Shifts:   [][]int64{{32400, 61200}}, // 9 AM to 5 PM
				Timezone: "America/Los_Angeles",
			},
		},
	}

	mockClient.AddResponse("/workers/worker_123/schedule", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedSchedule,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	schedule, err := client.GetSchedule("worker_123")

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, schedule.Entries, 1)
	testingutil.AssertEqual(t, "2023-01-01", schedule.Entries[0].Date)

	mockClient.AssertRequestMade("GET", "/workers/worker_123/schedule")
}

func TestClient_SetSchedule(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedSchedule := onfleet.WorkerScheduleEntries{
		Entries: []onfleet.WorkerSchedule{
			{
				Date:     "2023-01-02",
				Shifts:   [][]int64{{28800, 64800}}, // 8 AM to 6 PM
				Timezone: "America/Los_Angeles",
			},
		},
	}

	mockClient.AddResponse("/workers/worker_123/schedule", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedSchedule,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	inputSchedule := onfleet.WorkerScheduleEntries{
		Entries: []onfleet.WorkerSchedule{
			{
				Date:     "2023-01-02",
				Shifts:   [][]int64{{28800, 64800}},
				Timezone: "America/Los_Angeles",
			},
		},
	}

	schedule, err := client.SetSchedule("worker_123", inputSchedule)

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, schedule.Entries, 1)
	testingutil.AssertEqual(t, "2023-01-02", schedule.Entries[0].Date)

	mockClient.AssertRequestMade("POST", "/workers/worker_123/schedule")
}

func TestClient_ListWorkersByLocation(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := onfleet.WorkersByLocation{
		Workers: []onfleet.Worker{
			testingutil.GetSampleWorker(),
		},
	}

	mockClient.AddResponse("/workers/location", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	params := onfleet.WorkersByLocationListQueryParams{
		Longitude: -122.4194,
		Latitude:  37.7749,
		Radius:    5000, // 5km radius
	}

	response, err := client.ListWorkersByLocation(params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, response.Workers, 1)
	testingutil.AssertEqual(t, expectedResponse.Workers[0].ID, response.Workers[0].ID)

	mockClient.AssertRequestMade("GET", "/workers/location")
}

func TestClient_ListTasks(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := onfleet.WorkerTasks{
		Tasks: []onfleet.Task{
			testingutil.GetSampleTask(),
		},
		LastId: "last_task_456",
	}

	mockClient.AddResponse("/workers/worker_123/tasks", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	params := &onfleet.WorkerTasksListQueryParams{
		From:         1640995200,
		To:           1672531199,
		IsPickupTask: "false",
	}

	response, err := client.ListTasks("worker_123", params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, response.Tasks, 1)
	testingutil.AssertEqual(t, "last_task_456", response.LastId)

	mockClient.AssertRequestMade("GET", "/workers/worker_123/tasks")
}

func TestClient_ListTasks_NilParams(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := onfleet.WorkerTasks{
		Tasks:  []onfleet.Task{},
		LastId: "",
	}

	mockClient.AddResponse("/workers/worker_123/tasks", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

	response, err := client.ListTasks("worker_123", nil)

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, response.Tasks, 0)

	mockClient.AssertRequestMade("GET", "/workers/worker_123/tasks")
}

// Table-driven test for different worker account statuses
func TestClient_Get_DifferentAccountStatuses(t *testing.T) {
	tests := []struct {
		name   string
		status onfleet.WorkerAccountStatus
	}{
		{"accepted worker", onfleet.WorkerAccountStatusAccepted},
		{"invited worker", onfleet.WorkerAccountStatusInvited},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedWorker := testingutil.GetSampleWorker()
			expectedWorker.AccountStatus = tt.status

			mockClient.AddResponse("/workers/worker_123", testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedWorker,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

			worker, err := client.Get("worker_123")

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.status, worker.AccountStatus)
		})
	}
}

// Table-driven test for different vehicle types
func TestClient_Create_DifferentVehicleTypes(t *testing.T) {
	tests := []struct {
		name        string
		vehicleType onfleet.WorkerVehicleType
	}{
		{"car worker", onfleet.WorkerVehicleTypeCar},
		{"bicycle worker", onfleet.WorkerVehicleTypeBicycle},
		{"motorcycle worker", onfleet.WorkerVehicleTypeMotorcycle},
		{"truck worker", onfleet.WorkerVehicleTypeTruck},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedWorker := testingutil.GetSampleWorker()
			expectedWorker.Vehicle.Type = tt.vehicleType

			mockClient.AddResponse("/workers", testingutil.MockResponse{
				StatusCode: 201,
				Body:       expectedWorker,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

			params := testingutil.GetSampleWorkerCreateParams()
			params.Vehicle.Type = tt.vehicleType

			worker, err := client.Create(params)

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.vehicleType, worker.Vehicle.Type)
		})
	}
}

// Test different query parameters
func TestClient_ListWithQuery_DifferentFilters(t *testing.T) {
	tests := []struct {
		name   string
		params onfleet.WorkerListQueryParams
	}{
		{
			name: "filter by state",
			params: onfleet.WorkerListQueryParams{
				States: "0,1", // Unassigned and assigned
			},
		},
		{
			name: "filter by teams",
			params: onfleet.WorkerListQueryParams{
				Teams: "team_123,team_456",
			},
		},
		{
			name: "filter by phones",
			params: onfleet.WorkerListQueryParams{
				Phones: "+15551234567,+15559876543",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedResponse := []map[string]any{
				{
					"id":   "worker_123",
					"name": "John Doe",
				},
			}

			mockClient.AddResponse("/workers", testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedResponse,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

			response, err := client.ListWithQuery(tt.params)

			testingutil.AssertNoError(t, err)
			testingutil.AssertLen(t, response, 1)
		})
	}
}

// Test error scenarios
func TestClient_ErrorScenarios(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{"get unauthorized", "GET", "/workers/worker_123", 401},
		{"create forbidden", "POST", "/workers", 403},
		{"update not found", "PUT", "/workers/nonexistent", 404},
		{"delete server error", "DELETE", "/workers/worker_123", 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			mockClient.AddResponse(tt.url, testingutil.MockResponse{
				StatusCode: tt.statusCode,
				Body:       testingutil.GetSampleErrorResponse(),
			})

			client := Plug("test_api_key", nil, "https://api.example.com/workers", mockClient.MockCaller)

			var err error
			switch tt.method {
			case "GET":
				_, err = client.Get("worker_123")
			case "POST":
				_, err = client.Create(testingutil.GetSampleWorkerCreateParams())
			case "PUT":
				_, err = client.Update("nonexistent", onfleet.WorkerUpdateParams{})
			case "DELETE":
				err = client.Delete("worker_123")
			}

			testingutil.AssertError(t, err)
		})
	}
}