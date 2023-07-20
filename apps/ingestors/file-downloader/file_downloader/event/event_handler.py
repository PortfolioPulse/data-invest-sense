import asyncio
import json
from pyrepository.interfaces.ingestors.dtos import Config, MessageParameters
from pylog.log import setup_logging
from functools import wraps
# from jobs.jobs import trigger_job

logger = setup_logging(__name__)


def validate_event_data(func):
    @wraps(func)
    async def wrapper(self, job_config: Config, message_body: str):
        message_params = EventHandler.parse_message_body(message_body)
        EventHandler.validate_message_params(message_params)
        await func(self, job_config, message_body)
    return wrapper


class EventHandler:
    def __init__(self, job_config: Config, message_body: str, queue: asyncio.Queue):
        self.__input_queue = queue
        self.__job_config = job_config
        self.__message_body = message_body

    async def process_event(self):
        await self._process_event(self.__job_config, self.__message_body)

    @validate_event_data
    async def _process_event(self, job_config: Config, message_body: str):
        logger.info(job_config)
        message_params = EventHandler.parse_message_body(message_body)
        # module_name = f"event_stream.jobs.{job_config.jobType}.job_handler"
        # logger.info(module_name)
        # module_name = f"file_downloader.jobs.{job_config.jobType}.job_handler"
        # logger.info(module_name)
        # try:
        #     job_module = importlib.import_module(module_name)
        # except ModuleNotFoundError:
        #     logger.error(f"Module '{module_name}' not found")
        #     return

        # Trigger the job
        # trigger_job(message_params.input)

        await self.__input_queue.put(message_params)

    @staticmethod
    def parse_message_body(message_body: str) -> MessageParameters:
        try:
            parsed_body = json.loads(message_body)
        except json.JSONDecodeError as e:
            logger.error(f"Failed to parse message body: {e}")
            raise ValueError("Invalid message body")

        message_input = parsed_body.get("input", {})
        message_metadata = parsed_body.get("metadata", {})

        logger.info("Message body parsed successfully")
        return MessageParameters(
            input=message_input,
            metadata=message_metadata
        )

    @staticmethod
    def validate_message_params(message_params: MessageParameters):
        if not message_params.input:
            logger.error("Invalid message: Missing input")
            raise ValueError("Invalid message: Missing input")

        if not message_params.metadata:
            logger.error("Invalid message: Missing metadata")
            raise ValueError("Invalid message: Missing metadata")
        # Additional validation rules
        # ...

        logger.info("Message parameters validated successfully")
