from configs.loader import Config
from pylog.log import setup_logging
from controller.events.event import subscribe

logger = setup_logging(__name__)

def handle_job_should_start_event(job_config: Config):
    logger.info(f"Received job_should_start event: {job_config}")
    yield job_config


def set_mongo_event_handlers():
    subscribe("check_config", handle_job_should_start_event)
