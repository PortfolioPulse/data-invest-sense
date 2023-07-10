import os
import json
from pathlib import Path
from pylog.log import setup_logging
from typing import Dict, Type, TypeVar

logger = setup_logging(__name__)

T = TypeVar('T')


def read_job_configs(path: str, config_type: str, job_configs: Dict[str, list] = None, data_class: Type[T] = None) -> Dict[str, list[T]]:
    if job_configs is None:
        job_configs = {}
    directory_path = Path(path)
    for filename in os.listdir(path):
        file_path = directory_path.joinpath(filename)
        if file_path.is_dir():
            read_job_configs(file_path, config_type, job_configs, data_class)  # <- Fix: Use file_path as the new path
        elif file_path.name == '{}.json'.format(config_type):
            with open(file_path) as f:
                job_config = json.load(f)
                context = file_path.parent.parent.name
                if context not in job_configs:
                    job_configs[context] = []
                if data_class is not None:
                    job_config_obj = data_class(**job_config)
                    job_configs[context].append(job_config_obj)
                else:
                    job_configs[context].append(job_config)
    return job_configs
