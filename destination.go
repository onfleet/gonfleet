package onfleet

type Destination struct {
	Address          DestinationAddress  `json:"address"`
	GooglePlaceId    string              `json:"googlePlaceId"`
	ID               string              `json:"id"`
	Location         DestinationLocation `json:"location"`
	Metadata         []Metadata          `json:"metadata"`
	Notes            string              `json:"notes"`
	TimeCreated      int64               `json:"timeCreated"`
	TimeLastModified int64               `json:"timeLastModified"`
	Warnings         []any               `json:"warnings"`
}

type DestinationLocation []float64

type DestinationAddress struct {
	Apartment  string `json:"apartment"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Name       string `json:"name,omitempty"`
	Number     string `json:"number"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
	Street     string `json:"street"`
	Unparsed   string `json:"unparsed,omitempty"`
}

type DestinationCreateParams struct {
	Address  DestinationAddress       `json:"address"`
	Location DestinationLocation      `json:"location,omitempty"`
	Metadata []Metadata               `json:"metadata,omitempty"`
	Notes    string                   `json:"notes,omitempty"`
	Options  *DestinationOptionsParam `json:"options,omitempty"`
}

type DestinationOptionsParam struct {
	Language string `json:"language,omitempty"`
}
