package templatesupport

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

var Client HTTPClient

func Parse(resourcesLocation, name string, data interface{}) string {
	var path string
	if strings.HasPrefix(resourcesLocation, "https://raw.githubusercontent.com/") {
		get, getErr := Client.Get(fmt.Sprintf("%s/%s.sh", resourcesLocation, name))
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(get.Body)

		if getErr != nil || get.StatusCode != 200 {
			panic("unable to get desired resource.")
		}
		path = filepath.Join(os.TempDir(), fmt.Sprintf("%s.sh", name))

		all, _ := io.ReadAll(get.Body)
		file, _ := os.Create(path)
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		_, _ = fmt.Fprintf(file, string(all))
	} else {
		path = filepath.Join(resourcesLocation, "./"+name+".sh")
	}

	var parsed bytes.Buffer
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
