import asyncio
import json
import time

from controller.events.event import EventObserver
from controller.events.listeners.kafka_listener import KafkaEventListener
from controller.events.listeners.mongo_listener import MongoDBEventListener
from jobs.jobs import JobHandler
from pylog.log import setup_logging
from pyrabbit.consumer import RabbitMQConsumer
from pyrepository.interfaces.ingestors.dtos import (
    Config, MessageMetadataOriginsOutput, MessageMetadataOutput, MetadataOutput, MessageStatus
)

logger = setup_logging(__name__)

class CustomJSONEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, (MessageMetadataOutput, MessageMetadataOriginsOutput, MetadataOutput, MessageStatus)):
            return obj.__dict__
        return super().default(obj)

class Controller:
    def __init__(self, config: Config, aio_queue: asyncio, rabbitmq_service: RabbitMQConsumer):
        self.__config = config
        self.__aio_queue = aio_queue
        self.__rabbitmq_service = rabbitmq_service

    async def consume(self):
        while True:
            message = await self.__aio_queue.get()
            logger.info(f"Received message from queue '{self.__config.jobMetadataParams.context}.{self.__config.jobMetadataParams._id}': {message}")
            observer = EventObserver(self.__config)
            MongoDBEventListener(observer)
            # KafkaEventListener(observer)
            await observer.post_event("check_config", self.__config)
            for listener_handler, result in observer.results.items():
                logger.info(f"listener_handler: {listener_handler}")
            job_data = JobHandler(self.__config).run(message)
            time.sleep(5)
            await self.__rabbitmq_service.publish_message(
                "services",
                "feedback",
                json.dumps(job_data, cls=CustomJSONEncoder)
            )
            logger.info("Published message to service")


