import asyncio
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config
from controller.controller import process_input


logger = setup_logging(__name__)


class Event:
    @staticmethod
    async def consume_queue(job_config, rabbitmq_service, queue_name):
        async def callback(message):
            logger.info(f"Received message from queue '{queue_name}': {message.body.decode()}")
            message_body = message.body.decode()
            await process_input(job_config, message_body)
            await message.ack()

        async with rabbitmq_service.connection:
            channel = await rabbitmq_service.connection.channel()
            await channel.set_qos(prefetch_count=1)
            queue = await channel.declare_queue(queue_name)
            await queue.consume(callback)

            while True:
                await asyncio.sleep(0.1)
