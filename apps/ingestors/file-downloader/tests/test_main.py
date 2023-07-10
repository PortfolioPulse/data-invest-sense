from file_downloader.main import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello ingestors-file-downloader"
