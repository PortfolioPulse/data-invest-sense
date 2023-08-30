# FIXME: This is a temporary solution. We need to find a better way to load configs.
import json
from dataclasses import dataclass, field
from pydotenv.dotenv import DotEnvLoader
from pathlib import Path
from typing import Dict, Any
import asyncio
from pylog.log import setup_logging
from pycontroller.client import async_pycontroller_client

logger = setup_logging(__name__)

@dataclass
class JobMetadataParams:
    _id: str
    name: str
    context: str
    source: str
    service: str

@dataclass
class JobParams:
    jobHandler: str
    active: bool
    url: str

@dataclass
class Config:
    jobMetadataParams: JobMetadataParams
    jobParams: JobParams


class SetConfigParams:

    def set_job_params(self, configRaw):
        return JobParams(
            jobHandler=configRaw["serviceParameters"]["jobHandler"],
            active=configRaw["active"],
            url=configRaw["jobParameters"]["url"],
        )

    def set_job_metadata_params(self, configRaw):
        return JobMetadataParams(
            _id=configRaw["id"],
            name=configRaw["name"],
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
