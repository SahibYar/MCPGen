package open_api_loader

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// 1. Valid file input test
func TestParseOpenAPISpecsFile_ValidFile(t *testing.T) {
	summary, err := ParseOpenAPISpecsFile("testdata/valid_openapi.yaml")
	assert.NoError(t, err)
	assert.Contains(t, summary, "OpenAPI version:")
	assert.Contains(t, summary, "paths")
	assert.Contains(t, summary, "schemas")
}

// 2. File not found test
func TestParseOpenAPISpecsFile_FileNotFound(t *testing.T) {
	_, err := ParseOpenAPISpecsFile("testdata/missing_file.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot read file")
}

// 3. Invalid YAML format
func TestParseOpenAPISpecsFile_InvalidYAML(t *testing.T) {
	// Create a temporary invalid file
	tmpFile := "testdata/invalid_openapi.yaml"
	_ = os.WriteFile(tmpFile, []byte("invalid: yaml: : content"), 0644)
	defer os.Remove(tmpFile)

	_, err := ParseOpenAPISpecsFile(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create document")
}

// 4. OpenAPI model build errors (simulate)
func TestParseOpenAPISpecsFile_ModelErrors(t *testing.T) {
	_, err := ParseOpenAPISpecsFile("testdata/model_error_openapi.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create v3 model")
}
