package codegenerator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RAGProvider struct {
	Endpoint string // e.g., "http://localhost:8000/generate"
}

func (r *RAGProvider) GenerateCode(prompt string) (string, error) {
	requestBody := map[string]string{"prompt": prompt}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal RAG request body: %w", err)
	}

	resp, err := http.Post(r.Endpoint, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to call RAG service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("RAG service error: %s, body: %s", resp.Status, string(body))
	}

	var respData struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", fmt.Errorf("failed to decode RAG service response: %w", err)
	}
	return respData.Code, nil
}
