package arazzo_parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const validSpec = `
info:
  title: "Test"
  version: "1.0"
workflows:
  - workflowID: "wf1"
    steps:
      - id: "step1"
        type: "task"
`

const invalidSpec = `
info:
  title: "Invalid"
  version: "1.0"
workflows:
  - workflowID: "wf1"
    steps:
      - id: "step1"
        type: "invalid_type"
`

func createTempSpec(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "spec.yaml")
	require.NoError(t, os.WriteFile(path, []byte(content), 0644))
	return path
}

func TestRead_ValidSpec(t *testing.T) {
	file := createTempSpec(t, validSpec)
	output, validationErrs, err := Read(file)
	require.NoError(t, err)
	require.Empty(t, validationErrs)
	require.Contains(t, output, "Speakeasy Bar Workflows")
}

func TestRead_InvalidPath(t *testing.T) {
	_, _, err := Read("nonexistent.yaml")
	require.Error(t, err)
}

func TestValidate_ValidSpec(t *testing.T) {
	file := createTempSpec(t, validSpec)
	valid, validationErrs, err := Validate(file)
	require.NoError(t, err)
	require.True(t, valid)
	require.Empty(t, validationErrs)
}

func TestValidate_InvalidPath(t *testing.T) {
	_, _, err := Validate("missing.yaml")
	require.Error(t, err)
}

func TestValidate_InvalidSpec(t *testing.T) {
	file := createTempSpec(t, invalidSpec)
	valid, validationErrs, err := Validate(file)
	require.NoError(t, err)
	require.False(t, valid)
	require.NotEmpty(t, validationErrs)
}

func TestWalk_ValidSpec(t *testing.T) {
	file := createTempSpec(t, validSpec)
	ids, err := Walk(file)
	require.NoError(t, err)
	require.Contains(t, ids, "wf1")
}

func TestWalk_InvalidPath(t *testing.T) {
	_, err := Walk("notfound.yaml")
	require.Error(t, err)
}

func TestWalk_InvalidSpec(t *testing.T) {
	file := createTempSpec(t, "bad yaml: :::")
	_, err := Walk(file)
	require.Error(t, err)
}

func TestRead_InvalidSpec(t *testing.T) {
	file := createTempSpec(t, "bad: : yaml")
	_, _, err := Read(file)
	require.Error(t, err)
}

func TestRead_OutputFormat(t *testing.T) {
	file := createTempSpec(t, validSpec)
	output, _, err := Read(file)
	require.NoError(t, err)
	require.True(t, strings.Contains(output, "Speakeasy Bar Workflows"))
}
