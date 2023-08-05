package event

import "time"

type SchemaInputCreated struct {
	Name    string
	Payload interface{}
}

func NewSchemaInputCreated() *SchemaInputCreated {
	return &SchemaInputCreated{
		Name: "SchemaInputCreated",
	}
}

func (e *SchemaInputCreated) GetName() string {
	return e.Name
}

func (e *SchemaInputCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *SchemaInputCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *SchemaInputCreated) GetDateTime() time.Time {
	return time.Now()
}
