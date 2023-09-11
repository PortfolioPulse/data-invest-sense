import asyncio
from pylog.log import setup_logging
from pyrabbit.consumer import RabbitMQConsumer
from pysd.service_discovery import new_from_env
from consumer.event import Event
from controller.controller import Controller
from configs.loader import fetch_configs

logger = setup_logging(__name__, log_level="DEBUG")


async def consume_messages(sd, configs):
    rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
    await rabbitmq_service.connect()
    tasks = list()

    for _context, context_configs in configs.items():
        for config_name, config in context_configs.items():
            queue_name = f"{_context}.{config.jobMetadataParams.service}.inputs.{config.jobMetadataParams.source}"
            exchange_name = "services"
            routing_key = f"{config.jobMetadataParams.service}.inputs.{config.jobMetadataParams.source}"

            aio_queue = asyncio.Queue()
            tasks.append(
                asyncio.create_task(Event.consume_queue(config, rabbitmq_service, exchange_name, queue_name, routing_key, aio_queue))
            )
            tasks.append(
                asyncio.create_task(Controller(config, aio_queue, rabbitmq_service).consume())
            )
    await asyncio.gather(*tasks)
    for task in tasks:
        task.cancel()
    await rabbitmq_service.close_connection()

async def main():
    logger.info("Starting File Downloader Service")
    sd = new_from_env()
    configs = await fetch_configs()

    await consume_messages(sd, configs)



if __name__ == "__main__":
    asyncio.run(main())
