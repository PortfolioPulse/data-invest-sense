from pylog.log import setup_logging
from controller.events.event import EventObserver

logger = setup_logging(__name__)



class MongoDBEventListener:
    def __init__(self, observer: EventObserver):
        self.__observer = observer
        self.set_event_handlers()

    def set_event_handlers(self):
        self.__observer.add_subscribe("check_config", self.handle_job_should_start_event)

    async def handle_job_should_start_event(self, event_data):
        # logger.info(f"[MONGO EVENT] handle_job_should_start_event: {event_data}")
        event_result = self.__observer.config
        return event_result






