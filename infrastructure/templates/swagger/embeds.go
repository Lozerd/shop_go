package templates

import _ "embed"

//go:embed swagger_index.tmpl
var SwaggerIndexTpl string

//go:embed swagger.json
var SwaggerJsonTpl string
