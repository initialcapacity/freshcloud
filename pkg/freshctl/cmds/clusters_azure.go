package cmds

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/azuresupport"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	clustersCmd.AddCommand(azureClusterCmd)
	azureClusterCmd.AddCommand(azureClustersCreateResourceGroupCmd)
	azureClusterCmd.AddCommand(azureClustersDeleteResourceGroupCmd)
	azureClusterCmd.AddCommand(azureClustersCreateCmd)
	azureClusterCmd.AddCommand(azureConfigureCmd)
	azureClusterCmd.AddCommand(azureConfigureRegistryCmd)
}

var azureClusterCmd = &cobra.Command{
	Use:   "az",
	Short: "Manage azure clusters",
}

var azureClustersCreateResourceGroupCmd = &cobra.Command{
	Use:   "create-resource-group",
	Short: "Create an azure resource group",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"AZURE_LOCATION",
			"AZURE_RESOURCE_GROUP",
		)
		writeCommands(cmd.OutOrStderr(), azuresupport.CreateResourceGroupCmd(resourcesLocation, env))
	},
}

var azureClustersDeleteResourceGroupCmd = &cobra.Command{
	Use:   "delete-resource-group",
	Short: "Delete an azure resource group",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"AZURE_RESOURCE_GROUP",
		)
		writeCommands(cmd.OutOrStderr(), azuresupport.DeleteResourceGroupCmd(resourcesLocation, env))
	},
}

var azureClustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an azure cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"AZURE_RESOURCE_GROUP",
		)
		writeCommands(cmd.OutOrStderr(), azuresupport.CreateClustersCmd(resourcesLocation, env))
	},
}

var azureConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"K8S_CLUSTER_NAME",
			"AZURE_RESOURCE_GROUP",
		)
		writeCommands(cmd.OutOrStderr(), azuresupport.ConfigureCmd(resourcesLocation, env))
	},
}

var azureConfigureRegistryCmd = &cobra.Command{
	Use:   "configure-registry",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"REGISTRY_CLUSTER_NAME",
			"AZURE_RESOURCE_GROUP",
		)
		writeCommands(cmd.OutOrStderr(), azuresupport.ConfigureRegistryCmd(resourcesLocation, env))
	},
}
