package team

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_Get(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedTeam := testingutil.GetSampleTeam()
	mockClient.AddResponse("/teams/team_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedTeam,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	team, err := client.Get("team_123")

	assert.NoError(t, err)
	assert.Equal(t, expectedTeam.ID, team.ID)
	assert.Equal(t, expectedTeam.Name, team.Name)
	assert.Equal(t, expectedTeam.EnableSelfAssignment, team.EnableSelfAssignment)

	mockClient.AssertRequestMade("GET", "/teams/team_123")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_List(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedTeams := []onfleet.Team{
		testingutil.GetSampleTeam(),
	}

	mockClient.AddResponse("/teams", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedTeams,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	teams, err := client.List()

	assert.NoError(t, err)
	assert.Len(t, teams, 1)
	assert.Equal(t, expectedTeams[0].ID, teams[0].ID)

	mockClient.AssertRequestMade("GET", "/teams")
}

func TestClient_Create(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedTeam := testingutil.GetSampleTeam()
	mockClient.AddResponse("/teams", testingutil.MockResponse{
		StatusCode: 201,
		Body:       expectedTeam,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	params := onfleet.TeamCreateParams{
		Name:                 "New Team",
		EnableSelfAssignment: true,
		Workers:              []string{"worker_123"},
		Managers:             []string{"admin_123"},
		Hub:                  "hub_123",
	}

	team, err := client.Create(params)

	assert.NoError(t, err)
	assert.Equal(t, expectedTeam.ID, team.ID)

	mockClient.AssertRequestMade("POST", "/teams")
}

func TestClient_Update(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedTeam := testingutil.GetSampleTeam()
	expectedTeam.Name = "Updated Team Name"

	mockClient.AddResponse("/teams/team_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedTeam,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	params := onfleet.TeamUpdateParams{
		Name: "Updated Team Name",
	}

	team, err := client.Update("team_123", params)

	assert.NoError(t, err)
	assert.Equal(t, expectedTeam.ID, team.ID)
	assert.Equal(t, "Updated Team Name", team.Name)

	mockClient.AssertRequestMade("PUT", "/teams/team_123")
}

func TestClient_Delete(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/teams/team_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       map[string]interface{}{"success": true},
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	err := client.Delete("team_123")

	assert.NoError(t, err)
	mockClient.AssertRequestMade("DELETE", "/teams/team_123")
}

func TestClient_AutoDispatch(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := onfleet.TeamAutoDispatch{
		DispatchId: "dispatch_123",
	}

	mockClient.AddResponse("/teams/team_123/dispatch", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	params := onfleet.TeamAutoDispatchParams{
		MaxAllowedDelay:        300,   // 5 minutes
		MaxTasksPerRoute:       10,
		RouteEnd:              "hub_123",
		ScheduleTimeWindow:    []int64{28800, 64800}, // 8 AM to 6 PM
		ServiceTime:           600,   // 10 minutes
		TaskTimeWindow:        []int64{32400, 61200}, // 9 AM to 5 PM
	}

	response, err := client.AutoDispatch("team_123", &params)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.DispatchId, response.DispatchId)

	mockClient.AssertRequestMade("POST", "/teams/team_123/dispatch")
}

func TestClient_GetWorkerEta(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := onfleet.TeamWorkerEta{
		WorkerId: "worker_123",
		Vehicle:  onfleet.WorkerVehicleTypeCar,
		Steps: []onfleet.TeamWorkerEtaStep{
			{
				Location:       onfleet.DestinationLocation{-122.4194, 37.7749},
				TravelTime:     900,  // 15 minutes
				ServiceTime:    300,  // 5 minutes
				CompletionTime: 1640995200,
				Distance:       5.2,
			},
		},
	}

	mockClient.AddResponse("/teams/team_123/estimate", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	params := onfleet.TeamWorkerEtaQueryParams{
		DropoffLocation: "destination_456",
		PickupLocation:  "destination_789",
		PickupTime:      1640995200,
		ServiceTime:     300,
	}

	response, err := client.GetWorkerEta("team_123", params)

	assert.NoError(t, err)
	assert.Equal(t, "worker_123", response.WorkerId)
	assert.Len(t, response.Steps, 1)

	mockClient.AssertRequestMade("GET", "/teams/team_123/estimate")
}

func TestClient_ListTasks(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedResponse := onfleet.TeamTasks{
		Tasks: []onfleet.Task{
			testingutil.GetSampleTask(),
		},
		LastId: "last_task_789",
	}

	mockClient.AddResponse("/teams/team_123/tasks", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedResponse,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

	params := onfleet.TeamTasksListQueryParams{
		From:         1640995200,
		To:           1672531199,
		IsPickupTask: "false",
		LastId:       "prev_task_456",
	}

	response, err := client.ListTasks("team_123", &params)

	assert.NoError(t, err)
	assert.Len(t, response.Tasks, 1)
	assert.Equal(t, "last_task_789", response.LastId)

	mockClient.AssertRequestMade("GET", "/teams/team_123/tasks")
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
			name:       "get not found",
			method:     "GET",
			url:        "/teams/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Get("nonexistent")
				return err
			},
		},
		{
			name:       "create invalid",
			method:     "POST",
			url:        "/teams",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.TeamCreateParams{})
				return err
			},
		},
		{
			name:       "update not found",
			method:     "PUT",
			url:        "/teams/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Update("nonexistent", onfleet.TeamUpdateParams{})
				return err
			},
		},
		{
			name:       "delete forbidden",
			method:     "DELETE",
			url:        "/teams/team_123",
			statusCode: 403,
			operation: func(client *Client) error {
				return client.Delete("team_123")
			},
		},
		{
			name:       "auto dispatch error",
			method:     "POST",
			url:        "/teams/team_123/dispatch",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.AutoDispatch("team_123", &onfleet.TeamAutoDispatchParams{})
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

			client := Plug("test_api_key", nil, "https://api.example.com/teams", mockClient.MockCaller)

			err := tt.operation(client)
			assert.Error(t, err)
		})
	}
}