import asyncio
from dataclasses import dataclass
from typing import Dict, Any
import json
import importlib
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config, MessageParameters
from jobs.jobs import trigger_job

logger = setup_logging(__name__)

class Controller:
    def __init__(self, queue):
        self.queue = queue

    async def consume(self):
        # Consumes items from the queue
        while True:
            item = await self.queue.get()
            print(f"Consumed: {item}")
            trigger_job(item)
            await asyncio.sleep(0.5)
            self.queue.task_done()
