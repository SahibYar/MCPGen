import requests
from bs4 import BeautifulSoup
from modelcontextprotocol.rag import RAGPipeline

rag = RAGPipeline(vector_db='redis', redis_url='redis://localhost:6379/0')

# Download the README.md from GitHub
url = "https://raw.githubusercontent.com/modelcontextprotocol/python-sdk/refs/heads/main/README.md"
response = requests.get(url)
readme_text = response.text

# Optionally, you could chunk the README for better retrieval granularity
# For simplicity, we'll ingest the whole README as one document
rag.add_document(readme_text, metadata={"source": url})

print("modelcontextprotocol/python-sdk README.md ingested into Redis!")
