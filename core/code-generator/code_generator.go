package codegenerator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// CompiledFlow is a placeholder for the actual compiled flow structure.
type CompiledFlow struct {
	WorkflowID string
	Steps      []CompiledStep
}

type CompiledStep struct {
	StepID   string
	Type     string
	Inputs   map[string]interface{}
	Next     []string
}

// CodeGenerator generates Go server code using OpenAI API.
type CodeGenerator struct {
	Flow       *CompiledFlow
	OutputDir  string
	LLM        LLMProvider // Strategy Pattern: pluggable provider
}

// GenerateServerCode constructs a prompt, calls OpenAI API, and writes code to file.
func (cg *CodeGenerator) GenerateServerCode() error {
	prompt, err := cg.constructPrompt()
	if err != nil {
		return err
	}

	if cg.LLM == nil {
		return fmt.Errorf("no LLM provider configured")
	}

	code, err := cg.LLM.GenerateCode(prompt)
	if err != nil {
		return err
	}

	outputFile := cg.OutputDir + "/server.go"
	if err := ioutil.WriteFile(outputFile, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write generated code: %w", err)
	}
	return nil
}

func (cg *CodeGenerator) constructPrompt() (string, error) {
	flowJson, err := json.MarshalIndent(cg.Flow, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal flow: %w", err)
	}
	prompt := "Generate idiomatic Go server code for the following workflow. Include endpoint handlers, orchestration, hooks, and error handling.\n" + string(flowJson)
	return prompt, nil
}


