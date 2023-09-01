package event

import "time"

type StagingJobCreated struct {
     Name    string
     Payload interface{}
}

func NewStagingJobCreated() *StagingJobCreated {
     return &StagingJobCreated{
          Name: "StagingJobCreated",
     }
}

func (e *StagingJobCreated) GetName() string {
     return e.Name
}

func (e *StagingJobCreated) GetPayload() interface{} {
     return e.Payload
}

func (e *StagingJobCreated) SetPayload(payload interface{}) {
     e.Payload = payload
}

func (e *StagingJobCreated) GetDateTime() time.Time {
     return time.Now()
}
