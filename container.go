package onfleet

type Container struct {
	ActiveTask       *string       `json:"activeTask"`
	ID               string        `json:"id"`
	Organization     string        `json:"organization"`
	TimeCreated      int64         `json:"timeCreated"`
	TimeLastModified int64         `json:"timeLastModified"`
	Type             ContainerType `json:"type"`
	Tasks            []string      `json:"tasks"`
	Worker           string        `json:"worker,omitempty"`
	Team             string        `json:"team,omitempty"`
}

type ContainerType string

const (
	ContainerTypeOrganization ContainerType = "ORGANIZATION"
	ContainerTypeTeam         ContainerType = "TEAM"
	ContainerTypeWorker       ContainerType = "WORKER"
)
