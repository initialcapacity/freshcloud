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
	parsed := templatesupport.Parse("", "version", struct {
		Version string
	}{Version: "1.0"})
	assert.Equal(t, "freshctl version 1.0", parsed)
}

func TestParseViaUrl(t *testing.T) {
	templatesupport.Client = mock{body: "freshctl version 2.0"}
	location := "https://raw.githubusercontent.com/initialcapacity/freshcloud/main/pkg/freshctl/resources"
	parsed := templatesupport.Parse(location, "version", struct {
		Version string
	}{Version: "2.0"})
	assert.Equal(t, "freshctl version 2.0", parsed)
}

type mock struct {
	err  error
	body string
}

func (m mock) Get(url string) (resp *http.Response, err error) {
	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(m.body))}, m.err
}

func TestParseViaUrl_GetFails(t *testing.T) {
	templatesupport.Client = mock{err: errors.New("oops")}
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	location := "https://raw.githubusercontent.com/"
	_ = templatesupport.Parse(location, "version", struct {
		Version string
	}{Version: "3.4"})
}

func TestParseViaLocation(t *testing.T) {
	tempDir := os.TempDir()
	_ = ioutil.WriteFile(filepath.Join(tempDir, "version.sh"), []byte("freshctl version {{index .Version}}"), 0644)

	resp := templatesupport.Parse(tempDir, "version", struct {
		Version string
	}{Version: "4.2"})
	assert.Equal(t, "freshctl version 4.2", string(resp))
}

func TestParseViaLocation_missingResources(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()

	tempDir := os.TempDir()
	_ = ioutil.WriteFile(filepath.Join(tempDir, "version.sh"), []byte("freshctl version {{index .Version}}"), 0644)

	_ = templatesupport.Parse(tempDir, "ver", struct {
		Version string
	}{Version: "6.1"})
}

func TestParse_badTemplate(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	_ = templatesupport.Parse("", "version", struct {
		Bad string
	}{Bad: "1.0"})
}
