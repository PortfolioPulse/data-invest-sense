import unittest
import io
from unittest.mock import Mock
from pyminio.pyminio import MinioClient
from minio.error import S3Error


class CustomS3Error(S3Error):
    def __init__(self, message, resource, request_id, host_id, response, code=None):
        super().__init__(message=message, resource=resource, request_id=request_id, host_id=host_id, response=response, code=code)


class TestMinio(unittest.TestCase):
    def setUp(self):
        self.client = MinioClient(
            endpoint="localhost:9000",
            access_key="minio-access-key",
            secret_key="minio-secret-key",
        )

    def tearDown(self):
        pass

    def test_create_bucket(self):
        self.client.client.make_bucket = Mock()

        bucket_name = "test-bucket"
        self.client.create_bucket(bucket_name)
        self.client.client.make_bucket.assert_called_with(bucket_name)

    def test_create_bucket_error(self):
        class MockS3Error(Exception):
            pass

        def mock_make_bucket(bucket_name):
            raise MockS3Error("Bucket already exists")

        self.client.client.make_bucket = mock_make_bucket

        bucket_name = "existing-bucket"
        with self.assertRaises(Exception) as context:
            self.client.create_bucket(bucket_name)
        self.assertEqual(str(context.exception), "Bucket already exists")

    def test_list_buckets(self):
        self.client.client.list_buckets = Mock(return_value=["bucket1", "bucket2"])

        buckets = self.client.list_buckets()

        self.client.client.list_buckets.assert_called()
        self.assertEqual(buckets, ["bucket1", "bucket2"])

    def test_upload_file(self):
        self.client.client.fput_object = Mock()

        bucket_name = "test-bucket"
        object_name = "test-object"
        file_path = "./reference_files/test_file.txt"
        obtained_uri = self.client.upload_file(bucket_name, object_name, file_path)
        self.client.client.fput_object.assert_called_with(bucket_name, object_name, file_path)
        self.assertEqual(obtained_uri, "http://localhost:9000/test-bucket/test-object")

    def test_upload_bytes(self):
        self.client.client.put_object = Mock()

        bucket_name = "test-bucket"
        object_name = "test-object"
        bytes_data = b"sample bytes data"
        obtained_uri = self.client.upload_bytes(bucket_name, object_name, bytes_data)

        self.client.client.put_object.assert_called_with(bucket_name, object_name, bytes_data, len(bytes_data))
        self.assertEqual(obtained_uri, "http://localhost:9000/test-bucket/test-object")


    def test_download_file_content(self):
        # Mock the fget_object method to return sample content
        self.client.client.fget_object = Mock(return_value=b"sample content")

        bucket_name = "test-bucket"
        object_name = "test-object"
        file_path = "./response/downloaded.txt"
        self.client.download_file(bucket_name, object_name, file_path)

        self.client.client.fget_object.assert_called_with(bucket_name, object_name, file_path)

    def test_list_objects(self):
        # Mock the list_objects method to return sample object names
        self.client.client.list_objects = Mock(return_value=["object1", "object2"])

        bucket_name = "test-bucket"
        objects = self.client.list_objects(bucket_name)

        self.client.client.list_objects.assert_called_with(bucket_name)
        self.assertEqual(objects, ["object1", "object2"])

    def test_upload_nonexistent_file(self):
        bucket_name = "test-bucket"
        object_name = "nonexistent-object"
        file_path = "nonexistent-file.txt"
        with self.assertRaises(Exception) as context:
            self.client.upload_file(bucket_name, object_name, file_path)

    def test_download_nonexistent_file(self):
        bucket_name = "test-bucket"
        object_name = "nonexistent-object"
        file_path = "./reference_files/nonexistent.txt"

        with self.assertRaises(Exception) as context:
            self.client.download_file(bucket_name, object_name, file_path)
        self.assertTrue("Error downloading file" in str(context.exception))

    def test_large_file_upload(self):
        self.client.client.put_object = Mock()

        bucket_name = "test-bucket"
        object_name = "large-object"
        large_data = b"0" * (10 * 1024 * 1024)  # 10 MB

        self.client.upload_bytes(bucket_name, object_name, large_data)

        self.client.client.put_object.assert_called_with(bucket_name, object_name, large_data, len(large_data))

    def test_list_buckets_error(self):
        error = CustomS3Error("Error listing buckets", "resource", "request_id", "host_id", "response")
        self.client.client.list_buckets = Mock(side_effect=error)

        with self.assertRaises(Exception) as context:
            self.client.list_buckets()
        self.assertTrue("Error listing buckets" in str(context.exception))

    def test_download_file_error(self):
        error = CustomS3Error("Error downloading file", "resource", "request_id", "host_id", "response")
        self.client.client.fget_object = Mock(side_effect=error)

        bucket_name = "test-bucket"
        object_name = "test-object"
        file_path = "./response/downloaded.txt"
        with self.assertRaises(Exception) as context:
            self.client.download_file(bucket_name, object_name, file_path)
        self.assertTrue("Error downloading file" in str(context.exception))

    def test_upload_bytes_error(self):
        error = CustomS3Error("Upload error", "resource", "request_id", "host_id", "response")
        self.client.client.put_object = Mock(side_effect=error)

        bucket_name = "test-bucket"
        object_name = "test-object"
        bytes_data = b"sample bytes data"
        with self.assertRaises(Exception) as context:
            self.client.upload_bytes(bucket_name, object_name, bytes_data)
        self.assertTrue("Error uploading bytes" in str(context.exception))

    def test_list_objects_error(self):
        error = CustomS3Error("List objects error", "resource", "request_id", "host_id", "response")
        self.client.client.list_objects = Mock(side_effect=error)

        bucket_name = "test-bucket"
        with self.assertRaises(Exception) as context:
            self.client.list_objects(bucket_name)
        self.assertTrue("Error listing objects" in str(context.exception))

if __name__ == '__main__':
    unittest.main()


