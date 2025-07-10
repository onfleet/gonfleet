package routePlan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_Create(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedRoutePlan := testingutil.GetSampleRoutePlan()
	mockClient.AddResponse("/route-plans", testingutil.MockResponse{
		StatusCode: 201,
		Body:       expectedRoutePlan,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

	params := onfleet.RoutePlanParams{
		Name:          "Evening Delivery Route",
		StartTime:     1641031200,
		TaskIds:       []string{"task_444", "task_555"},
		Color:         "#33A1FF",
		VehicleType:   "VAN",
		Worker:        "worker_456",
		Team:          "team_123",
		StartAt:       onfleet.PositionEnumHub,
		EndAt:         onfleet.PositionEnumHub,
		StartingHubId: "hub_123",
		EndingHubId:   "hub_456",
		EndTime:       1641060000,
		Timezone:      "America/Los_Angeles",
	}

	routePlan, err := client.Create(params)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoutePlan.Id, routePlan.Id)
	assert.Equal(t, expectedRoutePlan.Name, routePlan.Name)
	assert.Equal(t, expectedRoutePlan.VehicleType, routePlan.VehicleType)

	mockClient.AssertRequestMade("POST", "/route-plans")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_Get(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedRoutePlan := testingutil.GetSampleRoutePlan()
	mockClient.AddResponse("/route-plans/routeplan_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedRoutePlan,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

	routePlan, err := client.Get("routeplan_123")

	assert.NoError(t, err)
	assert.Equal(t, expectedRoutePlan.Id, routePlan.Id)
	assert.Equal(t, expectedRoutePlan.Name, routePlan.Name)
	assert.Equal(t, expectedRoutePlan.State, routePlan.State)
	assert.Len(t, routePlan.Tasks, 3)

	mockClient.AssertRequestMade("GET", "/route-plans/routeplan_123")
}

func TestClient_Update(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedRoutePlan := testingutil.GetSampleRoutePlan()
	expectedRoutePlan.Name = "Updated Morning Route"
	expectedRoutePlan.Color = "#00FF00"

	mockClient.AddResponse("/route-plans/routeplan_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedRoutePlan,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

	params := onfleet.RoutePlanParams{
		Name:        "Updated Morning Route",
		Color:       "#00FF00",
		VehicleType: "TRUCK",
		StartTime:   1640995200,
	}

	routePlan, err := client.Update("routeplan_123", params)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoutePlan.Id, routePlan.Id)
	assert.Equal(t, "Updated Morning Route", routePlan.Name)
	assert.Equal(t, "#00FF00", routePlan.Color)

	mockClient.AssertRequestMade("PUT", "/route-plans/routeplan_123")
}

func TestClient_AddTasks(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedRoutePlan := testingutil.GetSampleRoutePlan()
	expectedRoutePlan.Tasks = []string{"task_111", "task_222", "task_333", "task_444", "task_555"}

	mockClient.AddResponse("/route-plans/routeplan_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedRoutePlan,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

	params := onfleet.RoutePlanAddTasksParams{
		Tasks: []string{"task_444", "task_555"},
	}

	routePlan, err := client.AddTasks("routeplan_123", params)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoutePlan.Id, routePlan.Id)
	assert.Len(t, routePlan.Tasks, 5)
	assert.Equal(t, "task_444", routePlan.Tasks[3])
	assert.Equal(t, "task_555", routePlan.Tasks[4])

	mockClient.AssertRequestMade("PUT", "/route-plans/routeplan_123")
}

func TestClient_List(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedRoutePlans := []onfleet.RoutePlan{
		testingutil.GetSampleRoutePlan(),
	}

	expectedResponse := onfleet.RoutePlansPaginated{
		RoutePlans: expectedRoutePlans,
		LastId:     "routeplan_123",
	}

	mockClient.AddResponse("/route-plans/all", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

	params := onfleet.RoutePlanListQueryParams{
		WorkerId:        "worker_123",
		StartTimeFrom:   1640995200,
		StartTimeTo:     1641081600,
		CreatedTimeFrom: 1640990000,
		CreatedTimeTo:   1641000000,
		HasTasks:        true,
		Limit:           10,
	}

	response, err := client.List(params)

	assert.NoError(t, err)
	assert.Len(t, response.RoutePlans, 1)
	assert.Equal(t, expectedRoutePlans[0].Id, response.RoutePlans[0].Id)
	assert.Equal(t, "routeplan_123", response.LastId)

	mockClient.AssertRequestMade("GET", "/route-plans/all")
}

func TestClient_Delete(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/route-plans/routeplan_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       map[string]interface{}{"success": true},
	})

	client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

	err := client.Delete("routeplan_123")

	assert.NoError(t, err)
	mockClient.AssertRequestMade("DELETE", "/route-plans/routeplan_123")
}

func TestClient_VehicleTypes(t *testing.T) {
	tests := []struct {
		name        string
		vehicleType string
	}{
		{
			name:        "car vehicle",
			vehicleType: "CAR",
		},
		{
			name:        "van vehicle",
			vehicleType: "VAN",
		},
		{
			name:        "truck vehicle",
			vehicleType: "TRUCK",
		},
		{
			name:        "motorcycle vehicle",
			vehicleType: "MOTORCYCLE",
		},
		{
			name:        "bicycle vehicle",
			vehicleType: "BICYCLE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedRoutePlan := testingutil.GetSampleRoutePlan()
			expectedRoutePlan.VehicleType = tt.vehicleType

			mockClient.AddResponse("/route-plans", testingutil.MockResponse{
				StatusCode: 201,
				Body:       expectedRoutePlan,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

			params := onfleet.RoutePlanParams{
				Name:        "Vehicle Type Test Route",
				StartTime:   1640995200,
				VehicleType: tt.vehicleType,
				Worker:      "worker_123",
			}

			routePlan, err := client.Create(params)

			assert.NoError(t, err)
			assert.Equal(t, tt.vehicleType, routePlan.VehicleType)
		})
	}
}

func TestClient_PositionTypes(t *testing.T) {
	tests := []struct {
		name     string
		startAt  onfleet.PositionEnum
		endAt    onfleet.PositionEnum
		startHub string
		endHub   string
	}{
		{
			name:     "hub to hub",
			startAt:  onfleet.PositionEnumHub,
			endAt:    onfleet.PositionEnumHub,
			startHub: "hub_123",
			endHub:   "hub_456",
		},
		{
			name:     "worker location to hub",
			startAt:  onfleet.PositionEnumWorkerLocation,
			endAt:    onfleet.PositionEnumHub,
			startHub: "",
			endHub:   "hub_456",
		},
		{
			name:     "worker address to worker location",
			startAt:  onfleet.PositionEnumWorkerAddress,
			endAt:    onfleet.PositionEnumWorkerLocation,
			startHub: "",
			endHub:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedRoutePlan := testingutil.GetSampleRoutePlan()
			if tt.startHub != "" {
				expectedRoutePlan.StartingHubId = &tt.startHub
			} else {
				expectedRoutePlan.StartingHubId = nil
			}
			if tt.endHub != "" {
				expectedRoutePlan.EndingHubId = &tt.endHub
			} else {
				expectedRoutePlan.EndingHubId = nil
			}

			mockClient.AddResponse("/route-plans", testingutil.MockResponse{
				StatusCode: 201,
				Body:       expectedRoutePlan,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

			params := onfleet.RoutePlanParams{
				Name:          "Position Test Route",
				StartTime:     1640995200,
				Worker:        "worker_123",
				StartAt:       tt.startAt,
				EndAt:         tt.endAt,
				StartingHubId: tt.startHub,
				EndingHubId:   tt.endHub,
			}

			routePlan, err := client.Create(params)

			assert.NoError(t, err)
			assert.Equal(t, expectedRoutePlan.Id, routePlan.Id)
			if tt.startHub != "" {
				assert.Equal(t, tt.startHub, *routePlan.StartingHubId)
			}
			if tt.endHub != "" {
				assert.Equal(t, tt.endHub, *routePlan.EndingHubId)
			}
		})
	}
}

func TestClient_RoutePlanStates(t *testing.T) {
	tests := []struct {
		name  string
		state string
	}{
		{
			name:  "active route plan",
			state: "active",
		},
		{
			name:  "completed route plan",
			state: "completed",
		},
		{
			name:  "draft route plan",
			state: "draft",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedRoutePlan := testingutil.GetSampleRoutePlan()
			expectedRoutePlan.State = tt.state

			mockClient.AddResponse("/route-plans/routeplan_123", testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedRoutePlan,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

			routePlan, err := client.Get("routeplan_123")

			assert.NoError(t, err)
			assert.Equal(t, tt.state, routePlan.State)
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
			name:       "create invalid start time",
			method:     "POST",
			url:        "/route-plans",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.RoutePlanParams{
					Name:      "Invalid Route",
					StartTime: 0, // Invalid start time
				})
				return err
			},
		},
		{
			name:       "create missing worker and team",
			method:     "POST",
			url:        "/route-plans",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.RoutePlanParams{
					Name:      "Missing Assignment Route",
					StartTime: 1640995200,
					// Missing both worker and team
				})
				return err
			},
		},
		{
			name:       "get not found",
			method:     "GET",
			url:        "/route-plans/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Get("nonexistent")
				return err
			},
		},
		{
			name:       "update not found",
			method:     "PUT",
			url:        "/route-plans/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Update("nonexistent", onfleet.RoutePlanParams{
					Name:      "Updated Route",
					StartTime: 1640995200,
				})
				return err
			},
		},
		{
			name:       "add tasks to completed route plan",
			method:     "PUT",
			url:        "/route-plans/completed_123",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.AddTasks("completed_123", onfleet.RoutePlanAddTasksParams{
					Tasks: []string{"task_123"},
				})
				return err
			},
		},
		{
			name:       "delete active route plan",
			method:     "DELETE",
			url:        "/route-plans/active_123",
			statusCode: 409,
			operation: func(client *Client) error {
				return client.Delete("active_123")
			},
		},
		{
			name:       "list unauthorized",
			method:     "GET",
			url:        "/route-plans/all",
			statusCode: 401,
			operation: func(client *Client) error {
				_, err := client.List(onfleet.RoutePlanListQueryParams{})
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

			client := Plug("test_api_key", nil, "https://api.example.com/route-plans", mockClient.MockCaller)

			err := tt.operation(client)
			assert.Error(t, err)
		})
	}
}