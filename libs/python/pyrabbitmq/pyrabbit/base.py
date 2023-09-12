import asyncio
import aio_pika
import urllib.parse
from pylog.log import setup_logging

logger = setup_logging(__name__)

class BaseRabbitMQ:
    def __init__(self, url):
        self.url = url
        self.connection = None
        self.exchange = None

    async def _connect(self):
        parsed_url = urllib.parse.urlparse(self.url)
        self.connection = await aio_pika.connect(
            host=parsed_url.hostname,
            port=parsed_url.port,
            login=parsed_url.username,
            password=parsed_url.password,
        )

    async def connect(self):
        while True:
            try:
                await self._connect()
                break
            except Exception as err:
                logger.error('[CONNECTION] - Could not connect to RabbitMQ, retrying in 2 seconds...')
                self.on_connection_error(err)
                await asyncio.sleep(2)

    def on_connection_error(self, error):
        logger.error(f"Connection error: {error}")
        logger.error(f"Connection parameters: {self.url}")

    async def create_channel(self):
        channel = await self.connection.channel()
        await channel.set_qos(prefetch_count=1)
        return channel

    async def declare_exchange(self, channel, exchange_name):
        self.exchange = await channel.declare_exchange(
            exchange_name, aio_pika.ExchangeType.TOPIC, durable=True
        )

    async def create_queue(self, channel, queue_name, exchange_name, routing_key):
        await self.declare_exchange(channel, exchange_name)
        queue = await channel.declare_queue(queue_name, durable=True)
        await queue.bind(self.exchange, routing_key)

        return queue

    async def close_connection(self):
        await self.connection.close()

    async def publish_message(self, exchange_name, routing_key, message):
        try:
            await self.exchange.publish(
                aio_pika.Message(
                    body=message.encode(),
                    delivery_mode=aio_pika.DeliveryMode.PERSISTENT,
                ),
                routing_key=routing_key,
            )
            logger.info(f"Published message to exchange '{exchange_name}' with routing key '{routing_key}'")
        except Exception as e:
            logger.error(f"Error while publishing message: {e}")
