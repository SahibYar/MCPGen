# RAG-Enabled MCP Code Generation: Testing Guide

This guide shows you how to set up, ingest data, run, and test the RAG (Retrieval Augmented Generation) workflow for Model Context Protocol (MCP) code generation using OpenAI, Redis, and the modelcontextprotocol/python-sdk.

---

## 1. Example Go Code to Call the Python RAG Service

```go
package main

import (
	"fmt"
	"log"
	"yourmodule/codegenerator" // adjust import path as needed
)

func main() {
	rag := &codegenerator.RAGProvider{
		Endpoint: "http://localhost:8000/generate",
	}

	prompt := "Generate idiomatic Go server code for a REST API with JWT authentication."
	code, err := rag.GenerateCode(prompt)
	if err != nil {
		log.Fatalf("RAG code generation failed: %v", err)
	}
	fmt.Println("Generated code:\n", code)
}
```
- Replace `"yourmodule/codegenerator"` with the correct import path.

---

## 2. Ingesting Documents into Redis

You need to populate Redis with relevant context (docs, code, specs) for RAG to work. Here’s a simple script:

```python
# rag_ingest.py
from modelcontextprotocol.rag import RAGPipeline
import glob

rag = RAGPipeline(vector_db='redis', redis_url='redis://localhost:6379/0')

# Example: Ingest a folder of .md, .go, .yaml files
files = glob.glob("path/to/your/docs/**/*.*", recursive=True)
for f in files:
    with open(f, "r") as file:
        content = file.read()
        rag.add_document(content, metadata={"filename": f})

print("Ingestion complete!")
```
- Replace `"path/to/your/docs"` with your documentation/code directory.
- Run with: `python rag_ingest.py`

---

## 3. End-to-End Test Instructions

### A. Start Redis
If not already running:
```bash
redis-server
```

### B. Ingest Documents
```bash
cd rag_service
python rag_ingest.py
```

### C. Start the Python RAG Service
```bash
export OPENAI_API_KEY=your-openai-key
export OPENAI_MODEL=gpt-4-1106-preview
export REDIS_URL=redis://localhost:6379/0
uvicorn rag_service:app --host 0.0.0.0 --port 8000
```

### D. Call from Go
Use the Go example above, or integrate into your MCPGen flow. The Go `RAGProvider` will send prompts to the Python service and receive generated code.

---

## 4. Troubleshooting

- **No context returned**: Ensure your documents are ingested and Redis is running.
- **OpenAI errors**: Check your API key, model name, and network/firewall settings.
- **Connection refused**: Make sure the Python service is running on the correct port.
- **Go import issues**: Double-check your Go module paths and dependencies.
- **Slow responses**: Try reducing `top_k` in the Python service or limit document sizes.

---

## 5. (Optional) Test the Python Service Directly

```bash
curl -X POST http://localhost:8000/generate -H "Content-Type: application/json" -d '{"prompt":"Generate a Go HTTP server with a /hello endpoint."}'
```

---

Let us know if you want a more advanced ingestion script, more Go integration examples, or help with debugging any part of this workflow!
