package open_api_loader

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/v2"
	"os"
)

func ParseSwaggerFile(filePath string) {
	// load a swagger specification from bytes
	swaggerSpec, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	// create a new document from specification bytes
	document, err := libopenapi.NewDocument(swaggerSpec)
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	// define variables to capture the v2 model, or any errors thrown
	var errors []error
	var v2Model *libopenapi.DocumentModel[v2.Swagger]

	// because we know this is a v2 spec,
	// we can build a ready to go model from it.
	v2Model, errors = document.BuildV2Model()

	// if anything went wrong when building the v2 model,
	// a slice of errors will be returned
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from "+
			"document: %d errors reported", len(errors)))
	}

	// get a count of the number of paths and schemas.
	paths := v2Model.Model.Paths.PathItems.Len()
	schemas := v2Model.Model.Definitions.Definitions.Len()

	// print the number of paths and schemas in the document
	fmt.Printf("There are %d paths and %d schemas"+
		" in the document", paths, schemas)

	// print the swagger version
	fmt.Printf("Swagger version: %s\n", v2Model.Model.Swagger)
	// print the swagger info
	fmt.Printf("Swagger info: %s\n", v2Model.Model.Info.Title)
	// print the swagger description
	fmt.Printf("Swagger description: %s\n", v2Model.Model.Info.Description)
	// print the swagger contact
	fmt.Printf("Swagger contact: %s\n", v2Model.Model.Info.Contact.Name)
	// print the swagger license
	fmt.Printf("Swagger license: %s\n", v2Model.Model.Info.License.Name)
	// print the swagger version
	fmt.Printf("Swagger version: %s\n", v2Model.Model.Info.Version)
	// print the swagger terms of service
	fmt.Printf("Swagger terms of service: %s\n", v2Model.Model.Info.TermsOfService)
	// print the swagger host
	fmt.Printf("Swagger host: %s\n", v2Model.Model.Host)
	// print the swagger base path
	fmt.Printf("Swagger base path: %s\n", v2Model.Model.BasePath)
	// print the swagger schemes
	fmt.Printf("Swagger schemes: %s\n", v2Model.Model.Schemes)
	// print the swagger produces
	fmt.Printf("Swagger produces: %s\n", v2Model.Model.Produces)
	// print the swagger consumes
	fmt.Printf("Swagger consumes: %s\n", v2Model.Model.Consumes)
	// print the swagger tags
	fmt.Printf("Swagger tags: %v\n", v2Model.Model.Tags)
	// print the swagger security
	fmt.Printf("Swagger security: %v\n", v2Model.Model.Security)
	// print the swagger external docs
	fmt.Printf("Swagger external docs: %s\n", v2Model.Model.ExternalDocs.Description)
}
