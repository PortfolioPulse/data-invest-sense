from collections import defaultdict

subscribers = defaultdict(list)

def subscribe(event_type: str, fn):
    if event_type not in subscribers:
        subscribers[event_type] = []
    subscribers[event_type].append(fn)


def post_event(event_type: str, event_data):
    if not event_type in subscribers:
        return
    for fn in subscribers[event_type]:
        yield fn(event_data)
