import asyncio
from dataclasses import dataclass
from typing import Dict, Any
import json
import importlib
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config, MessageParameters
from controller.events.event import post_event
from jobs.jobs import trigger_job
from controller.events.listeners.mongo_listener import set_mongo_event_handlers

logger = setup_logging(__name__)

def set_event_handlers():
    set_mongo_event_handlers()


class Controller:
    def __init__(self, config, aio_queue: asyncio):
        self.__config = config
        self.__aio_queue = aio_queue

    async def consume(self):
        while True:
            message = await self.__aio_queue.get()
            logger.info(f"Received message from queue '{self.__config.jobMetadataParams.context}.{self.__config.jobMetadataParams._id}': {message}")
            set_event_handlers()
            check_config = post_event("check_config", self.__config)
            logger.info(f"check_config: {list(check_config)}")
            trigger_job(message)

# {"input": {"referencia": "2023-07-21"}, "metadata": {"processingId": "123"}}
#
