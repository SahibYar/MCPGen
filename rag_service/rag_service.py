from fastapi import FastAPI, Request
from modelcontextprotocol.rag import RAGPipeline
import openai
import os

app = FastAPI()

# Configure RAG pipeline to use Redis
REDIS_URL = os.getenv("REDIS_URL", "redis://localhost:6379/0")
rag = RAGPipeline(vector_db='redis', redis_url=REDIS_URL)

# Read OpenAI API key directly from file
with open("/Users/sahibyar/Work/MCPGen/core/code-generator/.openai_key", "r") as f:
    OPENAI_API_KEY = f.read().strip()
OPENAI_MODEL = os.getenv("OPENAI_MODEL", "gpt-4-1106-preview")

@app.post('/generate')
async def generate(request: Request):
    data = await request.json()
    prompt = data['prompt']
    # Retrieve context from Redis using RAG
    context = rag.retrieve(prompt, top_k=5)
    augmented_prompt = context + '\n' + prompt
    # Call OpenAI LLM
    response = openai.ChatCompletion.create(
        model=OPENAI_MODEL,
        messages=[{"role": "user", "content": augmented_prompt}],
        max_tokens=2048,
        api_key=OPENAI_API_KEY
    )
    return {"code": response['choices'][0]['message']['content']}
