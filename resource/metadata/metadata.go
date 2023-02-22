package metadata

type Metadata struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`                 // "boolean" | "number" | "string" | "object" | "array"
	Subtype    string   `json:"subtype,omitempty"`    // "boolean" | "number" | "string" | "object" | "array"
	Visibility []string `json:"visibility,omitempty"` // "api" | "dashboard" | "worker"
	Value      any      `json:"value"`
}

type MatchedMetadataResult struct {
	ID       string     `json:"id"`
	Metadata []Metadata `json:"metadata"`
}
