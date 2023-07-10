import unittest
from dataclasses import dataclass
from pathlib import Path
from core.loaders.config_loader import read_job_configs


@dataclass
class JobConfig:
    field1: str
    field2: int


class TestReadJobConfigs(unittest.TestCase):

    def setUp(self):
        self._path = Path(__file__).parent.parent

    def test_read_job_configs_with_data_class(self):
        expected_result = {
            'br': [JobConfig(field1='value1', field2=123)],
            'us': [JobConfig(field1='value1', field2=123)]
        }
        result = read_job_configs(
            path=self._path.joinpath("reference_files/configs"),
            config_type="job-config",
            data_class=JobConfig
        )
        self.assertDictEqual(result, expected_result)


if __name__ == '__main__':
    unittest.main()
