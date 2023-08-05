from abc import ABC, abstractmethod
from configs.loader import JobParams

class Handler(ABC):
    def __init__(self, job_params: JobParams):
        self.job = job_params

    # @abstractmethod
    def _get_endpoint(self):
        raise NotImplementedError

    # @abstractmethod
    def run(self):
        raise NotImplementedError
