from typing import Dict
from pylog.log import setup_logging
from pyrepository.interfaces.ingestors.dtos import Config, JobParams, JobMetadataParams
from pycontroller.client import async_pycontroller_client

logger = setup_logging(__name__)


class SetConfigParams:

    def set_job_params(self, configRaw):
        return JobParams(
            jobHandler=configRaw["service_parameters"]["job_handler"],
            url=configRaw["job_parameters"]["url"],
        )

    def set_job_metadata_params(self, configRaw):
        return JobMetadataParams(
            _id=configRaw["id"],
            name=configRaw["name"],
            active=configRaw["active"],
            frequency=configRaw["frequency"],
            context=configRaw["context"],
            source=configRaw["source"],
            service=configRaw["service"],
        )


class ConfigLoader(SetConfigParams):
    def __init__(self) -> None:
        super().__init__()

    async def fetch_configs_for_service(self, service_name: str):
        client = async_pycontroller_client()
        configs = await client.list_all_configs_by_service(service_name)
        for config in configs:
            registered_config = Config(
                jobMetadataParams=self.set_job_metadata_params(config),
                jobParams=self.set_job_params(config),
            )
            register_config(
                registered_config.jobMetadataParams.context,
                registered_config.jobMetadataParams._id,
                registered_config
            )


mapping_config: Dict[str, Dict[str, Config]] = dict()


def register_config(context: str, config_id: str, config: Config):
    if context not in mapping_config:
        mapping_config[context] = dict()
    if config_id in mapping_config[context]:
        # TODO: Create an exception class
        logger.info(f"Warning: Duplicate config ID '{config_id}' for context '{context}'. Overwriting existing config.")
    mapping_config[context][config_id] = config


async def fetch_configs():
    service = "file-downloader"
    await ConfigLoader().fetch_configs_for_service(service_name=service)
    return mapping_config
