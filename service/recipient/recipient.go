package recipient

import (
	"github.com/onfleet/gonfleet/resource/metadata"
)

type Recipient struct {
	ID                   string              `json:"id,omitempty"`
	TimeCreated          int64               `json:"timeCreated,omitempty"`
	TimeLastModified     int64               `json:"timeLastModified,omitempty"`
	Name                 string              `json:"name,omitempty"`
	Phone                string              `json:"phone,omitempty"`
	Organization         string              `json:"organization,omitempty"`
	SkipSmsNotifications bool                `json:"skipSMSNotifications,omitempty"`
	Metadata             []metadata.Metadata `json:"metadata,omitempty"`
}