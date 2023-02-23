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
	Name       string             `json:"name"`
	Type       string             `json:"type"`              // "boolean" ; "number" ; "string" ; "object" ; "array"
	Subtype    string             `json:"subtype,omitempty"` // "boolean" ; "number" ; "string" ; "object" ; "array"
	Visibility []VisibilityOption `json:"visibility"`        // "api" ; "dashboard" ; "worker"
	Value      any                `json:"value"`
}

type MatchedMetadataResult struct {
	ID       string     `json:"id"`
	Metadata []Metadata `json:"metadata"`
}
