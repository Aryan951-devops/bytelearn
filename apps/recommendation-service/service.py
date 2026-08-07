import time
import uuid
import logging
import requests
from embedder import VectorEmbedder
from repository import EmbeddingRepository
from pydantic import BaseModel
from enum import Enum

logger = logging.getLogger(__name__)

class RecommendationEventType(str, Enum):
    VIDEO_CREATED = "video_created"
    VIDEO_UPDATED = "video_updated"
    VIDEO_DELETED = "video_deleted"

class RecommendationJob(BaseModel):
    video_id: uuid.UUID
    event_type: str

class RecommendationService:
    def __init__(self, repo: EmbeddingRepository, embedder: VectorEmbedder, api_gateway_url: str):
        self.repo = repo
        self.embedder = embedder
        self.api_gateway_url = api_gateway_url

    def process_event(self, event: dict):
        job = RecommendationJob.model_validate(event)

        event_type = job.event_type
        video_id = job.video_id

        if event_type in [RecommendationEventType.VIDEO_CREATED, RecommendationEventType.VIDEO_UPDATED]:
            self._sync_video(video_id)
        elif event_type == RecommendationEventType.VIDEO_DELETED:
            self.repo.delete_embedding(video_id)
            logger.info(f"Deleted embedding for video_id: {video_id}")
        else:
            logger.warning(f"Unknown event type: {event_type}")

    def _sync_video(self, video_id: uuid.UUID):
        meta = self._fetch_video_metadata(video_id)
        if not meta:
            return

        data = meta.get("data", "")
        video = data.get("video", "")
        title = video.get("title", "")
        desc = video.get("description", "")

        document = f"Title: {title}. Description: {desc}."
        print(document)

        vector = self.embedder.encode(document)

        self.repo.upsert_embedding(video_id, title, document, vector)
        logger.info(f"Upserted embedding for video_id: {video_id}")

    def _fetch_video_metadata(self, video_id: uuid.UUID) -> dict | None:
        try:
            # time.sleep(10) # for testing only. Need to remove this while pushing
            url = f"{self.api_gateway_url}/video/{video_id}"
            logger.info("Calling Gateway: %s", url)

            response = requests.get(url, timeout=5)
            if response.status_code == 200:
                return response.json()
            
            logger.error(f"Failed to fetch metadata, status: {response.status_code}")
        except Exception as e:
            logger.error(f"Error calling Gateway API for video_id {video_id}: {e}")
        return None

    def search_videos(self, query_text: str, limit: int = 10) -> list[dict]:
        if not query_text or not query_text.strip():
            return []
            
        # 1. Encode query string into vector embedding
        query_vector = self.embedder.encode(query_text)

        # 2. Query nearest neighbor vectors in DB
        results = self.repo.search_similar_embeddings(query_vector, limit=limit)
        return results