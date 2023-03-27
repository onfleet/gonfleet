package onfleet

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

type RecipientCreateParams struct {
	Metadata                  []Metadata `json:"metadata,omitempty"`
	Name                      string     `json:"name,omitempty"`
	Notes                     string     `json:"notes,omitempty"`
	Phone                     string     `json:"phone,omitempty"`
	SkipPhoneNumberValidation bool       `json:"skipPhoneNumberValidation,omitempty"`
	SkipSmsNotifications      bool       `json:"skipSMSNotifications,omitempty"`
	UseLongCodeForText        bool       `json:"useLongCodeForText,omitempty"`
}

type RecipientUpdateParams struct {
	Metadata             []Metadata `json:"metadata,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Notes                string     `json:"notes,omitempty"`
	SkipSmsNotifications bool       `json:"skipSMSNotifications,omitempty"`
}

type RecipientQueryKey string

const (
	RecipientQueryKeyName  RecipientQueryKey = "name"
	RecipientQueryKeyPhone RecipientQueryKey = "phone"
)
