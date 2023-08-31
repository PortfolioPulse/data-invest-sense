from abc import ABC, abstractmethod
from configs.loader import JobParams, JobMetadataParams

class Handler(ABC):
    def __init__(self, job_params: JobParams, job_metadata_params: JobMetadataParams):
        self.job = job_params
        self.job_metadata = job_metadata_params

    # @abstractmethod
    def _get_endpoint(self):
        raise NotImplementedError

    # @abstractmethod
    def run(self):
        raise NotImplementedError
