package templatesupport

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func Parse(resourcesDirectory, name string, data any) string {
	var parsed bytes.Buffer
	path := filepath.Join(resourcesDirectory, "./"+name+".sh")
	tmpl, err := template.New(filepath.Base(path)).Funcs(template.FuncMap{}).ParseFiles(path)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&parsed, data)
	if err != nil {
		panic(err)
	}
	return parsed.String()
}
