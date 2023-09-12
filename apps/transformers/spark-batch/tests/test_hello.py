"""Hello unit test module."""

from spark_batch.hello import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello transformers-spark-batch"
