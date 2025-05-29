package arazzo_parser

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/speakeasy-api/openapi/arazzo"
)

func Read(filePath string) (string, []error, error) {
	ctx := context.Background()
	r, err := os.Open(filePath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer r.Close()

	a, validationErrs, err := arazzo.Unmarshal(ctx, r)
	if err != nil {
		return "", nil, fmt.Errorf("unmarshal error: %w", err)
	}

	// Mutate the document
	a.Info.Title = "Speakeasy Bar Workflows"

	buf := bytes.NewBuffer([]byte{})
	if err := arazzo.Marshal(ctx, a, buf); err != nil {
		return "", nil, fmt.Errorf("marshal error: %w", err)
	}

	return buf.String(), validationErrs, nil
}

func Walk(filePath string) ([]string, error) {
	ctx := context.Background()
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	a, _, err := arazzo.Unmarshal(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	var workflowIDs []string
	err = arazzo.Walk(ctx, a, func(ctx context.Context, node, parent arazzo.MatchFunc, a *arazzo.Arazzo) error {
		return node(arazzo.Matcher{
			Workflow: func(workflow *arazzo.Workflow) error {
				workflowIDs = append(workflowIDs, workflow.WorkflowID)
				return nil
			},
		})
	})
	if err != nil {
		return nil, fmt.Errorf("walk error: %w", err)
	}

	return workflowIDs, nil
}

func Validate(filePath string) (bool, []error, error) {
	ctx := context.Background()
	r, err := os.Open(filePath)
	if err != nil {
		return false, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer r.Close()

	a, validationErrs, err := arazzo.Unmarshal(ctx, r)
	if err != nil {
		return false, nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return a.Valid, validationErrs, nil
}
