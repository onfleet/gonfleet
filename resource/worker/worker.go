package worker

import (
	"net/http"

	"github.com/onfleet/gonfleet/resource/destination"
	"github.com/onfleet/gonfleet/resource/metadata"
)

// Client for Workers resource
type Client struct {
	ApiKey     string
	HttpClient *http.Client
	Url        string
}

type WorkerUserData struct {
	AppVersion        string `json:"appVersion,omitempty"`
	BatteryLevel      int    `json:"batteryLevel,omitempty"`
	DeviceDescription string `json:"deviceDescription,omitempty"`
	Platform          string `json:"platform,omitempty"`
}

type WorkerAdditionalCapacities struct {
	CapacityA int `json:"capacityA,omitempty"`
	CapacityB int `json:"capacityB,omitempty"`
	CapacityC int `json:"capacityC,omitempty"`
}

type WorkerVehicleType string

const (
	WorkerVehicleTypeCar        WorkerVehicleType = "CAR"
	WorkerVehicleTypeTruck      WorkerVehicleType = "TRUCK"
	WorkerVehicleTypeBicycle    WorkerVehicleType = "BICYCLE"
	WorkerVehicleTypeMotorcycle WorkerVehicleType = "MOTORCYCLE"
)

type WorkerVehicle struct {
	ID               string            `json:"id,omitempty"`
	Type             WorkerVehicleType `json:"type,omitempty"`
	Description      *string           `json:"description,omitempty"`
	LicensePlate     *string           `json:"licensePlate,omitempty"`
	Color            *string           `json:"color,omitempty"`
	TimeLastModified int64             `json:"timeLastModified,omitempty"`
}

type AccountStatusOption string

const (
	AccountStatusAccepted AccountStatusOption = "ACCEPTED"
	AccountStatusInvited  AccountStatusOption = "INVITED"
)

// Worker refers to an Onfleet Worker.
// Reference https://docs.onfleet.com/reference/workers.
type Worker struct {
	ID                              string                     `json:"id,omitempty"`
	TimeCreated                     int64                      `json:"timeCreated,omitempty"`
	TimeLastModified                int64                      `json:"timeLastModified,omitempty"`
	Organization                    string                     `json:"organization,omitempty"`
	Name                            string                     `json:"name,omitempty"`
	DisplayName                     *string                    `json:"displayName,omitempty"`
	Phone                           string                     `json:"phone,omitempty"`
	ActiveTask                      *string                    `json:"activeTask,omitempty"`
	Tasks                           []string                   `json:"tasks,omitempty"`
	OnDuty                          bool                       `json:"onDuty,omitempty"`
	TimeLastSeen                    int64                      `json:"timeLastSeen,omitempty"`
	Capacity                        int                        `json:"capacity,omitempty"`
	AdditionalCapacities            WorkerAdditionalCapacities `json:"additionalCapacities,omitempty"`
	UserData                        WorkerUserData             `json:"userData,omitempty"`
	AccountStatus                   AccountStatusOption        `json:"accountStatus,omitempty"`
	Metadata                        []metadata.Metadata        `json:"metadata,omitempty"`
	TimeZone                        *string                    `json:"timezone,omitempty"`
	Teams                           []string                   `json:"teams,omitempty"`
	ImageUrl                        *string                    `json:"imageUrl,omitempty"`
	DelayTime                       *int64                     `json:"delayTime,omitempty"`
	Location                        *destination.Location      `json:"location,omitempty"`
	Vehicle                         WorkerVehicle              `json:"vehicle,omitempty"`
	HasRecentlyUsedSpoofedLocations bool                       `json:"hasRecentlyUsedSpoofedLocations,omitempty"`
}
