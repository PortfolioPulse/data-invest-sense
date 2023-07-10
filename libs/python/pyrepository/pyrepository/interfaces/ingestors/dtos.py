from dataclasses import dataclass


@dataclass
class Config:
    id: str
    name: str
    jobType: str
    context: str
    outputType: str


@dataclass
class MessageParameters:
    input: dict
    metadata: dict
