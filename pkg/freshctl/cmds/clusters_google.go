package cmds

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/googlecloudsupport"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	clustersCmd.AddCommand(googleClusterCmd)
	googleClusterCmd.AddCommand(googleServicesCmd)
	googleClusterCmd.AddCommand(googleClustersCreateCmd)
	googleClusterCmd.AddCommand(googleClustersListCmd)
	googleClusterCmd.AddCommand(googleClustersDeleteCmd)
	googleClusterCmd.AddCommand(googleConfigureCmd)
	googleClusterCmd.AddCommand(googleConfigureRegistryCmd)
	googleClusterCmd.AddCommand(googleCreateServiceAccountCmd)
}

var googleClusterCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Manage google cloud clusters",
}

var googleServicesCmd = &cobra.Command{
	Use:   "enable-services",
	Short: "Enable google cloud API services",
	Run: func(cmd *cobra.Command, args []string) {
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.EnableServicesCmd(resourcesLocation))
	},
}

var googleClustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"GCP_PROJECT_ID",
			"GCP_ZONE",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.CreateClustersCmd(resourcesLocation, env))
	},
}

var googleClustersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List google cloud clusters",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"GCP_PROJECT_ID",
			"GCP_ZONE",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ListClustersCmd(resourcesLocation, env))
	},
}

var googleClustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"GCP_PROJECT_ID",
			"GCP_ZONE",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.DeleteClustersCmd(resourcesLocation, env))
	},
}

var googleConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"GCP_PROJECT_ID",
			"GCP_ZONE",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ConfigureCmd(resourcesLocation, env))
	},
}

var googleConfigureRegistryCmd = &cobra.Command{
	Use:   "configure-registry",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"REGISTRY_CLUSTER_NAME",
			"GCP_PROJECT_ID",
			"GCP_ZONE",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ConfigureRegistryCmd(resourcesLocation, env))
	},
}

var googleCreateServiceAccountCmd = &cobra.Command{
	Use:   "create-service-account",
	Short: "Create a service account for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"GCP_PROJECT_ID",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.CreateServiceAccountCmd(resourcesLocation, env))
	},
}
