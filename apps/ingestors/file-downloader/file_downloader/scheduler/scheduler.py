import asyncio
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config
from pydotenv.dotenv import DotEnvLoader
from pathlib import Path
from controller.controller import process_input
from core.loaders.config_loader import read_job_configs


logger = setup_logging(__name__)

def get_job_configs_by_context_from_env(env: DotEnvLoader = DotEnvLoader()):
    return read_job_configs(
        path=Path(__file__).parent.parent.joinpath(
            "configs/{}".format(
                env.get_variable('CONTEXT_SERVICE')
            )
        ),
        config_type="job-config",
        data_class=Config
    )


class Scheduler:
    def __init__(self, config: Config):
        self.config = config

    async def run(self):
        while self.config.shouldRun:
            logger.info(f"Running alternative method for config: {self.config}")
            await process_input(self.config, "")
            await asyncio.sleep(30)

        configs = get_job_configs_by_context_from_env()
        for _context, context_configs in configs.items():
            for config in context_configs:
                if config.jobType == "scheduler" and config.active and config.shouldRun:
                    scheduler = Scheduler(config)
                    await asyncio.create_task(scheduler.run())
