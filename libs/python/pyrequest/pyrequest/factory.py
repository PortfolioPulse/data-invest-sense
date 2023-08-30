import httpx
import asyncio

class RateLimitedAsyncHttpClient:
    def __init__(self, base_url, max_calls, period):
        self.base_url = base_url
        self.max_calls = max_calls
        self.period = period
        self.semaphore = asyncio.Semaphore(max_calls)

    async def make_request(self, method, endpoint, data=None, params=None):
        url = self.base_url + endpoint
        async with self.semaphore:
            async with httpx.AsyncClient() as client:
                response = await client.request(method, url, json=data, params=params)
                response.raise_for_status()
                return response.json()
