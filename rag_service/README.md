# RAG Microservice for MCP Code Generation

This Python microservice uses the modelcontextprotocol/python-sdk and Redis as a vector database to enable Retrieval Augmented Generation (RAG) for Model Context Protocol (MCP) code generation.

## Features
- Stores and retrieves context from Redis using the python-sdk.
- Augments user prompts with retrieved context.
- Calls OpenAI's LLM (GPT-4.1 by default) to generate code.
- Designed to be called from the Go MCPGen project as a pluggable LLMProvider.

## Usage

1. **Install dependencies:**
   ```bash
   pip install -r requirements.txt
   ```
2. **Set environment variables:**
   - `REDIS_URL` (default: redis://localhost:6379/0)
   - `OPENAI_API_KEY` (your OpenAI key)
   - `OPENAI_MODEL` (default: gpt-4-1106-preview)
3. **Run the service:**
   ```bash
   uvicorn rag_service:app --host 0.0.0.0 --port 8000
   ```
4. **Call from Go:**
   Use the `RAGProvider` in Go, setting `Endpoint` to `http://localhost:8000/generate`.

## Example Request
```json
{
  "prompt": "Generate idiomatic Go server code for ..."
}
```

## Example Response
```json
{
  "code": "// generated Go code ..."
}
```

---

See the main MCPGen README for Go integration details.
