import os
from dataclasses import dataclass


class UnrecoverableError(Exception):
    pass


class ServiceUnavailableError(Exception):
    pass


@dataclass
class ServiceVars:
    rabbitmq: str = "RABBITMQ"
    lakeController: str = "LAKE_CONTROLLER"
    minio: str = "MINIO"


class ServiceDiscovery:
    def __init__(self, envvars):
        if envvars is None:
            raise UnrecoverableError('Environment variables not set')
        self._vars = envvars
        self._service_vars = ServiceVars()

    def _get_endpoint(self, var_name, service_name, protocol="http"):
        if var_name not in self._vars:
            raise ServiceUnavailableError(f'Environment variable {var_name} not set')
        tcp_addr = self._vars[var_name]
        gt_host = self._get_gateway_host(service_name)
        return tcp_addr.replace("tcp", protocol).replace("gateway_host", gt_host)

    def _get_gateway_host(self, service_name):
        if os.getenv('GATEWAY_ENVIRONMENT') is None:
            return 'localhost'
        return os.getenv(f'{service_name}_GATEWAY_HOST')

    def rabbitmq_endpoint(self):
        service_name = self._service_vars.rabbitmq
        return self._get_endpoint("RABBITMQ_PORT_6572_TCP", service_name, protocol="amqp")

    def lake_controller_endpoint(self):
        service_name = self._service_vars.lakeController
        endpoint = self._get_endpoint("LAKE_CONTROLLER_PORT_8000_TCP", service_name)
        if "localhost" in endpoint:
            endpoint = endpoint.replace("8000", "8002")
        return endpoint

    def minio_endpoint(self):
        service_name = self._service_vars.minio
        endpoint = self._get_endpoint("MINIO_PORT_9000_TCP", service_name)
        return endpoint

    def minio_access_key(self):
        return os.getenv("MINIO_ACCESS_KEY")

    def minio_secret_key(self):
        return os.getenv("MINIO_SECRET_KEY")

def new_from_env():
    return ServiceDiscovery(os.environ)
