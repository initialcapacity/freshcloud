package templatesupport_test

import (
	"errors"
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := templatesupport.Parse(resourcesLocation(), "test", struct {
		Name string
	}{Name: "world"})
	assert.Equal(t, "hi world", parsed)
}

func TestParseViaUrl(t *testing.T) {
	templatesupport.Client = &http.Client{}
	location := "https://raw.githubusercontent.com/initialcapacity/freshcloud/main/pkg/templatesupport/resources"
	parsed := templatesupport.Parse(location, "test", struct {
		Name string
	}{Name: "world"})
	assert.Equal(t, "hi world", parsed)
}

type mock struct {
	err error
}

func (m mock) Get(url string) (resp *http.Response, err error) {
	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(""))}, m.err
}

func TestParseViaUrl_GetFails(t *testing.T) {
	templatesupport.Client = mock{err: errors.New("oops")}
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	location := "https://raw.githubusercontent.com/"
	_ = templatesupport.Parse(location, "test", struct {
		Name string
	}{Name: "world"})
}

func TestParse_badTemplate(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	_ = templatesupport.Parse(resourcesLocation(), "test", struct {
		Bad string
	}{Bad: "world"})
}

func xTestParse_missingResources(t *testing.T) { // bring this back for custom directories
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	_ = templatesupport.Parse("_x_", "testz", struct {
		Name string
	}{Name: "world"})
}

func resourcesLocation() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(file, "../resources")
}
