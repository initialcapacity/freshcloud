package googlecloudsupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/googlecloudsupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestEnableServices(t *testing.T) {
	services := googlecloudsupport.EnableServicesCmd()
	assert.Contains(t, services, "gcloud services enable container.googleapis.com --quiet\n")
}

func TestCreateClusterCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := googlecloudsupport.CreateClustersCmd(resourcesDirectory, "aProject", "aZone --quiet", "aClusterName")
	assert.Contains(t, clusterCmd, "gcloud beta container --project \"aProject\"")
	assert.Contains(t, clusterCmd, "clusters create \"aClusterName\"")
}

func TestConfigureCmd(t *testing.T) {
	cmd := googlecloudsupport.ConfigureCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters get-credentials 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd)
}

func TestDeleteClustersCmdCmd(t *testing.T) {
	cmd := googlecloudsupport.DeleteClustersCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters delete 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd)
}

func TestListClustersCmdCmd(t *testing.T) {
	cmd := googlecloudsupport.ListClustersCmd("aProject", "aZone")
	assert.Equal(t, "gcloud container clusters list --project 'aProject' --zone 'aZone' --quiet", cmd)
}
