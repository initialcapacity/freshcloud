package googlecloudsupport_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/googlecloudsupport"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestEnableServices(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	servicesCmd := googlecloudsupport.EnableServicesCmd(resourcesDirectory)
	assert.Contains(t, servicesCmd[0], "gcloud services enable container.googleapis.com --quiet")
}

func TestCreateClusterCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	env := map[string]string{
		"GCP_PROJECT_ID":   "aProject",
		"GCP_ZONE":         "aZone",
		"GCP_CLUSTER_NAME": "aClusterName",
	}
	clusterCmd := googlecloudsupport.CreateClustersCmd(resourcesDirectory, env)
	expected := `gcloud container clusters create aClusterName --zone aZone --num-nodes 4`
	assert.Equal(t, clusterCmd[0], expected)
}

func TestListClustersCmdCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	env := map[string]string{
		"GCP_PROJECT_ID": "aProject",
		"GCP_ZONE":       "aZone",
	}
	cmd := googlecloudsupport.ListClustersCmd(resourcesDirectory, env)
	assert.Equal(t, "gcloud container clusters list --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}

func TestConfigureCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	env := map[string]string{
		"GCP_PROJECT_ID":   "aProject",
		"GCP_ZONE":         "aZone",
		"GCP_CLUSTER_NAME": "aClusterName",
	}
	clusterCmd := googlecloudsupport.ConfigureCmd(resourcesDirectory, env)
	assert.Equal(t, "gcloud container clusters get-credentials 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", clusterCmd[0])
}

func TestDeleteClustersCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	env := map[string]string{
		"GCP_PROJECT_ID":   "aProject",
		"GCP_ZONE":         "aZone",
		"GCP_CLUSTER_NAME": "aClusterName",
	}
	cmd := googlecloudsupport.DeleteClustersCmd(resourcesDirectory, env)
	assert.Equal(t, "gcloud container clusters delete 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}

func TestCreateServiceAccountCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	env := map[string]string{
		"GCP_PROJECT_ID": "aProject",
	}
	cmd := googlecloudsupport.CreateServiceAccountCmd(resourcesDirectory, env)
	assert.Equal(t, `mkdir -p .freshcloud
gcloud iam service-accounts create aProject-svc --display-name=aProject-svc
gcloud iam service-accounts keys create .freshcloud/aProject-svc.json \
  --iam-account=aProject-svc@aProject.iam.gserviceaccount.com
gcloud projects add-iam-policy-binding aProject \
  --member=serviceAccount:aProject-svc@aProject.iam.gserviceaccount.com \
  --role=roles/owner`, cmd[0])
}
