package onfleet

// Onfleet Recipient
// Reference https://docs.onfleet.com/reference/recipients
type Recipient struct {
	ID                   string     `json:"id"`
	TimeCreated          int64      `json:"timeCreated"`
	TimeLastModified     int64      `json:"timeLastModified"`
	Metadata             []Metadata `json:"metadata"`
	Name                 string     `json:"name"`
	Notes                string     `json:"notes"`
	Organization         string     `json:"organization"`
	Phone                string     `json:"phone"`
	SkipSmsNotifications bool       `json:"skipSMSNotifications"`
}
