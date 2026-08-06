import threading
import uvicorn
import logging

from config import Config
from embedder import VectorEmbedder
from repository import EmbeddingRepository
from service import RecommendationService

from app import create_app
from worker import start_worker


logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
)

def main():

    embedder = VectorEmbedder(Config.EMBEDDING_MODEL)

    repo = EmbeddingRepository(Config.DATABASE_URL)

    service = RecommendationService(
        repo,
        embedder,
        Config.API_GATEWAY_URL,
    )

    print("Starting Worker Thread which listens to redis!")
    worker = threading.Thread(
        target=start_worker,
        args=(service,),
        daemon=True,
    )
    worker.start()

    app = create_app(service)
    uvicorn.run(app, host="0.0.0.0", port=8001)

if __name__ == "__main__":
    main()