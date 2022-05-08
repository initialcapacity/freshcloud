package pipelines_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/pipelines"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPushPipelineImageCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	pushCmd := pipelines.PushPipelineImageCmd(resourcesDirectory, map[string]string{
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
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	deployCmd := pipelines.DeployPipelineCmd(resourcesDirectory, map[string]string{
		"GCP_PROJECT_ID":                  "aProject",
		"GCP_ZONE":                        "aZone",
		"REGISTRY_DOMAIN":                 "aDomain",
		"REGISTRY_PASSWORD":               "aPassword",
		"REGISTRY_CLUSTER_NAME":           "aCluster",
		"APP_NAME":                        "anApp",
		"APP_PIPELINE_CONFIGURATION_PATH": "aPath",
		"APP_PIPELINE_PATH":               "anotherPath",
	})
	assert.Contains(t, deployCmd[0], `gcloud container clusters get-credentials 'aCluster'`)
	assert.Contains(t, deployCmd[0], `CA=$(kubectl get secret/${NAME} -n anApp -o jsonpath='{.data.ca\.crt}')`)
}

func TestDeletePipelineCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	deleteCmd := pipelines.DeletePipelineCmd(resourcesDirectory, map[string]string{
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
