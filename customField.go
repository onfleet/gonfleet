package onfleet

type CustomField struct {
	Description string                        `json:"description"`
	AsArray     bool                          `json:"asArray"`
	Visibility  []CustomFieldVisibilityOption `json:"visibility"`
	Editability []CustomFieldVisibilityOption `json:"editability"`
	Key         string                        `json:"key"`
	Name        string                        `json:"name"`
	Type        CustomFieldValidDataType      `json:"type"`
	Contexts    []CustomFieldContext          `json:"context"`
	Value       any                           `json:"value"`
}

type CustomFieldParams struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type CustomFieldContext struct {
	IsRequired bool   `json:"isRequired"`
	Conditions []any  `json:"conditions"`
	Name       string `json:"name"`
}

type CustomFieldVisibilityOption string
type CustomFieldValidDataType string

const (
	CustomFieldVisibilityOptionAdmin  = "admin"
	CustomFieldVisibilityOptionAPI    = "api"
	CustomFieldVisibilityOptionWorker = "worker"

	CustomFieldValidDataTypeSingleLineText = "single_line_text_field"
	CustomFieldValidDataTypeMultiLineText  = "multi_line_text_field"
	CustomFieldValidDataTypeBoolean        = "boolean"
	CustomFieldValidDataTypeInteger        = "integer"
	CustomFieldValidDataTypeDecimal        = "decimal"
	CustomFieldValidDataTypeDate           = "date"
	CustomFieldValidDataTypeURL            = "Url"
)
