import uuid
import psycopg2
from pgvector.psycopg2 import register_vector
from pgvector import Vector

class EmbeddingRepository:
    def __init__(self, db_url: str):
        self.conn = psycopg2.connect(db_url)
        self.conn.autocommit = True
        register_vector(self.conn)

    def upsert_embedding(self, video_id: uuid.UUID, title: str, document: str, vector: list[float]):
        with self.conn.cursor() as cur:
            query = """
                INSERT INTO video_embeddings (video_id, title, document, embedding, updated_at)
                VALUES (%s, %s, %s, %s, NOW())
                ON CONFLICT (video_id) DO UPDATE
                SET title = EXCLUDED.title,
                    document = EXCLUDED.document,
                    embedding = EXCLUDED.embedding,
                    updated_at = NOW();
            """
            cur.execute(query, (str(video_id), title, document, vector))

    def delete_embedding(self, video_id: uuid.UUID):
        with self.conn.cursor() as cur:
            cur.execute("DELETE FROM video_embeddings WHERE video_id = %s;", (str(video_id),))


    def search_similar_embeddings(self, query_vector: list[float], limit: int = 10) -> list[dict]:
        with self.conn.cursor() as cur:
            vec_param = Vector(query_vector)

            # Cosine distance operator (<=>) orders by closeness (0 = exact match)
            query = """
                SELECT video_id, title, document, 1 - (embedding <=> %s) AS similarity
                FROM video_embeddings
                ORDER BY embedding <=> %s
                LIMIT %s;
            """
            cur.execute(query, (vec_param, vec_param, limit))
            rows = cur.fetchall()

            results = []
            for row in rows:
                results.append({
                    "video_id": str(row[0]),
                    "title": row[1],
                    "document": row[2],
                    "similarity": float(row[3])
                })
            return results