package resource

// Onfleet Recipient
// Reference https://docs.onfleet.com/reference/recipients
type Recipient struct {
	ID                   string     `json:"id,omitempty"`
	TimeCreated          int64      `json:"timeCreated,omitempty"`
	TimeLastModified     int64      `json:"timeLastModified,omitempty"`
	Metadata             []Metadata `json:"metadata,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Notes                string     `json:"notes,omitempty"`
	Organization         string     `json:"organization,omitempty"`
	Phone                string     `json:"phone,omitempty"`
	SkipSmsNotifications bool       `json:"skipSMSNotifications,omitempty"`
}
