package entity

import (
	"errors"
	"time"

	"libs/golang/go-jwt/jwt"
	"libs/golang/goid/md5"
	schemaJWTID "libs/golang/goid/schema"

	"github.com/go-chi/jwtauth"
)

type Schema struct {
	ID         schemaJWTID.ID         `json:"id"`
	SchemaType string                 `bson:"schema_type"`
	JsonSchema map[string]interface{} `bson:"json_schema"`
	SchemaID   md5.ID                 `bson:"schema_id"`
	Service    string                 `bson:"service"`
	Source     string                 `bson:"source"`
	CreatedAt  string
	UpdatedAt  string
}

func NewSchema(schemaType string, service string, source string, jsonSchema map[string]interface{}, tokenAuth *jwtauth.JWTAuth) (*Schema, error) {
	jwtToken, err := jwt.GenerateSchemaJWT(tokenAuth, jsonSchema)

	if err != nil {
		return nil, err
	}
	md5Schema := md5.NewMd5Hash(jwtToken)
	// schemaId, err := uuid.ParseID(md5Schema)
	// if err != nil {
	// 	return nil, err
	// }
	schema := &Schema{
		ID:         schemaJWTID.NewID(schemaType, service, source),
		SchemaType: schemaType,
		JsonSchema: jsonSchema,
		Service:    service,
		Source:     source,
		SchemaID:   md5Schema,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	err = schema.IsSchemaValid()
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (schema *Schema) IsSchemaValid() error {
	if schema.SchemaType == "" {
		return errors.New("schemaType is empty")
	}
	if schema.Service == "" {
		return errors.New("service is empty")
	}
	if schema.Source == "" {
		return errors.New("source is empty")
	}
	if schema.JsonSchema == nil {
		return errors.New("jsonSchema is empty")
	}
	return nil
}
