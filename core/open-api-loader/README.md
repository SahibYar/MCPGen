## Responsibilities:
*	Load and parse Swagger files
*	Extract endpoint details, schemas, security, parameters, etc.

# Methods:
```golang
type OpenAPILoader struct {
    SpecPath string
    Parsed   *OpenAPISpec
}

func (o *OpenAPILoader) Load() error
func (o *OpenAPILoader) GetEndpoints() []Endpoint
```