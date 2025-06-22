package codegenerator

// LLMProvider defines the interface for any large language model provider.
type LLMProvider interface {
	GenerateCode(prompt string) (string, error)
}

// OpenAIProvider implements LLMProvider for OpenAI models.
type OpenAIProvider struct {
	KeyPath string // Path to API key file
	Model   string // Model name (e.g., gpt-4-1106-preview)
}

// GenerateCode sends a prompt to the OpenAI API and returns the generated code.
func (o *OpenAIProvider) GenerateCode(prompt string) (string, error) {
	// Read OpenAI API key from file
	keyBytes, err := os.ReadFile(o.KeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read OpenAI API key from %s: %w", o.KeyPath, err)
	}
	apiKey := string(bytes.TrimSpace(keyBytes))
	if apiKey == "" {
		return "", fmt.Errorf("OpenAI API key is empty in %s", o.KeyPath)
	}

	// Prepare request body for OpenAI API
	apiURL := "https://api.openai.com/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": o.Model,
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
