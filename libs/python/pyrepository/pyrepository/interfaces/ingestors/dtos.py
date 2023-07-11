from dataclasses import dataclass


@dataclass
class Config:
    id: str
    name: str
    active: bool
    jobType: str
    shouldRun: bool
    context: str
    outputType: str


@dataclass
class MessageParameters:
    input: dict
    metadata: dict
