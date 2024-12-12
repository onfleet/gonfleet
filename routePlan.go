package onfleet

type PositionEnum string

const (
	PositionEnumHub            PositionEnum = "HUB"
	PositionEnumWorkerLocation PositionEnum = "WORKER_LOCATION"
	PositionEnumWorkerAddress  PositionEnum = "WORKER_ADDRESS"
)

type RoutePlanParams struct {
	Name          string       `json:"name"`
	StartTime     int64        `json:"startTime"`
	TaskIds       []string     `json:"tasks,omitempty"`
	Color         string       `json:"color,omitempty"`
	VehicleType   string       `json:"vehicleType,omitempty"`
	Worker        string       `json:"worker,omitempty"`
	Team          string       `json:"team,omitempty"`
	StartAt       PositionEnum `json:"start,omitempty"`
	EndAt         PositionEnum `json:"end,omitempty"`
	StartingHubId string       `json:"startingHubId,omitempty"`
	EndingHubId   string       `json:"endingHubId,omitempty"`
	EndTime       int64        `json:"endTime,omitempty"`
	Timezone      string       `json:"timezone,omitempty"`
}

type RoutePlan struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	State            string   `json:"state"`
	Color            string   `json:"color"`
	Tasks            []string `json:"tasks"`
	Organization     string   `json:"organization"`
	Team             *string  `json:"team"`
	Worker           string   `json:"worker"`
	VehicleType      string   `json:"vehicleType"`
	StartTime        int64    `json:"startTime"`
	EndTime          *int64   `json:"endTime"`
	ActualStartTime  *int64   `json:"actualStartTime"`
	ActualEndTime    *int64   `json:"actualEndTime"`
	StartingHubId    *string  `json:"startingHubId"`
	EndingHubId      *string  `json:"endingHubId"`
	ShortId          string   `json:"shortId"`
	TimeCreated      int64    `json:"timeCreated"`
	TimeLastModified int64    `json:"timeLastModified"`
}

type RoutePlanListQueryParams struct {
	WorkerId        string `json:"workerId,omitempty"`
	StartTimeTo     int64  `json:"startTimeTo,omitempty"`
	StartTimeFrom   int64  `json:"startTimeFrom,omitempty"`
	CreatedTimeTo   int64  `json:"createdTimeTo,omitempty"`
	CreatedTimeFrom int64  `json:"createdTimeFrom,omitempty"`
	HasTasks        bool   `json:"hasTasks,omitempty"`
	Limit           int64  `json:"limit,omitempty"`
}

type RoutePlanAddTasksParams struct {
	Tasks []string `json:"tasks"`
}

type RoutePlansPaginated struct {
	LastId     string      `json:"lastId,omitempty"`
	RoutePlans []RoutePlan `json:"routePlans"`
}
