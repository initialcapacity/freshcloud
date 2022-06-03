package pipelines_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/pipelines"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPushPipelineImageCmd(t *testing.T) {
	pushCmd := pipelines.PushPipelineImageCmd(resourcesLocation(), map[string]string{
		"REGISTRY_DOMAIN":   "aRegistryDomain",
		"REGISTRY_PASSWORD": "aPassword",
	})
	expected := `docker build --platform linux/amd64 --rm -t registry.aRegistryDomain/concourse-images/cluster-mgmt:latest \
  https://raw.githubusercontent.com/initialcapacity/freshcloud/main/pkg/freshctl/resources/Dockerfile
docker login -u admin -p aPassword https://registry.aRegistryDomain
docker push registry.aRegistryDomain/concourse-images/cluster-mgmt:latest`
	assert.Equal(t, expected, pushCmd[0])
}

func TestDeployPipelineCmd(t *testing.T) {
	deployCmd := pipelines.DeployPipelineCmd(resourcesLocation(), map[string]string{
		"REGISTRY_DOMAIN":                 "aDomain",
		"REGISTRY_PASSWORD":               "aPassword",
		"REGISTRY_CLUSTER_NAME":           "aCluster",
		"APP_NAME":                        "anApp",
		"APP_PIPELINE_CONFIGURATION_PATH": "aPath",
		"APP_PIPELINE_PATH":               "anotherPath",
	})
	assert.Contains(t, deployCmd[0], `CA=$(kubectl get secret/${NAME} -n anApp -o jsonpath='{.data.ca\.crt}')`)
}

func TestDeletePipelineCmd(t *testing.T) {
	deleteCmd := pipelines.DeletePipelineCmd(resourcesLocation(), map[string]string{
		"REGISTRY_DOMAIN":       "aDomain",
		"REGISTRY_PASSWORD":     "aPassword",
		"REGISTRY_CLUSTER_NAME": "aCluster",
		"APP_NAME":              "anApp",
	})
	expected := `fly login -c https://ci.aDomain -u admin -p aPassword -t aCluster
echo y | fly -t aCluster dp -p build-anApp
kubectl delete ns anApp`
	assert.Equal(t, expected, deleteCmd[0])
}

func resourcesLocation() string {
	_, file, _, _ := runtime.Caller(0)
	resourcesLocation := filepath.Join(file, "../../resources")
	return resourcesLocation
}
