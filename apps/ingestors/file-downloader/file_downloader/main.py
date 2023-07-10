"""Main Entrypoint File Downloader Service."""
# from config.config import read_job_configs
from core.loaders.config_loader import read_job_configs
from pathlib import Path
import os
import asyncio
from pylog.log import setup_logging
from pydotenv.dotenv import DotEnvLoader
from controller.controller import process_message
from pyrepository.interfaces.ingestors.dtos import Config
from pyrabbit.consumer import RabbitMQConsumer
from pysd.service_discovery import new_from_env


logger = setup_logging(__name__)


def get_job_configs_by_context_from_env(env: DotEnvLoader = DotEnvLoader()):
    return read_job_configs(
        path=Path(__file__).parent.joinpath(
            "configs/{}".format(
                env.get_variable('CONTEXT_SERVICE')
            )
        ),
        config_type="job-config",
        data_class=Config
    )

async def consume_queue(job_config, rabbitmq_service, queue_name):
    async def callback(message):
        # Process the message here
        logger.info(f"Received message from queue '{queue_name}': {message.body.decode()}")
        message_body = message.body.decode()
        await process_message(job_config, message_body)

        await message.ack()  # Acknowledge the message

    async with rabbitmq_service.connection:
        channel = await rabbitmq_service.connection.channel()
        await channel.set_qos(prefetch_count=1)  # Limit unacknowledged messages to 1
        queue = await channel.declare_queue(queue_name)
        await queue.consume(callback)

        while True:
            await asyncio.sleep(0.1)


async def main():
    logger.info("Starting File Downloader Service")
    sd = new_from_env()
    configs = get_job_configs_by_context_from_env()
    loop = asyncio.get_event_loop()
    tasks = []
    for _context, context_configs in configs.items():
        for config in context_configs:
            logger.info(f"Job Config: {config}")
            # TODO: Create GET queue name method
            queue_name = "{}.{}".format(_context, config.id)
            # RABBITMQ Connection
            rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
            while True:
                try:
                    await rabbitmq_service.connect()
                    break
                except Exception as err:
                    logger.error('[CONNECTION] - Could not connect to RabbitMQ, retrying in 2 seconds...')
                    await asyncio.sleep(2)
            await rabbitmq_service.connect()
            await rabbitmq_service.create_queue(queue_name)

            task = asyncio.create_task(consume_queue(config, rabbitmq_service, queue_name))
            tasks.append(task)

    await asyncio.gather(*tasks)  # Wait for all tasks to complete

    for task in tasks:
        task.cancel()

    await asyncio.sleep(0.1)
    await rabbitmq_service.close_connection()

if __name__ == "__main__":
    # logger.info(f"OS env RABBITMQ_HOST: {os.getenv('RABBITMQ_HOST')}")
    # logger.info(f"Service Discovery: {sd.rabbitmq_endpoint()}")
    asyncio.run(main())

