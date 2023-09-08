import asyncio
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config
from consumer.event_handler import EventHandler


logger = setup_logging(__name__)

class Event:
    @staticmethod
    async def consume_queue(job_config, rabbitmq_service, exchange_name, queue_name, routing_key, aio_queue):
        async def callback(message):
            message_body = message.body.decode()
            try:
                await EventHandler(job_config, message_body, aio_queue).process_event()
                logger.info(f"Processed message from queue '{queue_name}': {message.body.decode()}")
                await message.ack()
                logger.info(f"Acknowledged message from queue '{queue_name}'")
            except Exception as e:
                logger.error(f"Error processing message from queue '{queue_name}': {str(e)}")
                await message.reject(requeue=True)
                logger.warning(f"Rejected message from queue '{queue_name}', requeued.")

        channel = await rabbitmq_service.create_channel()
        queue = await rabbitmq_service.create_queue(channel, queue_name, exchange_name, routing_key)
        await rabbitmq_service.listen(queue, callback)

