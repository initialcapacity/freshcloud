package cmds

import (
	"github.com/spf13/cobra"
)

func init() {
	clustersCmd.AddCommand(awsClustersCmd)
	awsClustersCmd.AddCommand(awsServicesCmd)
	awsClustersCmd.AddCommand(awsClustersCreateCmd)
	awsClustersCmd.AddCommand(awsClustersListCmd)
	awsClustersCmd.AddCommand(awsClustersDeleteCmd)
	awsClustersCmd.AddCommand(awsConfigureCmd)
}

var awsClustersCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage aws clusters",
}

var awsServicesCmd = &cobra.Command{
	Use:   "enable-services",
	Short: "Enable amazon web API services",
	Run:   func(cmd *cobra.Command, args []string) { writeCommands(cmd.OutOrStderr(), []string{"todo"}) },
}

var awsClustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an aws cluster",
	Run:   func(cmd *cobra.Command, args []string) { writeCommands(cmd.OutOrStderr(), []string{"todo"}) },
}

var awsClustersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List aws clusters",
	Run:   func(cmd *cobra.Command, args []string) { writeCommands(cmd.OutOrStderr(), []string{"todo"}) },
}

var awsClustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an aws cluster",
	Run:   func(cmd *cobra.Command, args []string) { writeCommands(cmd.OutOrStderr(), []string{"todo"}) },
}

var awsConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for aws",
	Run:   func(cmd *cobra.Command, args []string) { writeCommands(cmd.OutOrStderr(), []string{"todo"}) },
}
