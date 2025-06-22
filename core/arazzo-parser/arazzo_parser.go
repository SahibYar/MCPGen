package arazzoparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Workflow represents a parsed workflow/task definition.
type Workflow struct {
	Name   string            `json:"name" yaml:"name"`
	Steps  []WorkflowStep    `json:"steps" yaml:"steps"`
	Format string            `json:"format" yaml:"-"`
}

type WorkflowStep struct {
	ID       string                 `json:"id" yaml:"id"`
	Type     string                 `json:"type" yaml:"type"`
	Inputs   map[string]interface{} `json:"inputs" yaml:"inputs"`
	Next     []string               `json:"next" yaml:"next"`
}

// ParseArazzoWorkflow parses a YAML or JSON Arazzo workflow file into a Workflow struct.
func ParseArazzoWorkflow(path string) (*Workflow, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open workflow file: %w", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read workflow file: %w", err)
	}

	var wf Workflow
	if yaml.Unmarshal(data, &wf) == nil && wf.Name != "" {
		wf.Format = "yaml"
		return &wf, nil
	}

	if json.Unmarshal(data, &wf) == nil && wf.Name != "" {
		wf.Format = "json"
		return &wf, nil
	}

	return nil, fmt.Errorf("file is not a valid Arazzo workflow YAML or JSON")
}
