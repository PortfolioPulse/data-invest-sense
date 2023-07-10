"""Base RabbitMQ Module."""
import aio_pika
import urllib.parse
from pylog.log import setup_logging

logger = setup_logging(__name__)


class BaseRabbitMQ:
    def __init__(self, url):
        self.url = url
        self.connection = None
        self.channel = None

    async def connect(self):
        parsed_url = urllib.parse.urlparse(self.url)
        self.connection = await aio_pika.connect_robust(
            host=parsed_url.hostname,
            port=parsed_url.port,
            login=parsed_url.username,
            password=parsed_url.password,
        )
        self.channel = await self.connection.channel()

    def on_connection_error(self, unused_connection, error):
        logger.error(f"Connection error: {error}")
        self.connection.close()

    async def create_queue(self, queue_name):
        queue = await self.channel.declare_queue(queue_name)

    async def close_connection(self):
        await self.channel.close()
        self.connection.close()

