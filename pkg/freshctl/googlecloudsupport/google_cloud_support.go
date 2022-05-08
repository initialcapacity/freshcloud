package googlecloudsupport

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func EnableServicesCmd(resourcesDirectory string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_cloud_services", nil)}
}

func ConfigureCmd(projectId, zone, clusterName string) []string {
	return []string{
		fmt.Sprintf("gcloud container clusters get-credentials '%s' --project '%v' --zone '%v' --quiet", clusterName, projectId, zone),
	}
}

func CreateClustersCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_cloud_cluster", envMap)}
}

func ListClustersCmd(projectId, zone string) []string {
	return []string{
		fmt.Sprintf("gcloud container clusters list --project '%v' --zone '%v' --quiet", projectId, zone),
	}
}

func DeleteClustersCmd(projectId, zone, clusterName string) []string {
	return []string{
		fmt.Sprintf("gcloud container clusters delete '%v' --project '%v' --zone '%v' --quiet", clusterName, projectId, zone),
	}
}

func CreateServiceAccountCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_cloud_service_account", envMap)}
}
