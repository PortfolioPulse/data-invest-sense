# get meta data
import importlib
from configs.loader import Config
from pylog.log import setup_logging
import time

logger = setup_logging(__name__)


class JobHandler:
    def __init__(self, config: Config):
        self.config = config
        self._import_job_as_module()

    def _import_job_as_module(self):
        self.module = importlib.import_module(f"jobs.handlers.{self.config.jobParams.jobHandler}.job")

    def run(self, data):
        logger.info(f"[RUNNING JOB] - Handler: {self.config.jobParams.jobHandler}")
        self.module.Job(self.config.jobParams).run(data)
