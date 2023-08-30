import asyncio
import aio_pika
import urllib.parse
from pylog.log import setup_logging

logger = setup_logging(__name__)

class BaseRabbitMQ:
    def __init__(self, url):
        self.url = url
        self.connection = None
        self.channel = None

    async def _connect(self):
        parsed_url = urllib.parse.urlparse(self.url)
        self.connection = await aio_pika.connect_robust(
            host=parsed_url.hostname,
            port=parsed_url.port,
            login=parsed_url.username,
            password=parsed_url.password,
        )
        self.channel = await self.connection.channel()

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

    async def declare_exchange(self, exchange_name):
        return await self.channel.declare_exchange(
            exchange_name, aio_pika.ExchangeType.DIRECT, durable=True
        )

    async def create_queue(self, queue_name, exchange_name, routing_key):
        await self.channel.set_qos(prefetch_count=1)

        exchange = await self.declare_exchange(exchange_name)

        queue = await self.channel.declare_queue(queue_name, durable=True)
        await queue.bind(exchange, routing_key)

        return queue

    async def close_connection(self):
        await self.channel.close()
        await self.connection.close()
