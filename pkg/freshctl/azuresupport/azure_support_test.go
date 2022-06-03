package azuresupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/azuresupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestConfigureCmd(t *testing.T) {
	env := map[string]string{
		"K8S_CLUSTER_NAME":     "aClusterName",
		"AZURE_RESOURCE_GROUP": "aResourceGroup",
	}
	clusterCmd := azuresupport.ConfigureCmd(resourcesLocation(), env)
	assert.Equal(t, "az aks get-credentials --name aClusterName --overwrite-existing --resource-group aResourceGroup", clusterCmd[0])
}

func TestDeleteClustersCmd(t *testing.T) {
	env := map[string]string{
		"K8S_CLUSTER_NAME":     "aClusterName",
		"AZURE_RESOURCE_GROUP": "aResourceGroup",
	}
	cmd := azuresupport.CreateClustersCmd(resourcesLocation(), env)
	assert.Equal(t, "az aks create --resource-group aResourceGroup --name aClusterName \\\n  --node-count 1 --generate-ssh-keys", cmd[0])
}

func TestCreateResourceGroupCmd(t *testing.T) {
	env := map[string]string{
		"AZURE_LOCATION":       "aLocation",
		"AZURE_RESOURCE_GROUP": "aResourceGroup",
	}
	cmd := azuresupport.CreateResourceGroupCmd(resourcesLocation(), env)
	assert.Equal(t, "az group create --name aResourceGroup --location aLocation", cmd[0])
}

func TestDeleteResourceGroupCmd(t *testing.T) {
	env := map[string]string{
		"AZURE_RESOURCE_GROUP": "aResourceGroup",
	}
	cmd := azuresupport.DeleteResourceGroupCmd(resourcesLocation(), env)
	assert.Equal(t, "az group delete --name 'aResourceGroup' --yes --no-wait", cmd[0])
}

func resourcesLocation() string {
	_, file, _, _ := runtime.Caller(0)
	resourcesLocation := filepath.Join(file, "../../resources")
	return resourcesLocation
}
