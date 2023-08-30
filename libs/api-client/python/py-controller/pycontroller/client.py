from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env

class AsyncPyControllerClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def create_config(self, data):
        endpoint = "/configs"
        return await self.client.make_request("POST", endpoint, data=data)

    async def list_all_configs(self):
        endpoint = "/configs"
        return await self.client.make_request("GET", endpoint)

    async def list_one_config_by_id(self, config_id):
        endpoint = f"/configs/{config_id}"
        return await self.client.make_request("GET", endpoint)

    async def list_all_configs_by_service(self, service_name):
        endpoint = f"/configs/service/{service_name}"
        return await self.client.make_request("GET", endpoint)

def async_pycontroller_client():
    sd = new_from_env()
    return AsyncPyControllerClient(sd.lake_controller_endpoint())


