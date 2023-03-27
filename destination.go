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

// Location is longitude and latitude.
// In that order :)
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
	Address  DestinationAddress  `json:"address"`
	Location DestinationLocation `json:"location,omitempty"`
	Notes    string              `json:"notes,omitempty"`
	Options  *DestinationOptions `json:"options,omitempty"`
}

type DestinationOptions struct {
	// Language is a ISO standard two letter country code
	Language string `json:"language,omitempty"`
}
