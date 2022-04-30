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
	clusterCmd := googlecloudsupport.CreateClustersCmd(resourcesDirectory, "aProject", "aZone", "aClusterName")
	expected := `gcloud beta container --project "aProject" \
    clusters create "aClusterName" \
        --zone "aZone" \
        --no-enable-basic-auth \
        --machine-type "n1-standard-2" \
        --image-type "COS" \
        --disk-type "pd-standard" \
        --disk-size "100" \
        --metadata disable-legacy-endpoints=true \
        --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
        --num-nodes "5" \
        --enable-stackdriver-kubernetes --enable-ip-alias \
        --network "projects/aProject/global/networks/default" \
        --subnetwork "projects/aProject/regions/us-central1/subnetworks/default" \
        --default-max-pods-per-node "110" \
        --no-enable-master-authorized-networks \
        --addons HorizontalPodAutoscaling,HttpLoadBalancing \
        --enable-autoupgrade \
        --enable-autorepair`
	assert.Equal(t, clusterCmd[0], expected)
}

func TestListClustersCmdCmd(t *testing.T) {
	cmd := googlecloudsupport.ListClustersCmd("aProject", "aZone")
	assert.Equal(t, "gcloud container clusters list --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}

func TestConfigureCmd(t *testing.T) {
	cmd := googlecloudsupport.ConfigureCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters get-credentials 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}

func TestDeleteClustersCmdCmd(t *testing.T) {
	cmd := googlecloudsupport.DeleteClustersCmd("aProject", "aZone", "aClusterName")
	assert.Equal(t, "gcloud container clusters delete 'aClusterName' --project 'aProject' --zone 'aZone' --quiet", cmd[0])
}
