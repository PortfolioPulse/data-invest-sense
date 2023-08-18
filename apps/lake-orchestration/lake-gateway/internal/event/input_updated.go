package event

import "time"

type InputUpdated struct {
	Name    string
	Payload interface{}
}

func NewInputUpdated() *InputUpdated {
	return &InputUpdated{
		Name: "InputCreated",
	}
}

func (e *InputUpdated) GetName() string {
	return e.Name
}

func (e *InputUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *InputUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *InputUpdated) GetDateTime() time.Time {
	return time.Now()
}
