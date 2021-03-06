package templatesupport

import (
	"bytes"
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl"
	"io"
	"io/fs"
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
	var fileBytes []byte
	if strings.HasPrefix(resourcesLocation, "https://raw.githubusercontent.com/") {
		get, getErr := Client.Get(fmt.Sprintf("%s/%s.sh", resourcesLocation, name))
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(get.Body)

		if getErr != nil || get.StatusCode != 200 {
			panic("unable to get desired resource.")
		}
		fileBytes, _ = io.ReadAll(get.Body)
	} else if strings.HasPrefix(resourcesLocation, "/") {
		f := filepath.Join(resourcesLocation, name+".sh")
		b, _ := os.ReadFile(f)
		fileBytes = b
	} else {
		b, _ := fs.ReadFile(freshctl.Resources, "resources/"+name+".sh")
		fileBytes = b
	}

	if len(fileBytes) == 0 {
		panic("unable to find resources.")
	}

	path := filepath.Join(os.TempDir(), fmt.Sprintf("%s.sh", name))
	file, _ := os.Create(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = fmt.Fprintf(file, string(fileBytes))

	var parsed bytes.Buffer
	tmpl, _ := template.New(filepath.Base(path)).Funcs(template.FuncMap{}).ParseFiles(path)
	err := tmpl.Execute(&parsed, data)
	if err != nil {
		panic(err)
	}
	return parsed.String()
}
