package open_api_loader

import (
	"MCPGen/core/utils"
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
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

	summary = fmt.Sprintf(
		"There are %d paths and %d schemas in the document\nSwagger version: %s\nSwagger info: %s\n",
		paths,
		schemas,
		v2Model.Model.Swagger,
		v2Model.Model.Info.Title,
	)

	// Optional fields
	var (
		description  = utils.SafeStr(v2Model.Model.Info.Description)
		contact      = utils.SafeStrPtr(v2Model.Model.Info.Contact, func(c *base.Contact) string { return c.Name })
		license      = utils.SafeStrPtr(v2Model.Model.Info.License, func(l *base.License) string { return l.Name })
		version      = utils.SafeStr(v2Model.Model.Info.Version)
		terms        = utils.SafeStr(v2Model.Model.Info.TermsOfService)
		host         = utils.SafeStr(v2Model.Model.Host)
		basePath     = utils.SafeStr(v2Model.Model.BasePath)
		schemes      = utils.SafeSlice(v2Model.Model.Schemes)
		produces     = utils.SafeSlice(v2Model.Model.Produces)
		consumes     = utils.SafeSlice(v2Model.Model.Consumes)
		tags         = utils.SafeSlice(v2Model.Model.Tags)
		security     = utils.SafeSlice(v2Model.Model.Security)
		externalDocs = utils.SafeStrPtr(v2Model.Model.ExternalDocs, func(e *base.ExternalDoc) string { return e.Description })
	)

	summary += fmt.Sprintf(
		"Swagger description: %s\nSwagger contact: %s\nSwagger license: %s\nSwagger version: %s\nSwagger terms of service: %s\nSwagger host: %s\nSwagger base path: %s\nSwagger schemes: %v\nSwagger produces: %v\nSwagger consumes: %v\nSwagger tags: %v\nSwagger security: %v\nSwagger external docs: %s\n",
		description,
		contact,
		license,
		version,
		terms,
		host,
		basePath,
		schemes,
		produces,
		consumes,
		tags,
		security,
		externalDocs,
	)
	return summary, nil
}
