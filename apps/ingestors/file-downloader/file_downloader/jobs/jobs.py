import importlib
from datetime import datetime

from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import (
    Config, MessageMetadataOriginsOutput, MessageMetadataOutput,
    MetadataOutput)

logger = setup_logging(__name__)


class JobHandler:
    def __init__(self, config: Config):
        self.config = config
        self._import_job_as_module()

    def _import_job_as_module(self):
        self.module = importlib.import_module(f"jobs.handlers.{self.config.jobParams.jobHandler}.job")

    def run(self, data):
        logger.info(f"[RUNNING JOB] - Handler: {self.config.jobParams.jobHandler}")
        job_data, job_status = self.module.Job(self.config.jobParams, self.config.jobMetadataParams).run(data)
        return {
            "data": job_data,
            "metadata": self._get_metadata(data),
            "status": job_status,
        }

    def _get_metadata(self, data):
        return MetadataOutput(
            input=MessageMetadataOutput(
                id=data.id,
                data=data.data,
                processing_id=data.metadata.processing_id,
                processing_timestamp=data.metadata.processing_timestamp,
                source=MessageMetadataOriginsOutput(
                    gateway=data.metadata.source,
                    controller=self.config.jobMetadataParams.source,
                )
            ),
            service=MessageMetadataOriginsOutput(
                gateway=data.metadata.service,
                controller=self.config.jobMetadataParams.service,
            ),
            processing_id=data.metadata.processing_id,
            processing_timestamp=datetime.now().strftime("%Y-%m-%dT%H:%M:%SZ"),
            target_endpoint=self.config.jobParams.url,
            job_frequency=self.config.jobMetadataParams.frequency,
        )
