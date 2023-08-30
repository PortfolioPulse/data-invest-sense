import asyncio
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config
from consumer.event_handler import EventHandler


logger = setup_logging(__name__)


class Event:
    @staticmethod
    async def consume_queue(job_config, rabbitmq_service, queue_name, aio_queue):
        async def callback(message):
            logger.info(f"Received message from queue '{queue_name}': {message.body.decode()}")
            message_body = message.body.decode()
            await EventHandler(job_config, message_body, aio_queue).process_event()

            await message.ack()

        async with rabbitmq_service.connection:
            channel = await rabbitmq_service.connection.channel()
            await channel.set_qos(prefetch_count=1)
            queue = await channel.declare_queue(queue_name, durable=True)
            await queue.consume(callback)
            while True:
                await asyncio.sleep(0.1)










class Event:
    @staticmethod
    async def consume_queue(job_config, rabbitmq_service, exchange_name, queue_name, routing_key, aio_queue):
        async def callback(message):
            logger.info(f"Received message from queue '{queue_name}': {message.body.decode()}")
            message_body = message.body.decode()
            await EventHandler(job_config, message_body, aio_queue).process_event()

            await message.ack()

        async with rabbitmq_service.connection:
            channel = await rabbitmq_service.connection.channel()
            await channel.set_qos(prefetch_count=1)

            # exchange_name = "your_exchange_name"  # Specify the exchange name you're using
            # routing_key = "your_routing_key"  # Specify the routing key for this queue

            await rabbitmq_service.create_queue(queue_name, exchange_name, routing_key)

            queue = await channel.declare_queue(queue_name, durable=True)
            await queue.consume(callback)

            while True:
                await asyncio.sleep(0.1)

