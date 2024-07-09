package docs

import "embed"

//go:embed api.yaml
var OpenAPISpec embed.FS
