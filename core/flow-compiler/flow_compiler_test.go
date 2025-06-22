package flowcompiler

import (
	"testing"
)

func TestFlowCompiler_Compile(t *testing.T) {
	endpoints := []Endpoint{
		{ID: "getUser", Path: "/user", Method: "GET"},
		{ID: "syncData", Path: "/sync", Method: "POST"},
	}
	flows := []FlowDefinition{
		{
			WorkflowID: "sync-user-data",
			Steps: []FlowStep{
				{ID: "step1", Call: "getUser", PreHook: "pre1.go", PostHook: "post1.go"},
				{ID: "step2", Call: "syncData", PreHook: "pre2.go", PostHook: "post2.go"},
			},
		},
	}
	compiler := NewFlowCompiler(endpoints, flows)
	compiled, err := compiler.Compile()
	if err != nil {
		t.Fatalf("Compile failed: %v", err)
	}
	if len(compiled) != 1 {
		t.Fatalf("Expected 1 compiled flow, got %d", len(compiled))
	}
	cf := compiled[0]
	if cf.WorkflowID != "sync-user-data" {
		t.Errorf("Expected WorkflowID 'sync-user-data', got '%s'", cf.WorkflowID)
	}
	if len(cf.Steps) != 2 {
		t.Errorf("Expected 2 steps, got %d", len(cf.Steps))
	}
	if cf.Steps[0].Endpoint.ID != "getUser" || cf.Steps[1].Endpoint.ID != "syncData" {
		t.Errorf("Step endpoint IDs do not match expected")
	}
}
