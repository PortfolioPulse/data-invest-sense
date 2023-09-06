package shared

type MetadataInputOrigin struct {
	Gateway    string `json:"gateway"`
	Controller string `json:"controller"`
}

type MetadataInput struct {
	ID                  string                 `json:"id"`
	Data                map[string]interface{} `json:"data"`
	ProcessingId        string                 `json:"processing_id"`
	ProcessingTimestamp string                 `json:"processing_timestamp"`
	Source              MetadataInputOrigin    `json:"source"`
}

type Metadata struct {
	Input               MetadataInput       `json:"input"`
	Service             MetadataInputOrigin `json:"service"`
	ProcessingId        string              `json:"processing_id"`
	ProcessingTimestamp string              `json:"processing_timestamp"`
	TargetEndpoint      string              `json:"target_endpoint"`
	JobFrequency        string              `json:"job_frequency"`
}

type Status struct {
     Code   int    `json:"code"`
     Detail string `json:"detail"`
}
