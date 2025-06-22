package flowcompiler

import (
	"fmt"
)

// FlowCompiler merges endpoints and workflows into executable flows.
type FlowCompiler struct {
	Endpoints []Endpoint
	Flows     []FlowDefinition
}

// NewFlowCompiler creates a new instance of FlowCompiler.
func NewFlowCompiler(endpoints []Endpoint, flows []FlowDefinition) *FlowCompiler {
	return &FlowCompiler{
		Endpoints: endpoints,
		Flows:     flows,
	}
}

// Compile merges endpoints and flows into CompiledFlow objects.
func (fc *FlowCompiler) Compile() ([]*CompiledFlow, error) {
	var compiledFlows []*CompiledFlow
	endpointMap := make(map[string]*Endpoint)
	for i, ep := range fc.Endpoints {
		endpointMap[ep.ID] = &fc.Endpoints[i]
	}

	for _, flow := range fc.Flows {
		cf := &CompiledFlow{
			WorkflowID: flow.WorkflowID,
		}
		for _, step := range flow.Steps {
			endpoint, ok := endpointMap[step.Call]
			if !ok {
				return nil, fmt.Errorf("endpoint with ID '%s' not found for step '%s' in workflow '%s'", step.Call, step.ID, flow.WorkflowID)
			}
			cf.Steps = append(cf.Steps, CompiledStep{
				StepID:   step.ID,
				Endpoint: endpoint,
				PreHook:  step.PreHook,
				PostHook: step.PostHook,
			})
		}
		compiledFlows = append(compiledFlows, cf)
	}
	return compiledFlows, nil
}
