package types

// VisibilityOption indicates who / what is able to view the details on a metadata object.
// Options are "api", "dashboard", "worker"
type MetadataVisibilityOption string

const (
	MetadataVisibilityOptionApi       MetadataVisibilityOption = "api"
	MetadataVisibilityOptionDashboard MetadataVisibilityOption = "dashboard"
	MetadataVisibilityOptionWorker    MetadataVisibilityOption = "worker"
)

// Onfleet Metadata.
// Reference https://docs.onfleet.com/reference/metadata
type Metadata struct {
	Name string `json:"name,omitempty"`
	// Subtype only required for Type of "array"
	// And can be one of the following "boolean", "number", "string", "object"
	Subtype string `json:"subtype,omitempty"`
	// Type can be one of the following "boolean", "number", "string", "object", "array"
	Type string `json:"type,omitempty"`
	// Value is any user assigned data that matches Type / Subtype
	Value any `json:"value,omitempty"`
	// Visibility lists who / what can view the metadata.
	// Options are "api", "dashboard", "worker"
	Visibility []MetadataVisibilityOption `json:"visibility,omitempty"`
}

type MetadataMatchedResult struct {
	ID       string     `json:"id,omitempty"`
	Metadata []Metadata `json:"metadata,omitempty"`
}
