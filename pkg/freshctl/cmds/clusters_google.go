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
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.EnableServicesCmd(resourcesDirectory))
	},
}

var googleClustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"GCP_PROJECT_ID",
			"GCP_ZONE",
			"GCP_CLUSTER_NAME",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.CreateClustersCmd(resourcesDirectory, env))
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
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ListClustersCmd(resourcesDirectory, env))
	},
}

var googleClustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"GCP_PROJECT_ID",
			"GCP_ZONE",
			"GCP_CLUSTER_NAME",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.DeleteClustersCmd(resourcesDirectory, env))
	},
}

var googleConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"GCP_PROJECT_ID",
			"GCP_ZONE",
			"GCP_CLUSTER_NAME",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.ConfigureCmd(resourcesDirectory, env))
	},
}

var googleCreateServiceAccountCmd = &cobra.Command{
	Use:   "create-service-account",
	Short: "Create a service account for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"GCP_PROJECT_ID",
		)
		writeCommands(cmd.OutOrStderr(), googlecloudsupport.CreateServiceAccountCmd(resourcesDirectory, env))
	},
}
