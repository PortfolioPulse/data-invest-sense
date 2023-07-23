from collections import defaultdict
import asyncio
from configs.loader import Config

class EventObserver:
    def __init__(self, config: Config):
        self.config = config
        self.subscribers = defaultdict(lambda: defaultdict(list))
        self.results = dict()

    def add_subscribe(self, event_type: str, fn):
        if event_type not in self.subscribers:
            self.subscribers[event_type] = []
        self.subscribers[event_type].append(fn)

    async def post_event(self, event_type: str, event_data):
        if event_type not in self.subscribers:
            return

        tasks = [fn(event_data) for fn in self.subscribers[event_type]]
        results = await asyncio.gather(*tasks)
        for fn, result in zip(self.subscribers[event_type], results):
            self.results[f"{fn.__self__.__class__.__name__}.{fn.__name__}"] = result
