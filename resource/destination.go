package resource

type Destination struct {
	Address          DestinationAddress  `json:"address,omitempty"`
	GooglePlaceId    string              `json:"googlePlaceId,omitempty"`
	ID               string              `json:"id,omitempty"`
	Location         DestinationLocation `json:"location,omitempty"`
	Metadata         []Metadata          `json:"metadata,omitempty"`
	Notes            string              `json:"notes,omitempty"`
	TimeCreated      int64               `json:"timeCreated,omitempty"`
	TimeLastModified int64               `json:"timeLastModified,omitempty"`
	Warnings         []any               `json:"warnings,omitempty"`
}

// Location is longitude and latitude.
// In that order :)
type DestinationLocation [2]float32

type DestinationAddress struct {
	Apartment  string `json:"apartment,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	Name       string `json:"name,omitempty"`
	Number     string `json:"number,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State      string `json:"state,omitempty"`
	Street     string `json:"street,omitempty"`
	Unparsed   string `json:"unparsed,omitempty"`
}

type DestinationCreationParams struct {
	Address  DestinationAddress  `json:"address,omitempty"`
	Location DestinationLocation `json:"location,omitempty"`
	Notes    string              `json:"notes,omitempty"`
}
