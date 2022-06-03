package awssupport

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func ConfigureCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "aws_clusters_configure", envMap)}
}

func ConfigureRegistryCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "aws_clusters_configure_registry", envMap)}
}

func CreateClustersCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "aws_clusters_create", envMap)}
}

func DeleteClustersCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "aws_clusters_delete", envMap)}
}
