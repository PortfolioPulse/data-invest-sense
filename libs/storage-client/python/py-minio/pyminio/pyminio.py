from minio import Minio
from minio.error import S3Error
from pysd.service_discovery import new_from_env
from pylog.log import setup_logging
from io import BytesIO


logger = setup_logging(__name__)


class MinioClient:
    def __init__(self, endpoint, access_key, secret_key, secure=False):
        self._endpoint = endpoint
        self.client = Minio(
            endpoint,
            access_key=access_key,
            secret_key=secret_key,
            secure=secure,
        )

    def create_bucket(self, bucket_name):
        try:
            self.client.make_bucket(bucket_name)
        except S3Error as err:
            raise Exception(f"Error creating bucket: {err}")
        logger.info(f"Bucket {bucket_name} created successfully")

    def list_buckets(self):
        try:
            return self.client.list_buckets()
        except S3Error as err:
            raise Exception(f"Error listing buckets: {err}")

    def _get_uri(self, bucket_name, object_name):
        return f"http://{self._endpoint}/{bucket_name}/{object_name}"

    def upload_file(self, bucket_name, object_name, file_path):
        try:
            self.client.fput_object(bucket_name, object_name, file_path)
        except S3Error as err:
            raise Exception(f"Error uploading file: {err}")
        return self._get_uri(bucket_name, object_name)

    def upload_bytes(self, bucket_name, object_name, bytes_data):
        try:
            data_stream = BytesIO(bytes_data)
            data_size = len(bytes_data)
            self.client.put_object(bucket_name, object_name, data_stream, data_size)
        except S3Error as err:
            raise Exception(f"Error uploading bytes: {err}")
        return self._get_uri(bucket_name, object_name)

    def download_file(self, bucket_name, object_name, file_path):
        try:
            self.client.fget_object(bucket_name, object_name, file_path)
        except S3Error as err:
            raise Exception(f"Error downloading file: {err}")

    def list_objects(self, bucket_name):
        try:
            return self.client.list_objects(bucket_name)
        except S3Error as err:
            raise Exception(f"Error listing objects: {err}")


def minio_client():
    sd = new_from_env()
    return MinioClient(
        endpoint=sd.minio_endpoint(),
        access_key=sd.minio_access_key(),
        secret_key=sd.minio_secret_key(),
        secure=False,
    )
