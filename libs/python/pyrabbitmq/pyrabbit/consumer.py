"""RabbitMQ Consumer Module."""
from pylog.log import setup_logging
import aio_pika
from pyrabbit.base import BaseRabbitMQ

logger = setup_logging(__name__)


class RabbitMQConsumer(BaseRabbitMQ):
    def __init__(self, url):
        super().__init__(url)

    async def listen(self, queue, callback):
        async with queue.iterator() as queue_iter:
            message: aio_pika.AbstractIncomingMessage
            async for message in queue_iter:
                await callback(message)
