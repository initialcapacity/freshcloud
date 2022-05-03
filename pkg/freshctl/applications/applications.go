package applications

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func PushImageCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "push_image", envMap)}
}

func DeployAppCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "deploy_app", envMap)}
}
