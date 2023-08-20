package event

import "time"

type SchemaCreated struct {
     Name    string
     Payload interface{}
}

func NewSchemaCreated() *SchemaCreated {
     return &SchemaCreated{
          Name: "SchemaCreated",
     }
}

func (e *SchemaCreated) GetName() string {
     return e.Name
}

func (e *SchemaCreated) GetPayload() interface{} {
     return e.Payload
}

func (e *SchemaCreated) SetPayload(payload interface{}) {
     e.Payload = payload
}

func (e *SchemaCreated) GetDateTime() time.Time {
     return time.Now()
}
