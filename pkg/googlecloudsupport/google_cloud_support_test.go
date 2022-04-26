package googlecloudsupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/googlecloudsupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestEnableServices(t *testing.T) {
	services := googlecloudsupport.EnableServicesCmd()
	assert.Contains(t, services, "gcloud services enable container.googleapis.com")
}

func TestCreateClusterCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../freshctl/resources")
	clusterCmd := googlecloudsupport.CreateClusterCmd(resourcesDirectory, "aProject", "aZone", "aClusterName")
	assert.Contains(t, clusterCmd, "gcloud beta container --project \"aProject\"")
	assert.Contains(t, clusterCmd, "clusters create \"aClusterName\"")
}

func TestConfigureCmd(t *testing.T) {
	cmd := googlecloudsupport.ConfigureCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters get-credentials 'aClusterName' --zone 'aZone' --project 'aProject'\n", cmd)
}
