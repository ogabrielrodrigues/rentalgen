package templates

import "embed"

//go:embed *.html styles.css
var FS embed.FS
