package entity

type MetadataInputOrigin struct {
	Gateway    string `json:"gateway"`
	Controller string `json:"controller"`
}

type MetadataInput struct {
	Data                map[string]interface{} `json:"data"`
	ProcessingId        string                 `json:"processing_id"`
	ProcessingTimestamp string                 `json:"processing_timestamp"`
	Source              MetadataInputOrigin    `json:"source"`
}

type ServiceFeedback struct {
	Data                map[string]interface{} `json:"data"`
	Metadata            MetadataInput          `json:"metadata"`
	Service             MetadataInputOrigin    `json:"service"`
	ProcessingId        string                 `json:"processing_id"`
	ProcessingTimestamp string                 `json:"processing_timestamp"`
}

func NewServiceFeedback(
	data map[string]interface{},
	metadataData map[string]interface{},
	metadataSourceGateway string,
	metadataSourceController string,
	metadataServiceGateway string,
	metadataServiceController string,
	service MetadataInputOrigin,
	processingId string,
	processingIdMetadata string,
	processingTimestamp string,
	processingTimestampMetadata string,
) *ServiceFeedback {
	return &ServiceFeedback{
		Data: data,
		Metadata: MetadataInput{
			Data: metadataData,
			Source: MetadataInputOrigin{
				Gateway:    metadataSourceGateway,
				Controller: metadataSourceController,
			},
			ProcessingId:        processingIdMetadata,
			ProcessingTimestamp: processingTimestampMetadata,
		},
		Service: MetadataInputOrigin{
			Gateway:    metadataServiceGateway,
			Controller: metadataServiceController,
		},
		ProcessingId:        processingId,
		ProcessingTimestamp: processingTimestamp,
	}
}
