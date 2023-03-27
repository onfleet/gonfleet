package onfleet

// Reference https://docs.onfleet.com/reference/teams
type Team struct {
	EnableSelfAssignment bool     `json:"enableSelfAssignment"`
	Hub                  *string  `json:"hub"`
	ID                   string   `json:"id"`
	Managers             []string `json:"managers"`
	Name                 string   `json:"name"`
	Tasks                []string `json:"tasks"`
	TimeCreated          int64    `json:"timeCreated"`
	TimeLastModified     int64    `json:"timeLastModified"`
	Workers              []string `json:"workers"`
}

// Reference https://docs.onfleet.com/reference/create-team
type TeamCreateParams struct {
	EnableSelfAssignment bool     `json:"enableSelfAssignment"`
	Hub                  string   `json:"hub,omitempty"`
	Managers             []string `json:"managers"`
	Name                 string   `json:"name"`
	Workers              []string `json:"workers"`
}

// Reference https://docs.onfleet.com/reference/update-team
type TeamUpdateParams struct {
	EnableSelfAssignment bool     `json:"enableSelfAssignment"`
	Hub                  string   `json:"hub,omitempty"`
	Managers             []string `json:"managers"`
	Name                 string   `json:"name"`
	Workers              []string `json:"workers"`
}

// Reference https://docs.onfleet.com/reference/team-auto-dispatch
type TeamAutoDispatch struct {
	DispatchId string `json:"dispatchId"`
}

// Reference https://docs.onfleet.com/reference/team-auto-dispatch
type TeamAutoDispatchParams struct {
	MaxAllowedDelay    int     `json:"maxAllowedDelay,omitempty"`
	MaxTasksPerRoute   int     `json:"maxTasksPerRoute,omitempty"`
	RouteEnd           string  `json:"routeEnd,omitempty"`
	ScheduleTimeWindow []int64 `json:"scheduleTimeWindow,omitempty"`
	ServiceTime        int     `json:"serviceTime,omitempty"`
	TaskTimeWindow     []int64 `json:"taskTimeWindow,omitempty"`
}

// Reference https://docs.onfleet.com/reference/delivery-estimate
type TeamWorkerEta struct {
	Steps    []TeamWorkerEtaStep `json:"steps"`
	Vehicle  WorkerVehicleType   `json:"vehicle"`
	WorkerId string              `json:"workerId"`
}

// Reference https://docs.onfleet.com/reference/delivery-estimate
type TeamWorkerEtaStep struct {
	CompletionTime int64               `json:"completionTime"`
	Distance       float64             `json:"distance"`
	Location       DestinationLocation `json:"location"`
	ServiceTime    float64             `json:"serviceTime"`
	TravelTime     float64             `json:"travelTime"`
}

// Reference https://docs.onfleet.com/reference/delivery-estimate
type TeamWorkerEtaQueryParams struct {
	DropoffLocation         string            `json:"dropoffLocation,omitempty"`
	PickupLocation          string            `json:"pickupLocation,omitempty"`
	PickupTime              int64             `json:"pickupTime,omitempty,string"`
	RestrictedVehiclesTypes WorkerVehicleType `json:"restrictedVehiclesTypes,omitempty"`
	ServiceTime             float64           `json:"serviceTime,omitempty,string"`
}

type TeamTasks struct {
	LastId string `json:"lastId,omitempty"`
	Tasks  []Task `json:"tasks"`
}

type TeamTasksListQueryParams struct {
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
