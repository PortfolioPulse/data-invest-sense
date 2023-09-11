package input

type SchemaDTO struct {
	SchemaType string                 `json:"schema_type"`
	Service    string                 `json:"service"`
	Source     string                 `json:"source"`
	JsonSchema map[string]interface{} `json:"json_schema"`
}
