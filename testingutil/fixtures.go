package testingutil

import (
	"time"

	"github.com/onfleet/gonfleet"
)

// API Response Fixtures

// GetSampleWebhook returns a sample Webhook struct for testing
func GetSampleWebhook() onfleet.Webhook {
	return onfleet.Webhook{
		ID:        "webhook_123",
		Name:      "Task Completion Webhook",
		Trigger:   0, // Task completed
		Url:       "https://api.example.com/webhook/onfleet",
		IsEnabled: true,
		Count:     42,
		Threshold: 5.0,
	}
}

// GetSampleHub returns a sample Hub struct for testing
func GetSampleHub() onfleet.Hub {
	return onfleet.Hub{
		ID:   "hub_123",
		Name: "Main Distribution Center",
		Address: onfleet.DestinationAddress{
			Street:     "123 Main Street",
			City:       "San Francisco",
			State:      "CA",
			PostalCode: "94105",
			Country:    "US",
		},
		Location: onfleet.DestinationLocation{-122.4194, 37.7749},
		Teams:    []string{"team_123", "team_456"},
	}
}

// GetSampleAdmin returns a sample Admin struct for testing
func GetSampleAdmin() onfleet.Admin {
	return onfleet.Admin{
		ID:               "admin_123",
		Email:            "admin@example.com",
		Name:             "John Admin",
		Phone:            "+15551234567",
		Type:             "standard",
		IsAccountOwner:   false,
		IsActive:         true,
		IsReadOnly:       false,
		Organization:     "org_456",
		Teams:            []string{"team_123"},
		TimeCreated:      1640995200,
		TimeLastModified: 1640995300,
		Metadata: []onfleet.Metadata{
			{
				Name:  "department",
				Type:  "string",
				Value: "operations",
			},
		},
	}
}

// GetSampleContainer returns a sample Container struct for testing
func GetSampleContainer() onfleet.Container {
	return onfleet.Container{
		ID:               "container_123",
		Type:             onfleet.ContainerTypeWorker,
		Worker:           "worker_456",
		Organization:     "org_789",
		ActiveTask:       GetStringPtr("task_111"),
		Tasks:            []string{"task_111", "task_222", "task_333"},
		TimeCreated:      1640995200,
		TimeLastModified: 1640995300,
	}
}

// GetSampleRoutePlan returns a sample RoutePlan struct for testing
func GetSampleRoutePlan() onfleet.RoutePlan {
	return onfleet.RoutePlan{
		Id:               "routeplan_123",
		Name:             "Morning Delivery Route",
		State:            "active",
		Color:            "#FF5733",
		Tasks:            []string{"task_111", "task_222", "task_333"},
		Organization:     "org_456",
		Team:             GetStringPtr("team_789"),
		Worker:           "worker_123",
		VehicleType:      "CAR",
		StartTime:        1640995200,
		EndTime:          GetInt64Ptr(1641038400),
		ActualStartTime:  GetInt64Ptr(1640995300),
		ActualEndTime:    nil,
		StartingHubId:    GetStringPtr("hub_456"),
		EndingHubId:      GetStringPtr("hub_789"),
		ShortId:          "MR123",
		TimeCreated:      1640990000,
		TimeLastModified: 1640995200,
	}
}


// GetSampleDeliveryManifest returns a sample DeliveryManifest struct for testing
func GetSampleDeliveryManifest() onfleet.DeliveryManifest {
	return onfleet.DeliveryManifest{
		DepartureTime: 1640995200,
		Driver: onfleet.Driver{
			Name:  "John Doe",
			Phone: "+15551234567",
		},
		HubAddress:    "123 Main Street, San Francisco, CA 94105",
		ManifestDate:  1640995200,
		Tasks:         []onfleet.Task{GetSampleTask()},
		TotalDistance: "25.3 miles",
		TurnByTurn: []onfleet.TurnByTurn{
			{
				DrivingDistance: "2.5 miles",
				EndAddress:      "456 Oak Avenue, San Francisco, CA 94103",
				ETA:             1640995800,
				StartAddress:    "123 Main Street, San Francisco, CA 94105",
				Steps:           []string{"Head north on Main St", "Turn right on Oak Ave", "Destination on right"},
			},
		},
		Vehicle: onfleet.WorkerVehicleParam{
			Type:        onfleet.WorkerVehicleTypeCar,
			Description: "Blue Honda Civic",
			LicensePlate: "ABC123",
		},
	}
}

// GetSampleTask returns a sample Task struct for testing
func GetSampleTask() onfleet.Task {
	return onfleet.Task{
		ID:        "task_123",
		ShortId:   "abc123",
		Creator:   "user_456",
		Executor:  "worker_789",
		Worker:    GetStringPtr("worker_789"),
		Merchant:  "merchant_123",
		Organization: "org_456",
		State:     onfleet.TaskStateAssigned,
		PickupTask: false,
		Quantity:   1.0,
		ServiceTime: 5.0,
		TimeCreated: time.Now().Unix() - 3600,
		TimeLastModified: time.Now().Unix() - 1800,
		TrackingUrl: "https://onfleet.com/track/abc123",
		TrackingViewed: false,
		Notes: "Test task notes",
		Destination: GetSampleDestination(),
		Recipients: []onfleet.Recipient{GetSampleRecipient()},
		CustomFields: []onfleet.CustomField{},
		Metadata: []onfleet.Metadata{},
		Dependencies: []string{},
		Feedback: []any{},
		AdditionalQuantities: onfleet.TaskAdditionalQuantities{
			QuantityA: 0.0,
			QuantityB: 0.0,
			QuantityC: 0.0,
		},
		Appearance: onfleet.TaskAppearance{
			TriangleColor: GetIntPtr(0),
		},
		Identity: onfleet.TaskIdentity{
			Checksum: nil,
			FailedScanCount: 0,
		},
		CompletionDetails: onfleet.TaskCompletionDetails{
			Success: false,
			Actions: []any{},
			Events: []onfleet.TaskCompletionEvent{},
			Distance: 0.0,
			FailureNotes: "",
			FailureReason: "",
			Notes: "",
			PhotoUploadId: nil,
			PhotoUploadIds: nil,
			SignatureUploadId: nil,
			Time: nil,
			UnavailableAttachments: []any{},
			FirstLocation: onfleet.DestinationLocation{},
			LastLocation: onfleet.DestinationLocation{},
		},
		Overrides: onfleet.TaskOverrides{
			RecipientName: nil,
			RecipientNotes: nil,
			RecipientSkipSmsNotifications: nil,
			UseMerchantForProxy: nil,
		},
		Container: &onfleet.TaskContainer{
			Type: onfleet.ContainerTypeWorker,
			Worker: "worker_789",
		},
		Barcodes: nil,
		CompleteAfter: nil,
		CompleteBefore: nil,
		EstimatedArrivalTime: nil,
		EstimatedCompletionTime: nil,
		ETA: nil,
		DelayTime: nil,
		ScanOnlyRequiredBarcodes: false,
	}
}

// GetSampleWorker returns a sample Worker struct for testing
func GetSampleWorker() onfleet.Worker {
	return onfleet.Worker{
		ID:          "worker_123",
		Name:        "John Doe",
		DisplayName: GetStringPtr("Johnny"),
		Phone:       "+15551234567",
		Capacity:    10.0,
		OnDuty:      true,
		Organization: "org_456",
		AccountStatus: onfleet.WorkerAccountStatusAccepted,
		TimeCreated: time.Now().Unix() - 86400,
		TimeLastModified: time.Now().Unix() - 1800,
		TimeLastSeen: time.Now().Unix() - 300,
		Location: onfleet.DestinationLocation{-122.4194, 37.7749}, // [longitude, latitude]
		Teams: []string{"team_123", "team_456"},
		Tasks: []string{"task_123", "task_456"},
		ActiveTask: GetStringPtr("task_123"),
		Metadata: []onfleet.Metadata{},
		AdditionalCapacities: onfleet.WorkerAdditionalCapacities{
			CapacityA: 5.0,
			CapacityB: 3.0,
			CapacityC: 2.0,
		},
		Vehicle: &onfleet.WorkerVehicle{
			ID:    "vehicle_123",
			Type:  onfleet.WorkerVehicleTypeCar,
			Color: GetStringPtr("Blue"),
			Description: GetStringPtr("2020 Honda Civic"),
			LicensePlate: GetStringPtr("ABC123"),
			TimeLastModified: time.Now().Unix() - 86400,
		},
		ImageUrl: GetStringPtr("https://example.com/avatar.jpg"),
		Timezone: GetStringPtr("America/Los_Angeles"),
		DelayTime: nil,
		HasRecentlyUsedSpoofedLocations: false,
		UserData: onfleet.WorkerUserData{
			AppVersion: "1.2.3",
			BatteryLevel: 85.5,
			DeviceDescription: "iPhone 12",
			Platform: "iOS",
		},
		Addresses: &onfleet.WorkerAddresses{
			Routing: &onfleet.WorkerAddressesRouting{
				ID: "address_123",
				Address: onfleet.DestinationAddress{
					Number: "123",
					Street: "Main St",
					City: "San Francisco",
					State: "CA",
					PostalCode: "94105",
					Country: "US",
				},
				Location: onfleet.DestinationLocation{-122.4194, 37.7749}, // [longitude, latitude]
				GooglePlaceId: "place_123",
				Notes: "Home address",
				Organization: "org_456",
				TimeCreated: time.Now().Unix() - 86400,
				TimeLastModified: time.Now().Unix() - 86400,
				CreatedByLocation: false,
				WasGeocoded: true,
			},
		},
		Analytics: nil,
	}
}

// GetSampleDestination returns a sample Destination struct for testing
func GetSampleDestination() onfleet.Destination {
	return onfleet.Destination{
		ID: "destination_123",
		Address: onfleet.DestinationAddress{
			Number:     "123",
			Street:     "Main St",
			City:       "San Francisco",
			State:      "CA",
			PostalCode: "94105",
			Country:    "US",
		},
		Location: onfleet.DestinationLocation{-122.4194, 37.7749}, // [longitude, latitude]
		GooglePlaceId: "place_123",
		Notes: "Front door delivery",
		TimeCreated: time.Now().Unix() - 3600,
		TimeLastModified: time.Now().Unix() - 1800,
		Metadata: []onfleet.Metadata{},
		Warnings: []any{},
	}
}

// GetSampleRecipient returns a sample Recipient struct for testing
func GetSampleRecipient() onfleet.Recipient {
	return onfleet.Recipient{
		ID:          "recipient_123",
		Name:        "Jane Smith",
		Phone:       "+15559876543",
		Organization: "org_456",
		SkipSmsNotifications: false,
		TimeCreated: time.Now().Unix() - 3600,
		TimeLastModified: time.Now().Unix() - 1800,
		Metadata: []onfleet.Metadata{},
		Notes: "Customer prefers text updates",
	}
}

// GetSampleTeam returns a sample Team struct for testing
func GetSampleTeam() onfleet.Team {
	return onfleet.Team{
		ID:     "team_123",
		Name:   "Delivery Team A",
		TimeCreated: time.Now().Unix() - 86400,
		TimeLastModified: time.Now().Unix() - 3600,
		Workers: []string{"worker_123", "worker_456"},
		Tasks:   []string{"task_123", "task_456"},
		Managers: []string{"admin_123"},
		Hub:     GetStringPtr("hub_123"),
		EnableSelfAssignment: true,
	}
}

// GetSampleOrganization returns a sample Organization struct for testing
func GetSampleOrganization() onfleet.Organization {
	return onfleet.Organization{
		ID:     "org_123",
		Name:   "Test Organization",
		Email:  "admin@testorg.com",
		DriverSupportEmail: "support@testorg.com",
		Timezone: "America/Los_Angeles",
		Country: "US",
		TimeCreated: time.Now().Unix() - 86400*30,
		TimeLastModified: time.Now().Unix() - 86400,
		Image: "https://example.com/logo.png",
		Delegatees: []string{},
	}
}



// Parameter Fixtures for Create/Update operations

// GetSampleTaskParams returns sample parameters for creating a task
func GetSampleTaskParams() onfleet.TaskParams {
	return onfleet.TaskParams{
		Destination: onfleet.DestinationCreateParams{
			Address: onfleet.DestinationAddress{
				Number:     "789",
				Street:     "Test Ave",
				City:       "San Francisco",
				State:      "CA",
				PostalCode: "94108",
				Country:    "US",
			},
			Notes: "Side entrance",
		},
		Recipients: []onfleet.RecipientCreateParams{
			{
				Name:  "Bob Johnson",
				Phone: "+15551112222",
				Notes: "Call upon arrival",
			},
		},
		PickupTask: false,
		Quantity:   2.0,
		ServiceTime: 10.0,
		Notes: "Handle with care",
		CompleteAfter: time.Now().Unix() + 3600,
		CompleteBefore: time.Now().Unix() + 7200,
	}
}

// GetSampleWorkerCreateParams returns sample parameters for creating a worker
func GetSampleWorkerCreateParams() onfleet.WorkerCreateParams {
	return onfleet.WorkerCreateParams{
		Name:     "Alice Cooper",
		Phone:    "+15553334444",
		Capacity: 15.0,
		Teams:    []string{"team_123"},
		DisplayName: "Alice",
		Vehicle: &onfleet.WorkerVehicleParam{
			Type:         onfleet.WorkerVehicleTypeBicycle,
			Color:        "Red",
			Description:  "Mountain bike",
			LicensePlate: "",
		},
		Addresses: &onfleet.WorkerAddressRoutingParam{
			Routing: "destination_456",
		},
		Metadata: []onfleet.Metadata{
			{
				Name:  "employee_id",
				Type:  "string",
				Value: "EMP001",
			},
		},
	}
}

// Error Response Fixtures

// GetSampleErrorResponse returns a sample error response for testing
func GetSampleErrorResponse() map[string]interface{} {
	return map[string]interface{}{
		"message": map[string]interface{}{
			"error":   2000,
			"message": "Invalid API key",
			"cause":   "The API key provided is not valid",
			"request": "12345-abcde-67890",
		},
	}
}

// GetSampleValidationErrorResponse returns a sample validation error response
func GetSampleValidationErrorResponse() map[string]interface{} {
	return map[string]interface{}{
		"message": map[string]interface{}{
			"error":   2001,
			"message": "Invalid parameters",
			"cause":   "The 'phone' parameter is required",
			"request": "12345-abcde-67891",
		},
	}
}