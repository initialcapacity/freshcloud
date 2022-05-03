package applications_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/applications"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/cmds"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"
)

func TestPushImageCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := applications.PushImageCmd(resourcesDirectory, map[string]string{
		"REGISTRY_DOMAIN":   "aRegistryDomain",
		"REGISTRY_PASSWORD": "aPassword",
		"APP_NAME":          "anApp",
		"APP_IMAGE_NAME":    "anImage",
	})
	expected := `curl --user "admin:aPassword" -X POST \
  https://registry.aRegistryDomain/api/v2.0/projects \
  -H "Content-type: application/json" --data \
  '{ "project_name": "'anApp'", "metadata":
   { "auto_scan": "true", "enable_content_trust":
     "false", "prevent_vul": "false", "public":
     "true", "reuse_sys_cve_whitelist": "true",
     "severity": "high" }
   }'
docker login -u admin -p aPassword https://registry.aRegistryDomain
docker tag anImage registry.aRegistryDomain/anApp/anImage:latest
docker push registry.aRegistryDomain/anApp/anImage:latest`
	assert.Equal(t, expected, clusterCmd[0])
}

func TestDeployAppCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")

	_ = syscall.Setenv("REGISTRY_DOMAIN", "aRegistryDomain")
	_ = syscall.Setenv("REGISTRY_PASSWORD", "aRegistryPassword")
	_ = syscall.Setenv("APP_NAME", "anApp")
	_ = syscall.Setenv("DOMAIN", "anAppDomain")
	_ = syscall.Setenv("APP_IMAGE_NAME", "anImage")
	_ = syscall.Setenv("APP_CONFIGURATION_PATH", "aPath")

	clusterCmd := applications.DeployAppCmd(resourcesDirectory, cmds.MakeEnvironmentMap(os.Environ()))

	expected := `mkdir -p .freshcloud
sha=$(curl -s --user "admin:aRegistryPassword" -X GET \
  "https://registry.aRegistryDomain/api/v2.0/projects/anApp/repositories/anImage/artifacts" \
  | jq -r '.[].digest'|head -1)
export IMAGE="registry.aRegistryDomain/anApp/anImage@${sha}"
echo "Found image ${IMAGE}"
kubectl create namespace anApp
envsubst < aPath > .freshcloud/anApp.yaml
kubectl apply -f .freshcloud/anApp.yaml
echo "Deploy anApp to https://anApp.anAppDomain"
echo "Remove the app by running - kubectl delete ns anApp"`
	assert.Equal(t, expected, clusterCmd[0])
}
