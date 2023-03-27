package onfleet

// Onfleet Worker.
// Reference https://docs.onfleet.com/reference/workers.
type Worker struct {
	AccountStatus                   WorkerAccountStatus        `json:"accountStatus"`
	ActiveTask                      *string                    `json:"activeTask"`
	AdditionalCapacities            WorkerAdditionalCapacities `json:"additionalCapacities"`
	Addresses                       *WorkerAddresses           `json:"addresses,omitempty"`
	Analytics                       *WorkerAnalytics           `json:"analytics,omitempty"`
	Capacity                        float64                    `json:"capacity"`
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
	Vehicle                         *WorkerVehicle             `json:"vehicle"`
}

type WorkerUserData struct {
	AppVersion        string  `json:"appVersion,omitempty"`
	BatteryLevel      float32 `json:"batteryLevel,omitempty"`
	DeviceDescription string  `json:"deviceDescription,omitempty"`
	Platform          string  `json:"platform,omitempty"`
}

type WorkerAdditionalCapacities struct {
	CapacityA float64 `json:"capacityA"`
	CapacityB float64 `json:"capacityB"`
	CapacityC float64 `json:"capacityC"`
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
	Date     string    `json:"date"`
	Shifts   [][]int64 `json:"shifts"`
	Timezone string    `json:"timezone"`
}

type WorkerScheduleEntries struct {
	Entries []WorkerSchedule `json:"entries"`
}

type WorkerAddresses struct {
	Routing *WorkerAddressesRouting `json:"routing"`
}

type WorkerAddressesRouting struct {
	Address           DestinationAddress  `json:"address"`
	CreatedByLocation bool                `json:"createdByLocation"`
	GooglePlaceId     string              `json:"googlePlaceId"`
	ID                string              `json:"id"`
	Location          DestinationLocation `json:"location"`
	Notes             string              `json:"notes"`
	Organization      string              `json:"organization"`
	TimeCreated       int64               `json:"timeCreated"`
	TimeLastModified  int64               `json:"timeLastModified"`
	WasGeocoded       bool                `json:"wasGeocoded"`
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

type WorkerGetQueryParams struct {
	// Analytics indicates whether analytics data should be includes on the retrieved worker object
	Analytics bool `json:"analytics,omitempty"`
	// Filter is comma separated list of worker fields to return
	Filter string `json:"filter,omitempty"`
	From   int64  `json:"from,omitempty,string"`
	To     int64  `json:"to,omitempty,string"`
}

type WorkerListQueryParams struct {
	// Filter is comma separated list of worker fields to return
	Filter string `json:"filter,omitempty"`
	// Phones is a comma separated list of workers' phone numbers
	Phones string `json:"phones,omitempty"`
	// States is comma separeted list of worker states
	States string `json:"states,omitempty"`
	// Teams is a comma separated list of team ids worker must be part of
	Teams string `json:"teams,omitempty"`
}

type WorkersByLocation struct {
	Workers []Worker `json:"workers"`
}

type WorkersByLocationListQueryParams struct {
	Longitude float64 `json:"longitude,string"`
	Latitude  float64 `json:"latitude,string"`
	Radius    float64 `json:"radius,omitempty,string"`
}

type WorkerTasks struct {
	LastId string `json:"lastId,omitempty"`
	Tasks  []Task `json:"tasks"`
}

type WorkerTasksListQueryParams struct {
	From int64 `json:"from,omitempty,string"`
	// IsPickupTask is a boolean represented as a string.
	//
	// E.g. "true" or "false".
	//
	// Set to empty string "" if both dropoff and pickup tasks should be returned.
	IsPickupTask string `json:"isPickupTask,omitempty"`
	LastId       string `json:"lastId,omitempty"`
	To           int64  `json:"to,omitempty,string"`
}

type WorkerCreateParams struct {
	Addresses   *WorkerParamsAddressRouting `json:"addresses,omitempty"`
	Capacity    float64                     `json:"capacity,omitempty"`
	DisplayName string                      `json:"displayName,omitempty"`
	Metadata    []Metadata                  `json:"metadata,omitempty"`
	Name        string                      `json:"name"`
	Phone       string                      `json:"phone"`
	Teams       []string                    `json:"teams"`
	Vehicle     *WorkerParamsVehicle        `json:"vehicle,omitempty"`
}

type WorkerParamsAddressRouting struct {
	// Routing should be set to a destination id
	Routing string `json:"routing"`
}

type WorkerParamsVehicle struct {
	Color        string            `json:"color,omitempty"`
	Description  string            `json:"description,omitempty"`
	LicensePlate string            `json:"licensePlate,omitempty"`
	Type         WorkerVehicleType `json:"type,omitempty"`
}

type WorkerUpdateParams struct {
	Addresses   *WorkerParamsAddressRouting `json:"addresses,omitempty"`
	Capacity    float64                     `json:"capacity,omitempty"`
	DisplayName string                      `json:"displayName,omitempty"`
	Metadata    []Metadata                  `json:"metadata,omitempty"`
	Name        string                      `json:"name,omitempty"`
	Teams       []string                    `json:"teams,omitempty"`
	Vehicle     *WorkerParamsVehicle        `json:"vehicle,omitempty"`
}
