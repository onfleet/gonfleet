package onfleet

type Organization struct {
	Country            string   `json:"country"`
	Delegatees         []string `json:"delegatees"`
	DriverSupportEmail string   `json:"driverSupportEmail"`
	Email              string   `json:"email"`
	ID                 string   `json:"id"`
	Image              string   `json:"image,omitempty"`
	Name               string   `json:"name"`
	TimeCreated        int64    `json:"timeCreated"`
	TimeLastModified   int64    `json:"timeLastModified"`
	Timezone           string   `json:"timezone"`
}

type OrganizationDelegate struct {
	Country            string `json:"country"`
	DriverSupportEmail string `json:"driverSupportEmail"`
	Email              string `json:"email"`
	ID                 string `json:"id"`
	IsFulfillment      bool   `json:"isFulfillment"`
	Name               string `json:"name"`
	Timezone           string `json:"timezone"`
}
