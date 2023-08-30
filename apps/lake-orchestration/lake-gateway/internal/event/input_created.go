package event

import "time"

type InputCreated struct {
	Name    string
	Payload interface{}
}

func NewInputCreated() *InputCreated {
	return &InputCreated{
		Name: "InputCreated",
	}
}

func (e *InputCreated) GetName() string {
	return e.Name
}

func (e *InputCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *InputCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *InputCreated) GetDateTime() time.Time {
	return time.Now()
}
