package arazzo_parser

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/speakeasy-api/openapi/arazzo"
)

func Read(filePath string) {
	ctx := context.Background()
	r, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// Unmarshal the Arazzo document which will also validate it against the Arazzo Specification
	a, validationErrs, err := arazzo.Unmarshal(ctx, r)
	if err != nil {
		panic(err)
	}

	// Validation errors are returned separately from any errors that block the document from being unmarshalled
	// allowing an invalid document to be mutated and fixed before being marshalled again
	for _, err := range validationErrs {
		fmt.Println(err.Error())
	}

	// Mutate the document by just modifying the returned Arazzo object
	a.Info.Title = "Speakeasy Bar Workflows"

	buf := bytes.NewBuffer([]byte{})

	// Marshal the document to a writer
	if err := arazzo.Marshal(ctx, a, buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

func Walk(filePath string) {
	ctx := context.Background()

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	a, _, err := arazzo.Unmarshal(ctx, f)
	if err != nil {
		panic(err)
	}

	err = arazzo.Walk(ctx, a, func(ctx context.Context, node, parent arazzo.MatchFunc, a *arazzo.Arazzo) error {
		return node(arazzo.Matcher{
			Workflow: func(workflow *arazzo.Workflow) error {
				fmt.Printf("Workflow: %s\n", workflow.WorkflowID)
				return nil
			},
		})
	})
	if err != nil {
		panic(err)
	}
}

func Validate(filePath string) {
	ctx := context.Background()
	r, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// Unmarshal the Arazzo document which will also validate it against the Arazzo Specification
	a, validationErrs, err := arazzo.Unmarshal(ctx, r)
	if err != nil {
		panic(err)
	}

	// Validation errors are returned separately from any errors that block the document from being unmarshalled
	// allowing an invalid document to be mutated and fixed before being marshalled again
	for _, err := range validationErrs {
		fmt.Println(err.Error())
	}

	fmt.Printf("Arazzo document is valid: %v\n", a.Valid)
}
