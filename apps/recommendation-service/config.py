import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    REDIS_HOST = os.getenv("REDIS_HOST", "localhost")
    REDIS_PORT = int(os.getenv("REDIS_PORT", 6379))
    REDIS_QUEUE = os.getenv("REDIS_QUEUE", "recommendation_jobs")
    
    DATABASE_URL = os.getenv("DATABASE_URL", "")
    API_GATEWAY_URL = os.getenv("API_GATEWAY_URL", "http://localhost:8080/api/v1")
    
    EMBEDDING_MODEL = os.getenv("EMBEDDING_MODEL", "all-MiniLM-L6-v2")
    WORKER_CONCURRENCY = int(os.getenv("WORKER_CONCURRENCY", 4))