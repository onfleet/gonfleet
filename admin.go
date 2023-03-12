package onfleet

// Onfleet Administrator.
// Reference https://docs.onfleet.com/reference/administrators.
type Admin struct {
	Email            string     `json:"email"`
	ID               string     `json:"id"`
	IsAccountOwner   bool       `json:"isAccountOwner"`
	IsActive         bool       `json:"isActive"`
	IsReadOnly       bool       `json:"isReadOnly"`
	Metadata         []Metadata `json:"metadata"`
	Name             string     `json:"name"`
	Organization     string     `json:"organization"`
	Phone            string     `json:"phone"`
	Teams            []string   `json:"teams"`
	TimeCreated      int64      `json:"timeCreated"`
	TimeLastModified int64      `json:"timeLastModified"`
	Type             string     `json:"type"`
}
