package googlecloudsupport

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func EnableServicesCmd(resourcesDirectory string) []string {
	templateName := "google_cloud_services"
	return []string{
		templatesupport.Parse(resourcesDirectory, templateName, nil),
	}
}

func ConfigureCmd(projectId, zone, clusterName string) []string {
	return []string{
		fmt.Sprintf("gcloud container clusters get-credentials '%s' --project '%v' --zone '%v' --quiet", clusterName, projectId, zone),
	}
}

func CreateClustersCmd(resourcesDirectory, projectId, zone, clusterName string) []string {
	templateName := "google_cloud_cluster"
	data := struct {
		ProjectID   string
		Zone        string
		ClusterName string
	}{
		ProjectID:   projectId,
		Zone:        zone,
		ClusterName: clusterName,
	}
	return []string{
		templatesupport.Parse(resourcesDirectory, templateName, data),
	}
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
