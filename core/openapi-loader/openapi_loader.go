package openapiloader

// OpenAPILoader defines methods for loading and normalizing OpenAPI v2/v3 specs.
type OpenAPILoader interface {
	LoadSpec(path string) (*UnifiedAPISpec, error)
}

// UnifiedAPISpec is a normalized representation of an OpenAPI spec for downstream modules.
type UnifiedAPISpec struct {
	Version    string
	Endpoints  []APIEndpoint
	Schemas    map[string]interface{}
	Security   map[string]interface{}
}

type APIEndpoint struct {
	Path       string
	Method     string
	Operation  string
	Parameters []Parameter
}

type Parameter struct {
	Name     string
	In       string
	Required bool
	Schema   interface{}
}
