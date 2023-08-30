import unittest
import asyncio
from pyrabbit.consumer import RabbitMQConsumer
from aio_pika.robust_connection import RobustConnection

class RabbitMQConsumerTests(unittest.TestCase):
    def setUp(self):
        self.consumer = RabbitMQConsumer()

    async def tearDown(self):
        if self.consumer.connection is not None:
            await self.consumer.close_connection()

    def run_async(self, coro):
        self.loop = asyncio.get_event_loop()
        return self.loop.run_until_complete(coro)

    def test_connection(self):
        async def connection_test():
            await self.consumer.connect()
            self.assertIsInstance(self.consumer.connection, RobustConnection)
        self.run_async(connection_test())

    def test_create_queue(self):
        async def create_queue_test():
            await self.consumer.connect()
            await self.consumer.create_queue("test_queue")
            # Verify that the queue is created by checking its properties or existence
            # For example:
            # self.assertIn("test_queue", [queue.name for queue in self.consumer.channel.queues])

        self.run_async(create_queue_test())

    def test_consume_queue(self):
        async def consume_queue_test():
            await self.consumer.connect()

            async def callback(message):
                # Implement your callback logic here
                print("Received message:", message.body)

            await self.consumer.create_queue("test_queue")
            await self.consumer.consume_queue("test_queue", callback)
            # Publish a message to the "test_queue" using a separate process or thread
            await asyncio.sleep(5)  # Wait for some time to receive messages

        self.run_async(consume_queue_test())

    def test_all(self):
        self.test_connection()
        self.test_create_queue()
        self.test_consume_queue()

if __name__ == '__main__':
    unittest.main()
