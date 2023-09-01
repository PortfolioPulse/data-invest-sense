package usecase

type Metadata struct {
	ProcessingId        string `json:"processing_id"`
	ProcessingTimestamp string `json:"processing_timestamp"`
	Source              string `json:"source"`
	Service             string `json:"service"`
}

type Status struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type InputInputDTO struct {
	Data map[string]interface{} `json:"data"`
}

type InputOutputDTO struct {
	ID       string                 `json:"id"`
	Data     map[string]interface{} `json:"data"`
	Metadata Metadata               `json:"metadata"`
	Status   Status                 `json:"status"`
}

type InputStatusInputDTO struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
}

type ListInputOutputDTO struct {
	Inputs []InputOutputDTO `json:"inputs"`
}

type StagingJobInputDTO struct {
	InputId      string                 `json:"input_id"`
	Input        map[string]interface{} `json:"input"`
	Source       string                 `json:"source"`
	Service      string                 `json:"service"`
	ProcessingId string                 `json:"processing_id"`
}

type StagingJobOutputDTO struct {
	ID           string                 `json:"id"`
	InputId      string                 `json:"input_id"`
	Input        map[string]interface{} `json:"input"`
	Source       string                 `json:"source"`
	Service      string                 `json:"service"`
	ProcessingId string                 `json:"processing_id"`
}
