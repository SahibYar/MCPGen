Responsibilities:
*	Combine OpenAPI + Arazzo into an executable flow
*	Represent as DAG / sequence of operations
```golang
type FlowCompiler struct {
    Endpoints []Endpoint
    Flows     []FlowDefinition
}

func (fc *FlowCompiler) Compile() []*CompiledFlow
```