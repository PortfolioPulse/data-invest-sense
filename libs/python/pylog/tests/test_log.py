import logging
import unittest
from unittest.mock import patch

from pylog.log import setup_logging
from pythonjsonlogger import jsonlogger


class TestLogging(unittest.TestCase):

    @patch('logging.StreamHandler')
    @patch('pythonjsonlogger.jsonlogger.JsonFormatter')
    @patch('logging.getLogger')
    def test_setup_logging(self, mock_get_logger, mock_json_formatter, mock_stream_handler):
        # Mock the dependencies
        mock_logger = mock_get_logger.return_value
        mock_handler = mock_stream_handler.return_value

        # Call the function
        logger = setup_logging('my_module', propagate=True, log_level='DEBUG')

        # Assertions
        mock_get_logger.assert_called_with('my_module')
        mock_stream_handler.assert_called_once()
        mock_json_formatter.assert_called_with('%(levelname)s %(message)s ')
        mock_handler.setFormatter.assert_called_with(mock_json_formatter.return_value)
        mock_logger.addHandler.assert_called_with(mock_handler)
        mock_logger.setLevel.assert_called_with(logging.DEBUG)
        self.assertEqual(logger, mock_logger)
        self.assertEqual(logger.propagate, True)


if __name__ == '__main__':
    unittest.main()
