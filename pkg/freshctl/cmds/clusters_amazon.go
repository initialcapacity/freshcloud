package cmds

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/awssupport"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	clustersCmd.AddCommand(awsClustersCmd)
	awsClustersCmd.AddCommand(awsClustersCreateCmd)
	awsClustersCmd.AddCommand(awsClustersDeleteCmd)
	awsClustersCmd.AddCommand(awsConfigureCmd)
	awsClustersCmd.AddCommand(awsConfigureRegistryCmd)
}

var awsClustersCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage aws clusters",
}

var awsClustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an aws cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"AWS_REGION",
		)
		writeCommands(cmd.OutOrStderr(), awssupport.CreateClustersCmd(resourcesLocation, env))
	},
}

var awsClustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an aws cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"AWS_REGION",
		)
		writeCommands(cmd.OutOrStderr(), awssupport.DeleteClustersCmd(resourcesLocation, env))
	},
}

var awsConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for aws",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"AWS_REGION",
		)
		writeCommands(cmd.OutOrStderr(), awssupport.ConfigureCmd(resourcesLocation, env))
	},
}

var awsConfigureRegistryCmd = &cobra.Command{
	Use:   "configure-registry",
	Short: "Configure kubectl for aws",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"REGISTRY_CLUSTER_NAME",
			"AWS_REGION",
		)
		writeCommands(cmd.OutOrStderr(), awssupport.ConfigureRegistryCmd(resourcesLocation, env))
	},
}
