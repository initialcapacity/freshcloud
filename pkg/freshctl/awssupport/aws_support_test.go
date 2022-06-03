package awssupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/awssupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestConfigureCmd(t *testing.T) {
	env := map[string]string{
		"AWS_CLUSTER_NAME": "aClusterName",
		"AWS_REGION":       "aRegion",
	}
	clusterCmd := awssupport.ConfigureCmd(resourcesLocation(), env)
	assert.Equal(t, "aws eks --region aRegion update-kubeconfig --name aClusterName", clusterCmd[0])
}

func TestCreateClustersCmd(t *testing.T) {
	env := map[string]string{
		"AWS_CLUSTER_NAME": "aClusterName",
		"AWS_REGION":       "aRegion",
	}
	cmd := awssupport.CreateClustersCmd(resourcesLocation(), env)
	assert.Equal(t, `eksctl create cluster -f -
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: aClusterName
  region: aRegion
nodeGroups:
  - name: controllers
    labels: { role: controllers }
    instanceTypes:
    - "t3.small"
    desiredCapacity: 2
    iam:
      withAddonPolicies:
        certManager: true
        albIngress: true
    taints:
      controllers: "true:NoSchedule"
managedNodeGroups:
  - name: workers`, cmd[0])
}

func TestDeleteClustersCmd(t *testing.T) {
	env := map[string]string{
		"AWS_CLUSTER_NAME": "aClusterName",
		"AWS_REGION":       "aRegion",
	}
	cmd := awssupport.DeleteClustersCmd(resourcesLocation(), env)
	assert.Equal(t, "eksctl delete cluster --name aClusterName --region aRegion", cmd[0])
}

func resourcesLocation() string {
	_, file, _, _ := runtime.Caller(0)
	resourcesLocation := filepath.Join(file, "../../resources")
	return resourcesLocation
}
