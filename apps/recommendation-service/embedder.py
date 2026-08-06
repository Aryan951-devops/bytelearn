from sentence_transformers import SentenceTransformer

class VectorEmbedder:
    def __init__(self, model_name: str):
        # Load model ONCE during initialization
        self.model = SentenceTransformer(model_name)

    def encode(self, text: str) -> list[float]:
        embedding = self.model.encode(text)
        return embedding.tolist()