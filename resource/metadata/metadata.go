package metadata

// VisibilityOption indicates who / what is able to view the details on a metadata object.
// Options are "api", "dashboard", "worker"
type VisibilityOption string

const (
	VisibilityOptionApi       VisibilityOption = "api"
	VisibilityOptionDashboard VisibilityOption = "dashboard"
	VisibilityOptionWorker    VisibilityOption = "worker"
)

// Onfleet Metadata.
// Reference https://docs.onfleet.com/reference/metadata
type Metadata struct {
	Name string `json:"name,omitempty"`
	// Type can be one of the following "boolean", "number", "string", "object", "array"
	Type string `json:"type,omitempty"`
	// Subtype only required for Type of "array"
	// And can be one of the following "boolean", "number", "string", "object"
	Subtype string `json:"subtype,omitempty"`
	// Visibility lists who / what can view the metadata.
	// Options are "api", "dashboard", "worker"
	Visibility []VisibilityOption `json:"visibility,omitempty"`
	// Value is any user assigned data that matches Type / Subtype
	Value any `json:"value,omitempty"`
}

type MatchedMetadataResult struct {
	ID       string     `json:"id,omitempty"`
	Metadata []Metadata `json:"metadata,omitempty"`
}
