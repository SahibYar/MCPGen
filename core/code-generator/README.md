Responsibilities:
*	Generate MCP server code from CompiledFlow
*	Insert hooks, error handling, logging, etc.

```golang
type CodeGenerator struct {
    Flow *CompiledFlow
    OutputDir string
}
func (cg *CodeGenerator) GenerateServerCode() error
```