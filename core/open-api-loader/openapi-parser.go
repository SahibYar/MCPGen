package open_api_loader

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"os"
)

func ParseOpenAPISpecsFile(filePath string) {
	// load an OpenAPI 3 specification from bytes
	openApiSpecs, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	// create a new document from specification bytes
	document, err := libopenapi.NewDocument(openApiSpecs)
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	// because we know this is a v3 spec,
	// we can build a ready to go model from it.
	v3Model, errors := document.BuildV3Model()

	// if anything went wrong when building the v3 model,
	// a slice of errors will be returned
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
	}

	// get a count of the number of paths and schemas.
	paths := v3Model.Model.Paths.PathItems.Len()
	schemas := v3Model.Model.Components.Schemas.Len()

	// print the number of paths and schemas in the document
	fmt.Printf("There are %d paths and %d schemas "+
		"in the document", paths, schemas)
	// print the OpenAPI version
	fmt.Printf("OpenAPI version: %s\n", v3Model.Model.Info.Version)
	// print the OpenAPI info
	fmt.Printf("OpenAPI info: %s\n", v3Model.Model.Info.Title)
	// print the OpenAPI description
	fmt.Printf("OpenAPI description: %s\n", v3Model.Model.Info.Description)
	// print the OpenAPI contact
	fmt.Printf("OpenAPI contact: %s\n", v3Model.Model.Info.Contact.Name)
	// print the OpenAPI license
	fmt.Printf("OpenAPI license: %s\n", v3Model.Model.Info.License.Name)
}
