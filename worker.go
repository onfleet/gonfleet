package onfleet

// Onfleet Worker.
// Reference https://docs.onfleet.com/reference/workers.
type Worker struct {
	AccountStatus                   WorkerAccountStatus        `json:"accountStatus"`
	ActiveTask                      *string                    `json:"activeTask"`
	AdditionalCapacities            WorkerAdditionalCapacities `json:"additionalCapacities"`
	Analytics                       *WorkerAnalytics           `json:"analytics,omitempty"`
	Capacity                        *int                       `json:"capacity"`
	DelayTime                       *float64                   `json:"delayTime"`
	DisplayName                     *string                    `json:"displayName"`
	HasRecentlyUsedSpoofedLocations bool                       `json:"hasRecentlyUsedSpoofedLocations"`
	ID                              string                     `json:"id"`
	ImageUrl                        *string                    `json:"imageUrl"`
	Location                        DestinationLocation        `json:"location"`
	Metadata                        []Metadata                 `json:"metadata"`
	Name                            string                     `json:"name"`
	OnDuty                          bool                       `json:"onDuty"`
	Organization                    string                     `json:"organization"`
	Phone                           string                     `json:"phone"`
	Tasks                           []string                   `json:"tasks"`
	Teams                           []string                   `json:"teams"`
	TimeCreated                     int64                      `json:"timeCreated"`
	TimeLastModified                int64                      `json:"timeLastModified"`
	TimeLastSeen                    int64                      `json:"timeLastSeen"`
	UserData                        WorkerUserData             `json:"userData"`
	Timezone                        *string                    `json:"timezone"`
	Vehicle                         WorkerVehicle              `json:"vehicle"`
}

type WorkerUserData struct {
	AppVersion        string  `json:"appVersion,omitempty"`
	BatteryLevel      float32 `json:"batteryLevel,omitempty"`
	DeviceDescription string  `json:"deviceDescription,omitempty"`
	Platform          string  `json:"platform,omitempty"`
}

type WorkerAdditionalCapacities struct {
	CapacityA int `json:"capacityA"`
	CapacityB int `json:"capacityB"`
	CapacityC int `json:"capacityC"`
}

type WorkerVehicleType string

const (
	WorkerVehicleTypeCar        WorkerVehicleType = "CAR"
	WorkerVehicleTypeBicycle    WorkerVehicleType = "BICYCLE"
	WorkerVehicleTypeMotorcycle WorkerVehicleType = "MOTORCYCLE"
	WorkerVehicleTypeTruck      WorkerVehicleType = "TRUCK"
)

type WorkerVehicle struct {
	Color            *string           `json:"color"`
	Description      *string           `json:"description"`
	ID               string            `json:"id"`
	LicensePlate     *string           `json:"licensePlate"`
	TimeLastModified int64             `json:"timeLastModified"`
	Type             WorkerVehicleType `json:"type"`
}

type WorkerAccountStatus string

const (
	WorkerAccountStatusAccepted WorkerAccountStatus = "ACCEPTED"
	WorkerAccountStatusInvited  WorkerAccountStatus = "INVITED"
)

// Onfleet Worker Schedule
// Reference https://docs.onfleet.com/reference/get-workers-schedule
type WorkerSchedule struct {
	Date     string     `json:"data"`
	Shifts   [][2]int64 `json:"shifts"`
	Timezone string     `json:"timezone"`
}

type WorkerScheduleEntries struct {
	Entries []WorkerSchedule `json:"entries"`
}

type WorkerAnalytics struct {
	Distances  WorkerAnalyticsDistances  `json:"distances"`
	Events     []WorkerAnalyticsEvent    `json:"events"`
	TaskCounts WorkerAnalyticsTaskCounts `json:"taskCounts"`
	Times      WorkerAnalyticsTimes      `json:"times"`
}

type WorkerAnalyticsEvent struct {
	Action string `json:"action"`
	Time   int64  `json:"time"`
}

type WorkerAnalyticsDistances struct {
	Enroute float64 `json:"enroute"`
	Idle    float64 `json:"idle"`
}

type WorkerAnalyticsTimes struct {
	Enroute float64 `json:"enroute"`
	Idle    float64 `json:"idle"`
}

type WorkerAnalyticsTaskCounts struct {
	Failed    int `json:"failed"`
	Succeeded int `json:"succeeded"`
}

type WorkerGetParams struct {
	// Analytics indicates whether analytics data should be includes on the retrieved worker object
	Analytics bool `json:"analytics,omitempty"`
	From   int64  `json:"from,omitempty,string"`
	To     int64  `json:"to,omitempty,string"`
}
