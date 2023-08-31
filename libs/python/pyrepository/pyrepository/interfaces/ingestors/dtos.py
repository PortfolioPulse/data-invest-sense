from dataclasses import dataclass


@dataclass
class JobMetadataParams:
    _id: str
    name: str
    context: str
    source: str
    service: str

@dataclass
class JobParams:
    jobHandler: str
    active: bool
    url: str

@dataclass
class Config:
    jobMetadataParams: JobMetadataParams
    jobParams: JobParams

@dataclass
class MessageMetadata:
    processing_id: str
    processing_timestamp: str
    source: str
    service: str

@dataclass
class MessageParameters:
    id: str
    data: dict
    metadata: MessageMetadata


@dataclass
class MessageMetadataOriginsOutput:
    gateway: str
    controller: str

@dataclass
class MessageMetadataOutput:
    id: str
    data: dict
    processing_id: str
    processing_timestamp: str
    source: MessageMetadataOriginsOutput

@dataclass
class MetadataOutput:
    input: MessageMetadataOutput
    service: MessageMetadataOriginsOutput
    processing_id: str
    processing_timestamp: str

@dataclass
class MessageStatus:
    code: int
    detail: str
