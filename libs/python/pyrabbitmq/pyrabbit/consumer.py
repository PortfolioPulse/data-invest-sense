"""RabbitMQ Consumer Module."""
from pylog.log import setup_logging
from pyrabbit.base import BaseRabbitMQ

logger = setup_logging(__name__)


class RabbitMQConsumer(BaseRabbitMQ):
    def __init__(self, url):
        super().__init__(url)

    async def consume_queue(self, queue_name, callback):
        await self.channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=True)
        await self.channel.start_consuming()
