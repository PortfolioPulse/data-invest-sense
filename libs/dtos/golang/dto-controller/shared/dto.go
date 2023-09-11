package shared

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type ProcessingJobDependencies struct {
	Service             string `json:"service"`
	Source              string `json:"source"`
	ProcessingId        string `json:"processing_id"`
	ProcessingTimestamp string `json:"processing_timestamp"`
	StatusCode          int `json:"status_code"`
}

