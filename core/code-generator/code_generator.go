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
	OpenAIKey  string
	Flow       *CompiledFlow
	OutputDir  string
}

// GenerateServerCode constructs a prompt, calls OpenAI API, and writes code to file.
func (cg *CodeGenerator) GenerateServerCode() error {
	prompt, err := cg.constructPrompt()
	if err != nil {
		return err
	}

	code, err := cg.callOpenAI(prompt)
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

func (cg *CodeGenerator) callOpenAI(prompt string) (string, error) {
	// Read OpenAI API key from file
	keyPath := "./core/code-generator/.openai_key"
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read OpenAI API key from %s: %w", keyPath, err)
	}
	apiKey := string(bytes.TrimSpace(keyBytes))
	if apiKey == "" {
		return "", fmt.Errorf("OpenAI API key is empty in %s", keyPath)
	}

	// Prepare request body for OpenAI API (using gpt-3.5-turbo)
	apiURL := "https://api.openai.com/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": "gpt-4-1106-preview", // GPT-4.1 as default
		"messages": []map[string]string{{
			"role":    "user",
			"content": prompt,
		}},
		"max_tokens": 2048,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal OpenAI request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create OpenAI request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenAI API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("OpenAI API error: %s, body: %s", resp.Status, string(body))
	}

	var respData struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", fmt.Errorf("failed to decode OpenAI response: %w", err)
	}
	if len(respData.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI API")
	}
	return respData.Choices[0].Message.Content, nil
}
