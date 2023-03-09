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
type DestinationLocation [2]float32

type DestinationAddress struct {
	Apartment  string `json:"apartment"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Name       string `json:"name"`
	Number     string `json:"number"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
	Street     string `json:"street"`
	Unparsed   string `json:"unparsed,omitempty"`
}

type DestinationCreationParams struct {
	Address  DestinationAddress  `json:"address"`
	Location DestinationLocation `json:"location"`
	Notes    string              `json:"notes"`
}