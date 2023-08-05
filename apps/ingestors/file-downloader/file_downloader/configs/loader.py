import json
from dataclasses import dataclass, field
from pydotenv.dotenv import DotEnvLoader
from pathlib import Path
from typing import Dict, Any
import asyncio
from pylog.log import setup_logging

logger = setup_logging(__name__)

@dataclass
class JobMetadataParams:
    _id: str
    name: str
    context: str
    rootPath: Path

@dataclass
class JobParams:
    jobHandler: str
    active: bool
    url: str

@dataclass
class Config:
    jobMetadataParams: JobMetadataParams
    jobParams: JobParams

@dataclass
class LoadConfig:
    path: Path
    configRaw: dict = field(init=False)

    async def _read_config(self):
        with open(self.path, 'r') as file:
            self.configRaw = json.load(file)

    def _set_job_params(self):
        return JobParams(
            jobHandler=self.configRaw["jobHandler"],
            active=self.configRaw["active"],
            url=self.configRaw["jobParams"]["url"],
        )

    def _set_job_metadata_params(self):
        return JobMetadataParams(
            _id=self.configRaw["id"],
            name=self.configRaw["name"],
            context=self.configRaw["context"],
            rootPath=self.path.parent,
        )

    async def export_config(self):
        await self._read_config()
        return Config(
            jobMetadataParams=self._set_job_metadata_params(),
            jobParams=self._set_job_params()
        )


def find_config_files(config_type: str, env: DotEnvLoader = DotEnvLoader()):
    job_config_files = list(
        Path(__file__).parent.joinpath(
            env.get_variable('CONTEXT_SERVICE')
        ).rglob(config_type))
    return job_config_files

mapping_config: Dict[str, Dict[str, Config]] = dict()

def register_config(context: str, config_id: str, config: Config):
    if context not in mapping_config:
        mapping_config[context] = dict()
    if config_id in mapping_config[context]:
        # TODO: Create an exception class
        logger.info(f"Warning: Duplicate config ID '{config_id}' for context '{context}'. Overwriting existing config.")
    mapping_config[context][config_id] = config

async def _load_config(config_path):
    config_loader = LoadConfig(config_path)
    config = await config_loader.export_config()
    if config.jobParams.active:
        register_config(
            config.jobMetadataParams.context,
            config.jobMetadataParams._id,
            config
        )

async def read_config_async():
    config_files = find_config_files('job-config.json')
    await asyncio.gather(*[_load_config(config_path) for config_path in config_files])
    return mapping_config
