package types

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
	WorkerVehicleTypeBicycle    WorkerVehicleType = "BICYCLE"
	WorkerVehicleTypeMotorcycle WorkerVehicleType = "MOTORCYCLE"
	WorkerVehicleTypeTruck      WorkerVehicleType = "TRUCK"
)

type WorkerVehicle struct {
	Color            *string           `json:"color,omitempty"`
	Description      *string           `json:"description,omitempty"`
	ID               string            `json:"id,omitempty"`
	LicensePlate     *string           `json:"licensePlate,omitempty"`
	TimeLastModified int64             `json:"timeLastModified,omitempty"`
	Type             WorkerVehicleType `json:"type,omitempty"`
}

type AccountStatusOption string

const (
	AccountStatusAccepted AccountStatusOption = "ACCEPTED"
	AccountStatusInvited  AccountStatusOption = "INVITED"
)

// Onfleet Worker.
// Reference https://docs.onfleet.com/reference/workers.
type Worker struct {
	AccountStatus                   AccountStatusOption        `json:"accountStatus,omitempty"`
	ActiveTask                      *string                    `json:"activeTask,omitempty"`
	AdditionalCapacities            WorkerAdditionalCapacities `json:"additionalCapacities,omitempty"`
	Capacity                        int                        `json:"capacity,omitempty"`
	DelayTime                       *int64                     `json:"delayTime,omitempty"`
	DisplayName                     *string                    `json:"displayName,omitempty"`
	HasRecentlyUsedSpoofedLocations bool                       `json:"hasRecentlyUsedSpoofedLocations,omitempty"`
	ID                              string                     `json:"id,omitempty"`
	ImageUrl                        *string                    `json:"imageUrl,omitempty"`
	Location                        *Location                  `json:"location,omitempty"`
	Metadata                        []Metadata                 `json:"metadata,omitempty"`
	Name                            string                     `json:"name,omitempty"`
	OnDuty                          bool                       `json:"onDuty,omitempty"`
	Organization                    string                     `json:"organization,omitempty"`
	Phone                           string                     `json:"phone,omitempty"`
	Tasks                           []string                   `json:"tasks,omitempty"`
	Teams                           []string                   `json:"teams,omitempty"`
	TimeCreated                     int64                      `json:"timeCreated,omitempty"`
	TimeLastModified                int64                      `json:"timeLastModified,omitempty"`
	TimeLastSeen                    int64                      `json:"timeLastSeen,omitempty"`
	UserData                        WorkerUserData             `json:"userData,omitempty"`
	TimeZone                        *string                    `json:"timezone,omitempty"`
	Vehicle                         WorkerVehicle              `json:"vehicle,omitempty"`
}
