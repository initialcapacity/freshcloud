package pipelines

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func PushPipelineImageCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "pipelines_push_build_image", envMap)}
}

func DeployPipelineCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "pipelines_deploy_pipeline", envMap)}
}

func DeletePipelineCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "pipelines_delete_pipeline", envMap)}
}
