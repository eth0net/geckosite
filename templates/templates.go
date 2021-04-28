package templates

import (
	"embed"
	"html/template"
)

//go:embed *.gohtml
var fs embed.FS

// Parse takes a template name and returns the parsed template from the embedded FS.
func Parse(name string) *template.Template {
	return template.Must(template.ParseFS(fs, "layout.gohtml", name+".gohtml"))
}

// ParseInto takes an existing Template pointer and parses
// the specified template from the embedded FS into it.
func ParseInto(t *template.Template, name string) *template.Template {
	return template.Must(t.ParseFS(fs, "layout.gohtml", name+".gohtml"))
}
