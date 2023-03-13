package onfleet

type Hub struct {
	Address  DestinationAddress  `json:"address"`
	ID       string              `json:"id"`
	Location DestinationLocation `json:"location"`
	Name     string              `json:"name"`
	Teams    []string            `json:"teams"`
}

type HubCreateParams struct {
	Address DestinationAddress `json:"address"`
	Name    string             `json:"name"`
	Teams   []string           `json:"teams,omitempty"`
}

type HubUpdateParams struct {
	Address DestinationAddress `json:"address"`
	Name    string             `json:"name"`
	Teams   []string           `json:"teams,omitempty"`
}
