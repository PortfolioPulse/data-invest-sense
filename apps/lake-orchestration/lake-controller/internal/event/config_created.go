package event

import "time"

type ConfigCreated struct {
	Name    string
	Payload interface{}
}

func NewConfigCreated() *ConfigCreated {
	return &ConfigCreated{
		Name: "ConfigCreated",
	}
}

func (e *ConfigCreated) GetName() string {
	return e.Name
}

func (e *ConfigCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *ConfigCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ConfigCreated) GetDateTime() time.Time {
	return time.Now()
}
