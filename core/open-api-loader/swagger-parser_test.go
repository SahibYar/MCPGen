package open_api_loader

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseSwaggerFile_ValidFile(t *testing.T) {
	summary, err := ParseSwaggerFile("testdata/valid_swagger.yaml")
	assert.NoError(t, err)
	assert.Contains(t, summary, "There are")
	assert.Contains(t, summary, "Swagger version:")
	assert.Contains(t, summary, "Swagger info:")
}

func TestParseSwaggerFile_FileNotFound(t *testing.T) {
	summary, err := ParseSwaggerFile("testdata/non_existent.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create new document")
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_InvalidYAML(t *testing.T) {
	tmpFile := "testdata/invalid_yaml.yaml"
	_ = os.WriteFile(tmpFile, []byte("invalid: : yaml: : content"), 0644)
	defer os.Remove(tmpFile)

	summary, err := ParseSwaggerFile(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create new document")
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_ModelErrors(t *testing.T) {
	// This file has valid syntax but invalid Swagger structure
	summary, err := ParseSwaggerFile("testdata/model_error_swagger.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create v2 model")
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_MissingFieldsPanic(t *testing.T) {
	// This file is valid Swagger but has nil-able fields like Contact or License
	summary, err := ParseSwaggerFile("testdata/missing_fields_swagger.yaml")
	assert.Error(t, err)
	assert.Equal(t, "cannot create v2 model", err.Error())
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_EmptyFile(t *testing.T) {
	tmpFile := "testdata/empty.yaml"
	_ = os.WriteFile(tmpFile, []byte(""), 0644)
	defer os.Remove(tmpFile)

	summary, err := ParseSwaggerFile(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create new document")
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_NonSwaggerYAML(t *testing.T) {
	tmpFile := "testdata/non_swagger.yaml"
	_ = os.WriteFile(tmpFile, []byte("hello: world"), 0644)
	defer os.Remove(tmpFile)

	summary, err := ParseSwaggerFile(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "spec type not supported")
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_ValidMinimalSwagger(t *testing.T) {
	tmpFile := "testdata/minimal_swagger.yaml"
	_ = os.WriteFile(tmpFile, []byte(`
swagger: "2.0"
info:
  title: Minimal API
  version: "1.0.0"
paths:
  /ping:
    get:
      summary: Simple ping endpoint
      responses:
        200:
          description: pong
          schema:
            type: string
definitions: {}
`), 0644)
	defer os.Remove(tmpFile)

	summary, err := ParseSwaggerFile(tmpFile)
	assert.NoError(t, err)
	assert.Contains(t, summary, "paths and")
	assert.Contains(t, summary, "Swagger version")
}

func TestParseSwaggerFile_MissingPaths(t *testing.T) {
	tmpFile := "testdata/no_paths.yaml"
	_ = os.WriteFile(tmpFile, []byte(`
swagger: "2.0"
info:
  title: No Paths
  version: "1.0.0"
`), 0644)
	defer os.Remove(tmpFile)

	summary, err := ParseSwaggerFile(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create v3 model")
	assert.Empty(t, summary)
}

func TestParseSwaggerFile_MalformedPaths(t *testing.T) {
	tmpFile := "testdata/malformed_paths.yaml"
	_ = os.WriteFile(tmpFile, []byte(`
swagger: "2.0"
info:
  title: Broken Paths
  version: "1.0.0"
paths:
  /broken:
    get:
`), 0644)
	defer os.Remove(tmpFile)

	summary, err := ParseSwaggerFile(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot create v3 model")
	assert.Empty(t, summary)
}
