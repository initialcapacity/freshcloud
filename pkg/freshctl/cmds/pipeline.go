package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/pipelines"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pipelinesCmd)
	pipelinesCmd.AddCommand(pipelinePushImageCmd)
	pipelinesCmd.AddCommand(pipelineDeployCmd)
	pipelinesCmd.AddCommand(pipelineDeleteCmd)
}

var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "Manage pipelines",
}

var pipelinePushImageCmd = &cobra.Command{
	Use:   "push-build-image",
	Short: "Push a build image to the registry",
	Run: func(cmd *cobra.Command, args []string) {
		env := map[string]string{
			"REGISTRY_DOMAIN":   requiredEnv("REGISTRY_DOMAIN"),
			"REGISTRY_PASSWORD": requiredEnv("REGISTRY_PASSWORD"),
		}
		for _, s := range pipelines.PushPipelineImageCmd(resourcesDirectory, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var pipelineDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		env := map[string]string{
			"GCP_PROJECT_ID":                  requiredEnv("GCP_PROJECT_ID"),
			"GCP_ZONE":                        requiredEnv("GCP_ZONE"),
			"GCP_CLUSTER_NAME":                requiredEnv("GCP_CLUSTER_NAME"),
			"GCP_SERVICE_ACCOUNT_JSON":        requiredEnv("GCP_SERVICE_ACCOUNT_JSON"),
			"REGISTRY_DOMAIN":                 requiredEnv("REGISTRY_DOMAIN"),
			"REGISTRY_PASSWORD":               requiredEnv("REGISTRY_PASSWORD"),
			"REGISTRY_CLUSTER_NAME":           requiredEnv("REGISTRY_CLUSTER_NAME"),
			"DOMAIN":                          requiredEnv("DOMAIN"),
			"APP_NAME":                        requiredEnv("APP_NAME"),
			"APP_PIPELINE_CONFIGURATION_PATH": requiredEnv("APP_PIPELINE_CONFIGURATION_PATH"),
			"APP_PIPELINE_PATH":               requiredEnv("APP_PIPELINE_PATH"),
		}
		for _, s := range pipelines.DeployPipelineCmd(resourcesDirectory, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var pipelineDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		env := map[string]string{
			"REGISTRY_DOMAIN":       requiredEnv("REGISTRY_DOMAIN"),
			"REGISTRY_PASSWORD":     requiredEnv("REGISTRY_PASSWORD"),
			"REGISTRY_CLUSTER_NAME": requiredEnv("REGISTRY_CLUSTER_NAME"),
			"APP_NAME":              requiredEnv("APP_NAME"),
		}
		for _, s := range pipelines.DeletePipelineCmd(resourcesDirectory, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
