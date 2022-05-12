package googlecloudsupport

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func EnableServicesCmd(resourcesLocation string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "google_clusters_enable_services", nil)}
}

func ConfigureCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "google_clusters_configure", envMap)}
}

func CreateClustersCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "google_clusters_create", envMap)}
}

func ListClustersCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "google_clusters_list", envMap)}
}

func DeleteClustersCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "google_clusters_delete", envMap)}
}

func CreateServiceAccountCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "google_clusters_create_service_account", envMap)}
}
