# get meta data
from pylog.log import setup_logging
import time

logger = setup_logging(__name__)


def trigger_job(data):
    # Perform the job operations using the data
    logger.info("sleeping for 30 seconds")
    time.sleep(30)
    logger.info(f"Job triggered with data: {data}")
