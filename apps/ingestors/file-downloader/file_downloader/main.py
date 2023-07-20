import asyncio
import trio
from core.loaders.config_loader import read_job_configs
from pathlib import Path
from pylog.log import setup_logging
from pydotenv.dotenv import DotEnvLoader
# from controller.controller import process_input
from decorators.concurrency import run_with_concurrency
from pyrepository.interfaces.ingestors.dtos import Config
from pyrabbit.consumer import RabbitMQConsumer
from pysd.service_discovery import new_from_env
# from scheduler.scheduler import Scheduler
from event.event import Event
from functools import wraps
from controller.controller import Controller

from configs.loader import read_config_async
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


# @run_with_concurrency(5)
async def main():
    logger.info("Starting File Downloader Service")
    sd = new_from_env()
    configs = await read_config_async()
    # logger.info(f"Config: {configs}")
    async with trio.open_nursery() as nursery:
      for _context, context_configs in configs.items():
          for config in context_configs:
              logger.info(f"Job Config: {config}")
              queue_name = f"{_context}.{config.id}"
              rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
              await rabbitmq_service.connect()
              await rabbitmq_service.create_queue(queue_name)






    # configs = get_job_configs_by_context_from_env()
    # loop = asyncio.get_event_loop()
    # tasks = []

    # for _context, context_configs in configs.items():
    #     for config in context_configs:
    #         if config.jobType == "event" and config.active:
    #             logger.info(f"Job Config: {config}")
    #             queue_name = f"{_context}.{config.id}"
    #             rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
    #             await rabbitmq_service.connect()
    #             await rabbitmq_service.create_queue(queue_name)
    #             aio_queue = asyncio.Queue()
    #             task = asyncio.create_task( )
    #             tasks.append(task)
    #             task = asyncio.create_task(Controller(aio_queue).consume())
    # await asyncio.gather(*tasks)  # Wait for all tasks to complete

    # for task in tasks:
    #     task.cancel()

    # await asyncio.sleep(0.1)
    # await rabbitmq_service.close_connection()


if __name__ == "__main__":
    trio.run(main)
