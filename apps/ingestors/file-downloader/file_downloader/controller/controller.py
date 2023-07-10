from dataclasses import dataclass
from typing import Dict, Any
import json
import importlib
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config, MessageParameters


logger = setup_logging(__name__)


class MessageProcessor:
    @staticmethod
    async def process_message(job_config: Config, message_body: str):
        # Parse the message body and validate the fields
        logger.info(job_config)
        message_params = MessageProcessor.parse_message_body(message_body)
        MessageProcessor.validate_message_params(message_params)

        # Perform additional processing or trigger a job based on the message data
        # ...
        # Import the module based on job_config.id
        module_name = f"event_stream.jobs.{job_config.jobType}.job_handler"
        logger.info(module_name)
        # try:
        #     job_module = importlib.import_module(module_name)
        # except ModuleNotFoundError:
        #     logger.error(f"Module '{module_name}' not found")
        #     return

        # # Trigger the job
        # job_module.trigger_job(message_params.data)

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


async def process_message(job_config, message_body: str):
    await MessageProcessor.process_message(job_config, message_body)
