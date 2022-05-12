package applications

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func PushImageCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "applications_push_image", envMap)}
}

func DeployAppCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "applications_deploy_app", envMap)}
}
