package worker

import (
	"net/http"

	"github.com/onfleet/gonfleet/config"
	"github.com/onfleet/gonfleet/resource/metadata"
)

// Client for Workers resource
type Client struct {
	Config     config.Config
	HttpClient *http.Client
	SubPath    string
}

// Worker refers to an Onfleet Worker.
// Reference https://docs.onfleet.com/reference/workers.
type Worker struct {
	ID               string              `json:"id"`
	TimeCreated      int64               `json:"timeCreated"`
	TimeLastModified int64               `json:"timeLastModified"`
	Organization     string              `json:"organization"`
	Name             string              `json:"name"`
	DisplayName      string              `json:"displayName"`
	Phone            string              `json:"phone"`
	ActiveTask       *string             `json:"activeTask"`
	Tasks            []string            `json:"tasks"`
	OnDuty           bool                `json:"onDuty"`
	TimeLastSeen     int64               `json:"timeLastSeen"`
	Capacity         int                 `json:"capacity"`
	UserData         workerUserData      `json:"userData"`
	AccountStatus    string              `json:"accountStatus"`
	Metadata         []metadata.Metadata `json:"metadata"`
}

type workerUserData struct {
	AppVersion        string `json:"appVersion"`
	BatteryLevel      int    `json:"batteryLevel"`
	DeviceDescription string `json:"deviceDescription"`
	Platform          string `json:"platform"`
}
