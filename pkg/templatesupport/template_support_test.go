package templatesupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := templatesupport.Parse(resourcesLocation(), "test", struct {
		Name string
	}{Name: "world"})
	assert.Equal(t, "hi world", parsed)
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

func TestParse_missingResources(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	_ = templatesupport.Parse("_x_", "test", struct {
		Name string
	}{Name: "world"})
}

func resourcesLocation() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(file, "../resources")
}
