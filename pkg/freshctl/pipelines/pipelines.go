package pipelines

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func PushPipelineImageCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "push_build_image", envMap)}
}

func DeployPipelineCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "deploy_pipeline", envMap)}
}

func DeletePipelineCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "delete_pipeline", envMap)}
}
