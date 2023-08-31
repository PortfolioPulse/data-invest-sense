import asyncio
import json
from pyrepository.interfaces.ingestors.dtos import Config, MessageParameters, MessageMetadata
from pylog.log import setup_logging
from functools import wraps

logger = setup_logging(__name__)


def validate_event_data(func):
    @wraps(func)
    async def wrapper(self, job_config: Config, message_body: str):
        message_params = EventHandler.parse_message_body(message_body)
        EventHandler.validate_message_params(message_params)
        await func(self, job_config, message_body)
    return wrapper


class EventHandler:
    def __init__(self, job_config: Config, message_body: str, aio_queue: asyncio.Queue):
        self.__aio_queue = aio_queue
        self.__job_config = job_config
        self.__message_body = message_body

    async def process_event(self):
        await self._process_event(self.__job_config, self.__message_body)

    @validate_event_data
    async def _process_event(self, job_config: Config, message_body: str):
        logger.info(job_config)
        message_params = EventHandler.parse_message_body(message_body)

        await self.__aio_queue.put(message_params)

    @staticmethod
    def parse_message_body(message_body: str) -> MessageParameters:
        try:
            parsed_body = json.loads(message_body)
        except json.JSONDecodeError as e:
            logger.error(f"Failed to parse message body: {e}")
            raise ValueError("Invalid message body")

        message_id = parsed_body.get("id", "")
        message_data = parsed_body.get("data", {})
        message_metadata = parsed_body.get("metadata", {})

        logger.info("Message body parsed successfully")
        return MessageParameters(
            id=message_id,
            data=message_data,
            metadata=MessageMetadata(
                processing_id=message_metadata.get("processing_id", ""),
                processing_timestamp=message_metadata.get("processing_timestamp", ""),
                source=message_metadata.get("source", ""),
                service=message_metadata.get("service", ""),
            )
        )

    @staticmethod
    def validate_message_params(message_params: MessageParameters):
        if not message_params.data:
            logger.error("Invalid message: Missing input")
            raise ValueError("Invalid message: Missing input")

        if not message_params.metadata:
            logger.error("Invalid message: Missing metadata")
            raise ValueError("Invalid message: Missing metadata")

        if not message_params.id:
            logger.error("Invalid message: Missing id")
            raise ValueError("Invalid message: Missing id")

        logger.info("Message parameters validated successfully")
