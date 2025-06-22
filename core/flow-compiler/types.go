package flowcompiler

// Endpoint represents a parsed API endpoint from OpenAPI/Swagger.
type Endpoint struct {
	ID          string
	Path        string
	Method      string
	Summary     string
	Parameters  []Parameter
	Responses   map[string]Response
	// ... other fields as needed
}

type Parameter struct {
	Name     string
	In       string
	Required bool
	Schema   interface{}
}

type Response struct {
	Code    string
	Schema  interface{}
}

// FlowDefinition represents a parsed Arazzo workflow.
type FlowDefinition struct {
	WorkflowID string
	Steps      []FlowStep
	// ... other workflow fields
}

type FlowStep struct {
	ID        string
	Call      string
	PreHook   string
	PostHook  string
	// ... other fields
}

// CompiledFlow is the result of merging endpoints and workflows.
type CompiledFlow struct {
	WorkflowID string
	Steps      []CompiledStep
}

type CompiledStep struct {
	StepID     string
	Endpoint   *Endpoint
	PreHook    string
	PostHook   string
}
