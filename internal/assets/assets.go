package assets

import "embed"

//go:embed templates

//go:embed templates/.*.tmpl
var TemplatesFS embed.FS
