package pipelines

import (
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
)

func PushPipelineImageCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "pipelines_push_build_image", envMap)}
}

func DeployPipelineCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "pipelines_deploy_pipeline", envMap)}
}

func DeletePipelineCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "pipelines_delete_pipeline", envMap)}
}
