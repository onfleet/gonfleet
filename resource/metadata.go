package resource

// Onfleet Metadata.
// Reference https://docs.onfleet.com/reference/metadata
type Metadata struct {
	Name       string                     `json:"name,omitempty"`
	Subtype    string                     `json:"subtype,omitempty"`
	Type       string                     `json:"type,omitempty"`
	Value      any                        `json:"value,omitempty"`
	Visibility []MetadataVisibilityOption `json:"visibility,omitempty"`
}

type MetadataVisibilityOption string

const (
	MetadataVisibilityOptionApi       MetadataVisibilityOption = "api"
	MetadataVisibilityOptionDashboard MetadataVisibilityOption = "dashboard"
	MetadataVisibilityOptionWorker    MetadataVisibilityOption = "worker"
)

type MetadataMatchedResult struct {
	ID       string     `json:"id,omitempty"`
	Metadata []Metadata `json:"metadata,omitempty"`
}
