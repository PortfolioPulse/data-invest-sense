from jobs.handlers.handlers import Handler
from pylog.log import setup_logging
from configs.loader import JobParams
import time
from datetime import datetime
from dataclasses import dataclass
import requests


logger = setup_logging(__name__)


@dataclass
class Reference:
    year: int
    month: int
    day: int


@dataclass
class JobInput:
    reference: Reference


class Job(Handler):
    def __init__(self, job_params: JobParams):
        self.job_params = job_params
        super().__init__(job_params)

    def _get_job_input(self, data):
        try:
            return JobInput(
                reference=Reference(**data.data["reference"])
            )
        except KeyError as err:
            logger.error(f"Invalid job input: {err}")
            raise ValueError("Invalid job input")

    def _get_reference(self, reference: Reference):
        ref = datetime(reference.year, reference.month, reference.day)
        return ref.strftime("%Y%m%d")

    def _get_endpoint(self, input: JobInput):
        reference = self._get_reference(input.reference)
        return self.job_params.url.format(reference)

    def make_request(self, input: JobInput):
        endpoint = self._get_endpoint(input)
        logger.info(f"endpoint: {endpoint}")
        return requests.get(endpoint, verify=False)

    def run(self, data):
        input = self._get_job_input(data)
        logger.info("sleeping for 30 seconds")
        # time.sleep(30)
        logger.info(f"Job triggered with input: {input}")
        response = self.make_request(input)
        logger.info(f"Job _get_endpoint: {response.status_code}")


