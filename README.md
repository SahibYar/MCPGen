# MCPGen
**Auto-generate a multiservice MCP (Message Control Protocol) server from OpenAPI v2 (Swagger), v3 (Arazzo) Spec Files**
Turn API definitions and workflow specs into a smart,
pluggable server that can coordinate real tasks—locally or in the cloud. 

## What is MCPGen ?
MCGen is a tool that builds a fully functional MCP Server —
a control plane that coordinates operations across multiple APIs.
It takes
  * One or more **OpenAPI/Swagger specs** (REST APIs)
  * Optional Arazzo workflow specs - (task flows and orchestration)
...and generates a runnable sever that
  * Exposes task-based endpoints (like `/run-job`)
  * Executes coordinated workflows across services
  * Injects pre/post hooks, retries, validations
  * Wraps your microservices in a clean, centralized control layer

## Why Use MCP ?
Modern systems use **microservices**, each with their own APIs. Exposing all of them to external world via (CLI, Rest end-points) is messy and risky. 
RestAPIs and CLIs are great—but they're not always friendly for end user 
### MCPGen gives you one control point:
* One server, multiple services
* One endpoint, many API calls
* Easier to manage, monitor, and secure.
* Encapsulates logic in a clean interface.
* Lets clients say `POST /run-job`, and MCP figures out the rest
Think of it as:
> Swagger + Arazzo -> MCP -> Task Engine

## How It Works
#### 1. Install MCPGen
```bash
go install github.com/sahibyar/mcpgen@latest
```
#### 2. Prepare your specs
* Your OpenAPI/Swagger specs (can be local or remote), locally it can be `.json` or `.yaml` format
* Optional Arazzo task spec files (JSON or YAML)

#### 3. Generate the MCP Server
```bash
mcpgen \
  --specs ./specs/service-a.yaml, ./specs/service-b.yaml \
  --arazzo ./workflows/task.yaml \
  --output ./mcp-server
```

#### 4. Run the server
```bash
cd mcp-server
go run main.go
```

Now your clients can call:
```http
POST /run-task/sync-user-data
```
...and MCP will handle the entire coordinates flow. 

Task Coordination Layer (Arazzo + MCP)
MCPGen supports **task-based routing and orchestration** using **Arazzo specification files**, which defined how multiple APIs should work together in a single flow

Example Arazzo Workflow:
```yaml
workflowId: sync-user-data
summary: Sync user data
inputs:
  type: object
  properties:
        username:
          type: string
        password: 
          type: User ID to sync
description: Sync user info from Service A go Service B
steps:
  -  id: getUser
     call: user-data
     pre_hook: hooks/validate_user.go
  -  id: syncData
     call: data-service
     postHooks: hooks/log_result.go
```

Each task supports:
* id: Tied to OpenAPI operations
* service: From your provider Swagger specs
* preHook/postHook: Custom Go middleware or logic
* Optional condition, retries, timeout, etc.
## Features
* Supports **multiple OpenAPI specs**
* Supports **Arazzo task coordination specs**
* Generates a ready-to-run **MCP server in Go**
* **Pre- / post-hooks** injection points
* Validates request/response payloads.
* Configurable **middlewares, error handling, and retry logic**
* **Docker-ready** build output

## 🏗️ High-Level Architecture
```aiignore
+-----------------------------------------------------------------------+
|   +---------------------------+                                       |
|   |     OpenAPI Parser        |  <-- Reads Swagger files (YAML/JSON)  |
|   +---------------------------+                                       |
|                                                                       |
|                                                                       |
|   +---------------------------+                                       |
|   |    Arazzo Flow Parser     |  <-- Reads custom flow definitions    |
|   +---------------------------+                                       |
|                                                                       |
+-----------------------------------------------------------------------+
            |
            v
+----------------------------+
|     MCP Flow Compiler      |  <-- Combines OpenAPI + Arazzo into flow DAGs
+----------------------------+
            |
            v
+----------------------------+
|      Code Generator        |  <-- Generates Go/Python/Node MCP server
+----------------------------+
            |
            v
+----------------------------+
|     Generated MCP Server   |  <-- Supports hooks, auth, error handling
+----------------------------+
```

## 🧠 Data Flow Diagram
```aiignore
[Swagger YAML]         [Arazzo YAML]
     |                     |
     v                     v
[OpenAPILoader]       [ArazzoParser]
     |                     |
     +----------+----------+
                |
                v
         [FlowCompiler]
                |
                v
         [CodeGenerator]
                |
                v
       [MCP Server Output (Go/Py/Node)]
                |
                v
       [Includes Hooks, Logging, Auth]
```

## 🔄 Supported MCPGen Flows & Execution Diagram

The following diagram shows how MCPGen coordinates multiple services and flows, supporting advanced features like hooks, conditions, and multi-step orchestration:

```aiignore
[User/API/CLI Request]
         |
         v
+------------------------------+
|        MCPGen Server         |
+------------------------------+
         |
         v
+------------------------------+
|      Task Router/Dispatcher  |
+------------------------------+
         |
         v
+------------------------------+
|         Flow Engine          |
|  (Executes Steps Sequentially|
|   or Conditionally)          |
+------------------------------+
         |
         v
+------------------------------+
|        Step Executor         |
|  (Calls API, Runs Hook, etc) |
+------------------------------+
         |
         v
+------------------------------+
|   Aggregates Results/Errors  |
+------------------------------+
         |
         v
+------------------------------+
|     Returns Response         |
+------------------------------+
```

---

## 🤖 RAG-Enabled Code Generation: Go ↔ Python Integration

MCPGen can use Retrieval Augmented Generation (RAG) for smarter code generation by integrating Go with a Python microservice. Here’s how the flow works:

```aiignore
[Go: CodeGenerator (RAGProvider)]
         |
         |  (HTTP POST /generate)
         v
[Python: rag_service (FastAPI)]
         |
         |  (RAGPipeline retrieves context from Redis)
         v
[Redis Vector DB] <---+--- [Python: modelcontextprotocol/python-sdk]
         |
         |  (Augmented prompt)
         v
[Python: Calls OpenAI LLM]
         |
         |  (Generated code)
         v
[Python: Returns code to Go]
         |
         v
[Go: Uses generated code in MCP flow]
```

### Step-by-Step:
1. **Go (RAGProvider)** sends a prompt to the Python `/generate` endpoint.
2. **Python (rag_service)** retrieves relevant context from Redis using the python-sdk.
3. The context is combined with the prompt and sent to OpenAI’s LLM.
4. The generated code is returned to the Go app and used in the MCP flow.

This enables context-aware, retrieval-augmented code generation for your MCP workflows.

### Example Execution Flow
```yaml
flow: ProcessUserOrder
steps:
  - id: ValidateInput
    call: UserService.validateInput
    pre_hook: checkRateLimit
  - id: CreateOrder
    call: OrderService.createOrder
    post_hook: notifyAnalytics
  - id: Payment
    call: PaymentService.initiate
    conditional_on: CreateOrder.status == "success"
```

This flow will:
- Validate input (with pre-hook)
- Create an order (with post-hook)
- Only initiate payment if order creation succeeds

### Internal Representation (Go struct)
```go
// Represents a step in the flow
type Step struct {
	ID           string
	ServiceCall  Endpoint
	PreHook      string
	PostHook     string
	Condition    *ConditionExpr
}
```

### Use Case Example
- You have a `user-service` and a `sync-service`
- You want a single endpoint `/run-task/sync-user-data` that:
  - Fetches user info
  - Pushes data to an external system
  - Applies hooks and error handling

MCPGen lets you define this as a flow and exposes it as a single, orchestrated endpoint.

## ⚙️ Detailed Design: Execution Flow
Arazzo + OpenAPI combined to define flows like this:
```yaml
flow: ProcessUserOrder
steps:
  - id: ValidateInput
    call: UserService.validateInput
    pre_hook: checkRateLimit
  - id: CreateOrder
    call: OrderService.createOrder
    post_hook: notifyAnalytics
  - id: Payment
    call: PaymentService.initiate
    conditional_on: CreateOrder.status == "success"
```

## Internally becomes:
```golang
type Step struct {
	ID           string
	ServiceCall  Endpoint
	PreHook      string
	PostHook     string
	Condition    *ConditionExpr
}
```
## Use Case Example
You have:
* A `user-service` that fetches user info
* A `sync-service` that pushes data to external systems

You want:
* A CLI command like `mycli sync-user --id 123`

MCPGen gives you:
* A single `/run-task/sync-user-data` endpoint
* Internally calls both APIs, applies hooks, handles responses
* CLI stays simple, secure, and decoupled.

## LongTerm Roadmap
* Visual Arazzo editor (web-based)
* gRPC + REST hybrid workflows.
* Swagger UI integration for MCP endpoints.
* Live hot reload of specs and tasks
* Support for tracing + metrics exports

## Contributing
MCPGen is early, and we'd love your help. 
PRs welcome for:
* New hooks or middlewares
* Workflow enhancements
* Support for another language
* Docs, examples, tests

## License
MIT License