package googlecloudsupport

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func EnableServicesCmd(resourcesDirectory string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_clusters_enable_services", nil)}
}

func ConfigureCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_clusters_configure", envMap)}
}

func CreateClustersCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_clusters_create", envMap)}
}

func ListClustersCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_clusters_list", envMap)}
}

func DeleteClustersCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_clusters_delete", envMap)}
}

func CreateServiceAccountCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "google_clusters_create_service_account", envMap)}
}
