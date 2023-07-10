import os
from dotenv import load_dotenv

class DotEnvLoader:
    def __init__(self, path='.env'):
        self.path = path
        self.load()

    def load(self):
        load_dotenv(self.path)

    def get_variable(self, key):
        return os.getenv(key, "")
