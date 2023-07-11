import asyncio
from core.loaders.config_loader import read_job_configs
from pathlib import Path
from pylog.log import setup_logging
from pydotenv.dotenv import DotEnvLoader
from controller.controller import process_input
from pyrepository.interfaces.ingestors.dtos import Config
from pyrabbit.consumer import RabbitMQConsumer
from pysd.service_discovery import new_from_env
from scheduler.scheduler import Scheduler
from event.event import Event


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


async def main():
    logger.info("Starting File Downloader Service")
    sd = new_from_env()
    configs = get_job_configs_by_context_from_env()
    loop = asyncio.get_event_loop()
    tasks = []

    for _context, context_configs in configs.items():
        for config in context_configs:
            if config.jobType == "event" and config.active:
                logger.info(f"Job Config: {config}")
                queue_name = f"{_context}.{config.id}"
                rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
                await rabbitmq_service.connect()
                await rabbitmq_service.create_queue(queue_name)
                task = asyncio.create_task(Event.consume_queue(config, rabbitmq_service, queue_name))
                tasks.append(task)
            elif config.jobType == "scheduler" and config.active and config.shouldRun:
                scheduler = Scheduler(config)
                tasks.append(asyncio.create_task(scheduler.run()))

    await asyncio.gather(*tasks)  # Wait for all tasks to complete

    for task in tasks:
        task.cancel()

    await asyncio.sleep(0.1)
    await rabbitmq_service.close_connection()


if __name__ == "__main__":
    asyncio.run(main())
