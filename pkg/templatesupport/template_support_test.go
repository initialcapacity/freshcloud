package templatesupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParse(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../resources")
	parsed := templatesupport.Parse(resourcesDirectory, "test", struct {
		Name string
	}{Name: "world"})
	assert.Equal(t, "hi world", parsed)
}

func TestParse_missingResources(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../x_resources")
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	_ = templatesupport.Parse(resourcesDirectory, "test", struct {
		Name string
	}{Name: "world"})
}
