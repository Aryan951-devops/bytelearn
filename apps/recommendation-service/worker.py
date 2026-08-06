import json
import logging
import time
import redis

from config import Config

logger = logging.getLogger(__name__)


def start_worker(service):

    redis_client = redis.Redis(
        host=Config.REDIS_HOST,
        port=Config.REDIS_PORT,
        db=0,
        socket_timeout=None,
        socket_connect_timeout=5,
        health_check_interval=30,
    )

    logger.info(
        "Listening on queue '%s'...",
        Config.REDIS_QUEUE,
    )

    while True:
        try:
            # BRPOP blocks until a job arrives 
            _, raw_payload = redis_client.brpop(
                Config.REDIS_QUEUE,
                timeout=0,
            )

        except redis.exceptions.ConnectionError:
            logger.exception("Redis connection lost")
            time.sleep(5)
            continue

        except redis.exceptions.TimeoutError:
            logger.exception("Redis timeout")
            continue

        except Exception as e:
            logger.exception(f"Worker encountered error: {e}")
            time.sleep(1)

        try:
            event = json.loads(raw_payload.decode())
            logger.info(
                "Received event: %s",
                event,
            )

            service.process_event(event)

        except Exception:
            logger.exception(
                "Failed processing recommendation event"
            )