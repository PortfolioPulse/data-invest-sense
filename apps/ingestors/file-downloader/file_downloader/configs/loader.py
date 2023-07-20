from dataclasses import dataclass, field
import os
import json
from pathlib import Path
from pylog.log import setup_logging
from typing import Dict, Type, TypeVar

import trio

logger = setup_logging(__name__)


@dataclass
class JobMetadataParams:
    _id: str
    name: str
    context: str
    rootPath: Path


@dataclass
class JobParams:
    jobType: str
    active: bool


@dataclass
class Config:
    jobMetadataParams: JobMetadataParams
    jobParams: JobParams


@dataclass
class LoadConfig:
    path: Path
    configRaw: dict = field(init=False)

    async def __aenter__(self):
        self.file = await trio.open_file(self.path, 'r')
        return self

    async def __aexit__(self, exc_type, exc_value, traceback):
        await self.file.aclose()

    async def _read_config(self):
        async with self.file:
            self.configRaw = json.loads(await self.file.read())

    def _set_job_params(self):
        return JobParams(
            jobType=self.configRaw["jobType"],
            active=self.configRaw["active"],
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


def find_config_files(config_type: str):
    job_config_files = list(
        Path(__file__).parent.joinpath(
            os.getenv("CONTEXT_SERVICE", "")
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
    async with LoadConfig(config_path) as config_loader:
        config = await config_loader.export_config()
        if config.jobParams.active:
            register_config(
                config.jobMetadataParams.context,
                config.jobMetadataParams._id,
                config
            )


async def read_config_async():
    async with trio.open_nursery() as nursery:
        for _config_path in find_config_files('job-config.json'):
            nursery.start_soon(_load_config, _config_path)
    return mapping_config


