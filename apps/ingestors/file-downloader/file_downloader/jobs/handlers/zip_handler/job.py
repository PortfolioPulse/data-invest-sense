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
        return "landing-{context}-source-{source}".format(
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
        headers = {
            "Referer": "https://www.portaltransparencia.gov.br/download-de-dados/ceaf",
            "Sec-Fetch-Site": "same-origin",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
            "Accept-Encoding": "gzip, deflate, br",
            "Accept-Language": "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
            "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0",
            "Cookie": "NSC_JOibrajdb4h3qgcckulqmuceplvn5eb=ffffffff09e9012945525d5f4f58455e445a4a4229a0; sede=ffffffff09e9012245525d5f4f58455e445a4a4229a0; JSESSIONID=Ny2VOX2CrK9rbUSCh-f72i2l5nSVN4tLObFvoLuh.dzp-jboss1-01"
        }
        return requests.get(
            endpoint,
            verify=False,
            headers=headers,
            timeout=10*60,
        )

    def run(self, data):
        input = self._get_job_input(data)
        logger.info(f"Job triggered with input: {input}")
        response = self.make_request(input)
        minio = minio_client()
        partition = self._get_reference(input.reference)
        uri = minio.upload_bytes(self._get_bucket_name(), f"{partition}/{self.job_metadata.source}.zip", response.content)
        logger.info(f"File storage uri: {uri}")
        return {"documentUri": uri, "partition": partition}, self._get_status(response)
