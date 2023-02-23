package metadata

// VisibilityOption indicates who / what is able to view the details on a metadata object.
// Options are "api", "dashboard", "worker"
type VisibilityOption string

const (
	VisibilityOptionApi       VisibilityOption = "api"
	VisibilityOptionDashboard VisibilityOption = "dashboard"
	VisibilityOptionWorker    VisibilityOption = "worker"
)

type Metadata struct {
	// Name of the metadata object
	Name string `json:"name,omitempty"`
	// Type can be one of the following "boolean", "number", "string", "object", "array"
	Type string `json:"type"`
	// Subtype only required for Type of "array"
	// And can be one of the following "boolean", "number", "string", "object"
	Subtype string `json:"subtype,omitempty"`
	// Visibility lists who / what can view the metadata.
	// Options are "api", "dashboard", "worker"
	Visibility []VisibilityOption `json:"visibility"`
	// Value is any user assigned data that matches Type / Subtype
	Value any `json:"value"`
}

type MatchedMetadataResult struct {
	ID       string     `json:"id"`
	Metadata []Metadata `json:"metadata"`
}
