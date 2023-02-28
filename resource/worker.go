package resource

// Onfleet Worker.
// Reference https://docs.onfleet.com/reference/workers.
type Worker struct {
	AccountStatus                   WorkerAccountStatus        `json:"accountStatus,omitempty"`
	ActiveTask                      *string                    `json:"activeTask,omitempty"`
	AdditionalCapacities            WorkerAdditionalCapacities `json:"additionalCapacities,omitempty"`
	Capacity                        int                        `json:"capacity,omitempty"`
	DelayTime                       *int64                     `json:"delayTime,omitempty"`
	DisplayName                     *string                    `json:"displayName,omitempty"`
	HasRecentlyUsedSpoofedLocations bool                       `json:"hasRecentlyUsedSpoofedLocations,omitempty"`
	ID                              string                     `json:"id,omitempty"`
	ImageUrl                        *string                    `json:"imageUrl,omitempty"`
	Location                        *DestinationLocation       `json:"location,omitempty"`
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
	Timezone                        *string                    `json:"timezone,omitempty"`
	Vehicle                         WorkerVehicle              `json:"vehicle,omitempty"`
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

type WorkerAccountStatus string

const (
	WorkerAccountStatusAccepted WorkerAccountStatus = "ACCEPTED"
	WorkerAccountStatusInvited  WorkerAccountStatus = "INVITED"
)

// Onfleet Worker Schedule
// Reference https://docs.onfleet.com/reference/get-workers-schedule
type WorkerSchedule struct {
	Date     string     `json:"data,omitempty"`
	Shifts   [][2]int64 `json:"shifts,omitempty"`
	Timezone string     `json:"timezone,omitempty"`
}

type WorkerScheduleEntries struct {
	Entries []WorkerSchedule `json:"entries,omitempty"`
}
