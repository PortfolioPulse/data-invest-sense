package shared

type Status struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type Metadata struct {
	ProcessingId        string `json:"processing_id"`
	ProcessingTimestamp string `json:"processing_timestamp"`
	Source              string `json:"source"`
	Service             string `json:"service"`
}
