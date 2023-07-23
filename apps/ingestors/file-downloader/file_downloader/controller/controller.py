import asyncio
from dataclasses import dataclass
from typing import Dict, Any
import json
import importlib
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config, MessageParameters
from jobs.jobs import trigger_job
from controller.events.listeners.mongo_listener import MongoDBEventListener
from controller.events.listeners.kafka_listener import KafkaEventListener
from controller.events.event import EventObserver

logger = setup_logging(__name__)


class Controller:
    def __init__(self, config, aio_queue: asyncio):
        self.__config = config
        self.__aio_queue = aio_queue

    async def consume(self):
        while True:
            message = await self.__aio_queue.get()
            logger.info(f"Received message from queue '{self.__config.jobMetadataParams.context}.{self.__config.jobMetadataParams._id}': {message}")
            observer = EventObserver(self.__config)
            MongoDBEventListener(self.observer)
            KafkaEventListener(self.observer)
            await observer.post_event("check_config", self.__config)
            for listener_handler, result in observer.results.items():
                logger.info(f"listener_handler: {listener_handler}")
                logger.info(f"result: {result}")
            trigger_job(message)



