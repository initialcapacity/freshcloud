package azuresupport

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func ConfigureCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "azure_clusters_configure", envMap)}
}

func CreateClustersCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "azure_clusters_create", envMap)}
}

func CreateResourceGroupCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "azure_clusters_create_resource_group", envMap)}
}

func DeleteResourceGroupCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "azure_clusters_delete_resource_group", envMap)}
}
