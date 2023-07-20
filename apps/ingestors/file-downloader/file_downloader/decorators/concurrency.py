# import asyncio
from functools import wraps
import concurrent.futures


def run_with_concurrency(concurrency):
    def decorator(func):
        def wrapper(*args, **kwargs):
            with concurrent.futures.ThreadPoolExecutor(max_workers=concurrency) as executor:
                tasks = [executor.submit(func, *args, **kwargs) for _ in range(concurrency)]
                concurrent.futures.wait(tasks, return_when=concurrent.futures.ALL_COMPLETED)

        return wrapper
    return decorator
