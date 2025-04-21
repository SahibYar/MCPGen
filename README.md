# MCPGen
**Auto-generate a multi-service MCP (Message Control Protocol) server from OpenAPI (Swagger) + Arazzo spec files**
Turn API definitions and workflow specs into a smart, pluggable server that can coordinate real tasks - locally or in cloud. 

## What is MCPGen ?
MCGen is an opensource tool that builds builds a fully functional MCP Server - a control plane that coordinates operations across multiple APIs.
It takes
  * One or more **OpenAPI/Swagger specs** (REST APIs)
  * Optional Arazzo workflow specs - (task flows and orchestration)
...and generates a runnable sever that
  * Exposes task-based endpoints (like `/run-job`)
  * Executes coordinated workflows across services
  * Injects pre/post hooks, retries, validations
  * Wraps your microservices in a clean, centralized control layer

## Why Use MCPGen ?
Modern systems use **microservices**, each with their own APIs. Exposing all of them to clients (CLIs, users, integration) is messy and risky. 
Microservices and APIs are great - but they're not always friends for CLI tools, integrations, or task-based workflows. 

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
task: sync-user-data
description: Sync user info from Service A go Service B
steps:
  -  operationId: getUser
     service: user-data
     preHook: hooks/validate_user.go
  -  operationIdL syncData
     service: data-service
     postHooks: hooks/log_result.go
```

Each task supports:
* operationId: Tied to OpenAPI operations
* service: From your provider Swagger specs
* preHook/postHook: Custom Go middleware or logic
* Optional condition, retries, timeout, etc.
## Features
* Supports **multiple OpenAPI specs**
* Supports **Arazzo task coordination specs**
* Generates a ready-to-run **MCP server in Go**
* **Pre/post hooks** injection points
* Validates request/response payloads.
* Configurable **middlewares, error handling, and retry logic**
* **Docker-ready** build output

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

## Roadmap
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
* Support for other language
* Docs, examples, tests

## License
MIT License


