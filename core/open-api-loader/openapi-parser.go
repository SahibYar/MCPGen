package open_api_loader

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"os"
)

func ParseOpenAPISpecsFile(filePath string) (summary string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("cannot create v3 model")
			summary = ""
		}
	}()
	openApiSpecs, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read file: %w", err)
	}

	document, err := libopenapi.NewDocument(openApiSpecs)
	if err != nil {
		return "", fmt.Errorf("cannot create document: %w", err)
	}

	v3Model, errors := document.BuildV3Model()
	if len(errors) > 0 {
		return "", fmt.Errorf("cannot create v3 model: %v", errors)
	}

	paths := v3Model.Model.Paths.PathItems.Len()
	schemas := v3Model.Model.Components.Schemas.Len()

	summary = fmt.Sprintf(
		"There are %d paths and %d schemas in the document\nOpenAPI version: %s\nOpenAPI info: %s\nOpenAPI description: %s\nOpenAPI contact: %s\nOpenAPI license: %s\n",
		paths,
		schemas,
		v3Model.Model.Info.Version,
		v3Model.Model.Info.Title,
		v3Model.Model.Info.Description,
		v3Model.Model.Info.Contact.Name,
		v3Model.Model.Info.License.Name,
	)

	return summary, nil
}
