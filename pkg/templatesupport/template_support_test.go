package templatesupport_test

import (
	"errors"
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := templatesupport.Parse("", "test", struct {
		Name string
	}{Name: "world"})
	assert.Equal(t, "hi world", parsed)
}

func TestParseViaUrl(t *testing.T) {
	templatesupport.Client = &http.Client{}
	location := "https://raw.githubusercontent.com/initialcapacity/freshcloud/main/pkg/freshctl/resources"
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

func TestParseViaLocation(t *testing.T) {
	tempDir := os.TempDir()
	_ = ioutil.WriteFile(filepath.Join(tempDir, "test.sh"), []byte("hi {{.Name}}"), 0644)

	resp := templatesupport.Parse(tempDir, "test", struct {
		Name string
	}{Name: "world"})
	assert.Equal(t, "hi world", string(resp))
}

func TestParseViaLocation_missingResources(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()

	tempDir := os.TempDir()
	_ = ioutil.WriteFile(filepath.Join(tempDir, "test.sh"), []byte("hi {{.Name}}"), 0644)

	_ = templatesupport.Parse(tempDir, "tes", struct {
		Name string
	}{Name: "world"})
}

func TestParse_badTemplate(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	_ = templatesupport.Parse("", "test", struct {
		Bad string
	}{Bad: "world"})
}
