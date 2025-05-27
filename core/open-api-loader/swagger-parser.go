package open_api_loader

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"os"
)

func ParseSwaggerFile(filePath string) (summary string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("cannot create v2 model")
			summary = ""
		}
	}()

	swaggerSpec, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot create new document: %w", err)
	}

	document, err := libopenapi.NewDocument(swaggerSpec)
	if err != nil {
		return "", fmt.Errorf("cannot create new document: %w", err)
	}

	v2Model, errors := document.BuildV2Model()
	if len(errors) > 0 {
		return "", fmt.Errorf("cannot create v2 model from document: %d errors reported", len(errors))
	}

	paths := v2Model.Model.Paths.PathItems.Len()
	schemas := v2Model.Model.Definitions.Definitions.Len()

	// Accessing possibly-nil fields, but protected by recover()
	summary = fmt.Sprintf(
		"There are %d paths and %d schemas in the document\nSwagger version: %s\nSwagger info: %s\nSwagger description: %s\nSwagger contact: %s\nSwagger license: %s\nSwagger version: %s\nSwagger terms of service: %s\nSwagger host: %s\nSwagger base path: %s\nSwagger schemes: %v\nSwagger produces: %v\nSwagger consumes: %v\nSwagger tags: %v\nSwagger security: %v\nSwagger external docs: %s\n",
		paths,
		schemas,
		v2Model.Model.Swagger,
		v2Model.Model.Info.Title,
		v2Model.Model.Info.Description,
		v2Model.Model.Info.Contact.Name,
		v2Model.Model.Info.License.Name,
		v2Model.Model.Info.Version,
		v2Model.Model.Info.TermsOfService,
		v2Model.Model.Host,
		v2Model.Model.BasePath,
		v2Model.Model.Schemes,
		v2Model.Model.Produces,
		v2Model.Model.Consumes,
		v2Model.Model.Tags,
		v2Model.Model.Security,
		v2Model.Model.ExternalDocs.Description,
	)

	return summary, nil
}
