package usecase

type SchemaInputDTO struct {
	SchemaType string                 `json:"schema_type"`
	Service    string                 `json:"service"`
	Source     string                 `json:"source"`
	JsonSchema map[string]interface{} `json:"json_schema"`
}

type SchemaOutputDTO struct {
	ID         string                 `json:"id"`
	SchemaType string                 `json:"schema_type"`
	Service    string                 `json:"service"`
	Source     string                 `json:"source"`
	JsonSchema map[string]interface{} `json:"json_schema"`
	SchemaID   string                 `json:"schema_id"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"updated_at"`
}
