from pytube import YouTube
import ssl
from pylog.log import setup_logging

logger = setup_logging(__name__)

ssl._create_default_https_context = ssl._create_unverified_context

def download_video(link):
    youtubeObject = YouTube(link)
    logger.info(f"youtubeObject.vid_info: {youtubeObject.vid_info}")
    # logger.info(f"youtubeObject.embed_html: {youtubeObject.embed_html}")
    youtubeObject = youtubeObject.streams.get_highest_resolution()
    try:
        youtubeObject.download()
    except:
        logger.error("An error has occurred")
    # logger.info(f"youtubeObject.vid_info: {youtubeObject.vid_info}")
    # logger.info(f"youtubeObject.length: {youtubeObject.length}")
    # logger.info(f"youtubeObject.views: {youtubeObject.views}")
    # logger.info(f"youtubeObject.author: {youtubeObject.author}")
    logger.info("Download is completed successfully")


if __name__ == "__main__":
    link = "https://www.youtube.com/shorts/"
    download_video(link)
