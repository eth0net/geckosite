package templates

import (
	"embed"
	"fmt"
	"html/template"
	"strconv"
	"time"
)

//go:embed *.gohtml
var fs embed.FS

var baseTemplateFuncMap = template.FuncMap{
	"copyright": func(holder string, start int) string {
		var year string
		now := time.Now().Year()
		if start > 0 && start < now {
			year += strconv.Itoa(start) + "-"
		}
		year += strconv.Itoa(now)

		return fmt.Sprintf("Ⓒ Copyright %s %s", year, holder)
	},
}

// Parse takes a template name and returns the parsed template from the embedded FS.
func Parse(name string) *template.Template {
	t := template.New(name)
	return ParseInto(t, name)
}

// ParseInto takes an existing Template pointer and parses
// the specified template from the embedded FS into it.
func ParseInto(t *template.Template, name string) *template.Template {
	t = t.Funcs(baseTemplateFuncMap)
	return template.Must(t.ParseFS(fs, "layout.gohtml", name+".gohtml"))
}
