package webhook

import (
	"testing"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/testingutil"
)

func TestClient_List(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWebhooks := []onfleet.Webhook{
		testingutil.GetSampleWebhook(),
	}

	mockClient.AddResponse("/webhooks", testingutil.MockResponse{
		StatusCode: 200,
		Body:       expectedWebhooks,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/webhooks", mockClient.MockCaller)

	webhooks, err := client.List()

	testingutil.AssertNoError(t, err)
	testingutil.AssertLen(t, webhooks, 1)
	testingutil.AssertEqual(t, expectedWebhooks[0].ID, webhooks[0].ID)
	testingutil.AssertEqual(t, expectedWebhooks[0].Name, webhooks[0].Name)
	testingutil.AssertEqual(t, expectedWebhooks[0].Trigger, webhooks[0].Trigger)

	mockClient.AssertRequestMade("GET", "/webhooks")
	mockClient.AssertBasicAuth("test_api_key")
}

func TestClient_Create(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	expectedWebhook := testingutil.GetSampleWebhook()
	mockClient.AddResponse("/webhooks", testingutil.MockResponse{
		StatusCode: 201,
		Body:       expectedWebhook,
	})

	client := Plug("test_api_key", nil, "https://api.example.com/webhooks", mockClient.MockCaller)

	params := onfleet.WebhookCreateParams{
		Name:      "Task Completion Webhook",
		Trigger:   0, // Task completed
		Url:       "https://api.example.com/webhook/onfleet",
		Threshold: 5.0,
	}

	webhook, err := client.Create(params)

	testingutil.AssertNoError(t, err)
	testingutil.AssertEqual(t, expectedWebhook.ID, webhook.ID)
	testingutil.AssertEqual(t, expectedWebhook.Name, webhook.Name)
	testingutil.AssertEqual(t, expectedWebhook.Trigger, webhook.Trigger)

	mockClient.AssertRequestMade("POST", "/webhooks")
}

func TestClient_Delete(t *testing.T) {
	mockClient := testingutil.SetupTest(t)
	defer testingutil.CleanupTest(t, mockClient)

	mockClient.AddResponse("/webhooks/webhook_123", testingutil.MockResponse{
		StatusCode: 200,
		Body:       map[string]interface{}{"success": true},
	})

	client := Plug("test_api_key", nil, "https://api.example.com/webhooks", mockClient.MockCaller)

	err := client.Delete("webhook_123")

	testingutil.AssertNoError(t, err)
	mockClient.AssertRequestMade("DELETE", "/webhooks/webhook_123")
}

func TestClient_WebhookTriggerTypes(t *testing.T) {
	tests := []struct {
		name        string
		trigger     int
		description string
	}{
		{
			name:        "task completed",
			trigger:     0,
			description: "Task completed trigger",
		},
		{
			name:        "task failed",
			trigger:     1,
			description: "Task failed trigger",
		},
		{
			name:        "worker duty",
			trigger:     2,
			description: "Worker duty trigger",
		},
		{
			name:        "task started",
			trigger:     3,
			description: "Task started trigger",
		},
		{
			name:        "task eta",
			trigger:     4,
			description: "Task ETA trigger",
		},
		{
			name:        "task arrival",
			trigger:     5,
			description: "Task arrival trigger",
		},
		{
			name:        "task assignment",
			trigger:     6,
			description: "Task assignment trigger",
		},
		{
			name:        "task unassignment",
			trigger:     7,
			description: "Task unassignment trigger",
		},
		{
			name:        "task creation",
			trigger:     8,
			description: "Task creation trigger",
		},
		{
			name:        "task update",
			trigger:     9,
			description: "Task update trigger",
		},
		{
			name:        "task deletion",
			trigger:     10,
			description: "Task deletion trigger",
		},
		{
			name:        "sms recipient response missed",
			trigger:     12,
			description: "SMS recipient response missed trigger",
		},
		{
			name:        "sms recipient response replied",
			trigger:     13,
			description: "SMS recipient response replied trigger",
		},
		{
			name:        "auto dispatch completed",
			trigger:     14,
			description: "Auto dispatch completed trigger",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := testingutil.SetupTest(t)
			defer testingutil.CleanupTest(t, mockClient)

			expectedWebhook := testingutil.GetSampleWebhook()
			expectedWebhook.Trigger = tt.trigger

			mockClient.AddResponse("/webhooks", testingutil.MockResponse{
				StatusCode: 201,
				Body:       expectedWebhook,
			})

			client := Plug("test_api_key", nil, "https://api.example.com/webhooks", mockClient.MockCaller)

			params := onfleet.WebhookCreateParams{
				Name:    tt.description,
				Trigger: tt.trigger,
				Url:     "https://api.example.com/webhook/onfleet",
			}

			webhook, err := client.Create(params)

			testingutil.AssertNoError(t, err)
			testingutil.AssertEqual(t, tt.trigger, webhook.Trigger)
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
			name:       "create invalid url",
			method:     "POST",
			url:        "/webhooks",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.WebhookCreateParams{
					Name:    "Invalid Webhook",
					Trigger: 0,
					Url:     "invalid-url",
				})
				return err
			},
		},
		{
			name:       "create invalid trigger",
			method:     "POST",
			url:        "/webhooks",
			statusCode: 400,
			operation: func(client *Client) error {
				_, err := client.Create(onfleet.WebhookCreateParams{
					Name:    "Invalid Webhook",
					Trigger: 999,
					Url:     "https://api.example.com/webhook",
				})
				return err
			},
		},
		{
			name:       "delete not found",
			method:     "DELETE",
			url:        "/webhooks/nonexistent",
			statusCode: 404,
			operation: func(client *Client) error {
				return client.Delete("nonexistent")
			},
		},
		{
			name:       "list unauthorized",
			method:     "GET",
			url:        "/webhooks",
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

			client := Plug("test_api_key", nil, "https://api.example.com/webhooks", mockClient.MockCaller)

			err := tt.operation(client)
			testingutil.AssertError(t, err)
		})
	}
}