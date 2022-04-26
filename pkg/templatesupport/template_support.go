package templatesupport

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func Parse(resourcesDirectory, name string, data any) string {
	var parsed bytes.Buffer
	tmpl := filepath.Join(resourcesDirectory, "./"+name+".gotmpl")
	err := template.Must(template.New(filepath.Base(tmpl)).Funcs(
		template.FuncMap{}).ParseFiles(tmpl)).Execute(&parsed, data)
	if err != nil {
		panic(err)
	}
	return parsed.String()
}
