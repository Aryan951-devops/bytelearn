import logging
import uuid
from fastapi import FastAPI, Query, HTTPException, status
from pydantic import BaseModel


logger = logging.getLogger(__name__)

class SearchResultItem(BaseModel):
    video_id: uuid.UUID
    title: str
    similarity: float

class SearchResponse(BaseModel):
    results: list[SearchResultItem]
    count: int

def create_app(service):
    app = FastAPI(title="Recommendation Service")

    @app.get("/search", response_model=SearchResponse, status_code=status.HTTP_200_OK)
    def search(
        q: str = Query(..., min_length=1, description="Search term or topic"),
        limit: int = Query(10, ge=1, le=50, description="Max results to return (default 10)")
    ):
        try:
            results = service.search_videos(query_text=q, limit=limit)
            return {
                "results": results,
                "count": len(results)
            }
        except Exception as e:
            logger.exception("Error executing search query")
            raise HTTPException(
                status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
                detail="Failed to perform search"
            )

    return app