package container

import (
	"testing"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_GetWorkerContainer(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedContainer := testingutil.GetSampleContainer()
	mockClient.AddResponse("/containers/workers/worker_456", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedContainer,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

	container, err := client.Get("worker_456", onfleet.ContainerQueryKeyWorkers)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedContainer.ID, container.ID)
	testingutil.AssertEqual(t, expectedContainer.Type, container.Type)
	testingutil.AssertEqual(t, expectedContainer.Worker, container.Worker)
	testingutil.AssertLen(t, container.Tasks, 3)

	mockClient.AssertRequestMade("GET", "/containers/workers/worker_456")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_GetTeamContainer(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedContainer := testingutil.GetSampleContainer()
	expectedContainer.Type = onfleet.ContainerTypeTeam
	expectedContainer.Team = "team_123"
	expectedContainer.Worker = ""

	mockClient.AddResponse("/containers/teams/team_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedContainer,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

	container, err := client.Get("team_123", onfleet.ContainerQueryKeyTeams)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedContainer.ID, container.ID)
	testingutil.AssertEqual(t, onfleet.ContainerTypeTeam, container.Type)
	testingutil.AssertEqual(t, "team_123", container.Team)

	mockClient.AssertRequestMade("GET", "/containers/teams/team_123")
}

func TestClient_GetOrganizationContainer(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedContainer := testingutil.GetSampleContainer()
	expectedContainer.Type = onfleet.ContainerTypeOrganization
	expectedContainer.Team = ""
	expectedContainer.Worker = ""

	mockClient.AddResponse("/containers/organizations/org_789", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedContainer,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

	container, err := client.Get("org_789", onfleet.ContainerQueryKeyOrganizations)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedContainer.ID, container.ID)
	testingutil.AssertEqual(t, onfleet.ContainerTypeOrganization, container.Type)
	testingutil.AssertEqual(t, expectedContainer.Organization, container.Organization)

	mockClient.AssertRequestMade("GET", "/containers/organizations/org_789")
}

func TestClient_InsertTasksAtIndex(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedContainer := testingutil.GetSampleContainer()
	expectedContainer.Tasks = []string{"task_111", "task_444", "task_555", "task_222", "task_333"}

	mockClient.AddResponse("/containers/workers/worker_456", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedContainer,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

	params := onfleet.ContainerTaskInsertParams{
		Tasks: []any{
			map[string]interface{}{
				"id":    "task_444",
				"index": 1,
			},
			map[string]interface{}{
				"id":    "task_555",
				"index": 2,
			},
		},
		ConsiderDependencies: false,
	}

	container, err := client.InsertTasks("worker_456", onfleet.ContainerQueryKeyWorkers, params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedContainer.ID, container.ID)
	testingutil.AssertLen(t, container.Tasks, 5)
	testingutil.AssertEqual(t, "task_444", container.Tasks[1])
	testingutil.AssertEqual(t, "task_555", container.Tasks[2])

	mockClient.AssertRequestMade("PUT", "/containers/workers/worker_456")
}

func TestClient_InsertTasksWithDependencies(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedContainer := testingutil.GetSampleContainer()
	expectedContainer.Tasks = []string{"task_111", "task_222", "task_333", "task_666"}

	mockClient.AddResponse("/containers/workers/worker_456", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedContainer,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

	params := onfleet.ContainerTaskInsertParams{
		Tasks: []any{
			map[string]interface{}{
				"id": "task_666",
			},
		},
		ConsiderDependencies: true,
	}

	container, err := client.InsertTasks("worker_456", onfleet.ContainerQueryKeyWorkers, params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedContainer.ID, container.ID)
	testingutil.AssertLen(t, container.Tasks, 4)
	testingutil.AssertEqual(t, "task_666", container.Tasks[3])

	mockClient.AssertRequestMade("PUT", "/containers/workers/worker_456")
}

func TestClient_ContainerTypes(t *testing.T) {
	tests := []struct {
		name          string
		containerType onfleet.ContainerType
		queryKey      onfleet.ContainerQueryKey
		entityID      string
		workerField   string
		teamField     string
	}{
		{
			name:          "worker container",
			containerType: onfleet.ContainerTypeWorker,
			queryKey:      onfleet.ContainerQueryKeyWorkers,
			entityID:      "worker_123",
			workerField:   "worker_123",
			teamField:     "",
		},
		{
			name:          "team container",
			containerType: onfleet.ContainerTypeTeam,
			queryKey:      onfleet.ContainerQueryKeyTeams,
			entityID:      "team_456",
			workerField:   "",
			teamField:     "team_456",
		},
		{
			name:          "organization container",
			containerType: onfleet.ContainerTypeOrganization,
			queryKey:      onfleet.ContainerQueryKeyOrganizations,
			entityID:      "org_789",
			workerField:   "",
			teamField:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedContainer := testingutil.GetSampleContainer()
			expectedContainer.Type = tt.containerType
			expectedContainer.Worker = tt.workerField
			expectedContainer.Team = tt.teamField

			expectedURL := "/containers/" + string(tt.queryKey) + "/" + tt.entityID
			mockClient.AddResponse(expectedURL, testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedContainer,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

			container, err := client.Get(tt.entityID, tt.queryKey)

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.containerType, container.Type)
			testingutil.AssertEqual(t, tt.workerField, container.Worker)
			testingutil.AssertEqual(t, tt.teamField, container.Team)

			mockClient.AssertRequestMade("GET", expectedURL)
		})
	}
}

func TestClient_TaskOperations(t *testing.T) {
	tests := []struct {
		name                 string
		params               onfleet.ContainerTaskInsertParams
		expectedTasksLength  int
		considerDependencies bool
	}{
		{
			name: "append tasks without index",
			params: onfleet.ContainerTaskInsertParams{
				Tasks: []any{
					map[string]interface{}{"id": "task_new1"},
					map[string]interface{}{"id": "task_new2"},
				},
				ConsiderDependencies: false,
			},
			expectedTasksLength:  5,
			considerDependencies: false,
		},
		{
			name: "insert tasks at specific indices",
			params: onfleet.ContainerTaskInsertParams{
				Tasks: []any{
					map[string]interface{}{
						"id":    "task_new3",
						"index": 0,
					},
					map[string]interface{}{
						"id":    "task_new4",
						"index": 2,
					},
				},
				ConsiderDependencies: false,
			},
			expectedTasksLength:  5,
			considerDependencies: false,
		},
		{
			name: "append with dependency consideration",
			params: onfleet.ContainerTaskInsertParams{
				Tasks: []any{
					map[string]interface{}{"id": "task_dependent"},
				},
				ConsiderDependencies: true,
			},
			expectedTasksLength:  4,
			considerDependencies: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedContainer := testingutil.GetSampleContainer()
			// Adjust tasks based on operation
			switch tt.name {
			case "append tasks without index":
				expectedContainer.Tasks = []string{"task_111", "task_222", "task_333", "task_new1", "task_new2"}
			case "insert tasks at specific indices":
				expectedContainer.Tasks = []string{"task_new3", "task_111", "task_new4", "task_222", "task_333"}
			case "append with dependency consideration":
				expectedContainer.Tasks = []string{"task_111", "task_222", "task_333", "task_dependent"}
			}

			mockClient.AddResponse("/containers/workers/worker_456", testingutil.MockResponse{
				StatusCode: 200,
				Body:       expectedContainer,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

			container, err := client.InsertTasks("worker_456", onfleet.ContainerQueryKeyWorkers, tt.params)

			testingutil.AssertNoError(t, err)
			testingutil.AssertLen(t, container.Tasks, tt.expectedTasksLength)

			mockClient.AssertRequestMade("PUT", "/containers/workers/worker_456")
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
			name:       "get worker container not found",
			method:     "GET",
			url:        "/containers/workers/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Get("nonexistent", onfleet.ContainerQueryKeyWorkers)
				return err
			},
		},
		{
			name:       "get team container not found",
			method:     "GET",
			url:        "/containers/teams/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				_, err := client.Get("nonexistent", onfleet.ContainerQueryKeyTeams)
				return err
			},
		},
		{
			name:       "insert tasks invalid container",
			method:     "PUT",
			url:        "/containers/workers/inactive_worker",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.InsertTasks("inactive_worker", onfleet.ContainerQueryKeyWorkers, onfleet.ContainerTaskInsertParams{
					Tasks: []any{
						map[string]interface{}{"id": "task_123"},
					},
				})
				return err
			},
		},
		{
			name:       "insert tasks invalid task id",
			method:     "PUT",
			url:        "/containers/workers/worker_456",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.InsertTasks("worker_456", onfleet.ContainerQueryKeyWorkers, onfleet.ContainerTaskInsertParams{
					Tasks: []any{
						map[string]interface{}{"id": "invalid_task"},
					},
				})
				return err
			},
		},
		{
			name:       "unauthorized access",
			method:     "GET",
			url:        "/containers/workers/worker_456",
			statusCode: 401,
			operation: func(client *Client) error {
				_, err := client.Get("worker_456", onfleet.ContainerQueryKeyWorkers)
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

			client := Plug("test_api_key", nil, "https://api.example.com/containers", mockClient.MockCaller)

			err := tt.operation(client)
			testingutil.AssertError(t, err)
		})
	}
}