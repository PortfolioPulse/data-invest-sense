from jobs.handlers.handlers import Handler
from pylog.log import setup_logging
from configs.loader import JobParams, JobMetadataParams
from pyrepository.interfaces.ingestors.dtos import MessageStatus
from datetime import datetime
from dataclasses import dataclass
import requests
from pyminio.pyminio import minio_client


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
    def __init__(self, job_params: JobParams, job_metadata_params: JobMetadataParams):
        self.job_params = job_params
        self.job_metadata_params = job_metadata_params
        super().__init__(job_params, job_metadata_params)

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

    def _get_bucket_name(self):
        return "raw-{context}-source-{source}".format(
            context=self.job_metadata_params.context,
            source=self.job_metadata_params.source,
        )

    def _get_status(self, response):
        return MessageStatus(
            code=response.status_code,
            detail=response.reason,
        )

    def make_request(self, input: JobInput):
        endpoint = self._get_endpoint(input)
        logger.info(f"endpoint: {endpoint}")
        return requests.get(endpoint, verify=False)

    def run(self, data):
        input = self._get_job_input(data)
        logger.info(f"Job triggered with input: {input}")
        response = self.make_request(input)
        minio = minio_client()
        logger.info(f"Job _get_endpoint: {response.status_code}")
        uri = minio.upload_bytes(self._get_bucket_name(), "test-object", response.content)
        logger.info(f"File storage uri: {uri}")
        return {"documentUri": uri}, self._get_status(response)
