package support

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func CopyResourcesCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "resources_copy_locally", envMap)}
}
