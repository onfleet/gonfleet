package onfleet

// Onfleet Metadata.
// Reference https://docs.onfleet.com/reference/metadata
type Metadata struct {
	Name       string                     `json:"name"`
	Subtype    string                     `json:"subtype,omitempty"`
	Type       string                     `json:"type"`
	Value      any                        `json:"value"`
	Visibility []MetadataVisibilityOption `json:"visibility,omitempty"`
}

type MetadataVisibilityOption string

const (
	MetadataVisibilityOptionApi       MetadataVisibilityOption = "api"
	MetadataVisibilityOptionDashboard MetadataVisibilityOption = "dashboard"
	MetadataVisibilityOptionWorker    MetadataVisibilityOption = "worker"
)
