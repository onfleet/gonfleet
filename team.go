package onfleet

// Onfleet Team.
// Reference https://docs.onfleet.com/reference/teams.
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

type TeamCreateParams struct {
	EnableSelfAssignment bool     `json:"enableSelfAssignment"`
	Hub                  string   `json:"hub,omitempty"`
	Managers             []string `json:"managers"`
	Name                 string   `json:"name"`
	Workers              []string `json:"workers"`
}

type TeamUpdateParams struct {
	EnableSelfAssignment bool     `json:"enableSelfAssignment"`
	Hub                  string   `json:"hub,omitempty"`
	Managers             []string `json:"managers"`
	Name                 string   `json:"name"`
	Workers              []string `json:"workers"`
}
