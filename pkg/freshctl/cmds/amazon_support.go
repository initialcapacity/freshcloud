package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	servicesCmd.AddCommand(awsServicesCmd)
	awsServicesCmd.AddCommand(aservicesCmd)

	clustersCmd.AddCommand(awsClustersCmd)
	awsClustersCmd.AddCommand(aservicesCmd)
	awsClustersCmd.AddCommand(aclustersCreateCmd)
	awsClustersCmd.AddCommand(aclustersDeleteCmd)
	awsClustersCmd.AddCommand(aconfigureCmd)
}

var awsServicesCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage aws infrastructure.",
}

var awsClustersCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage aws infrastructure.",
}

var aservicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Enable amazon web services",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var aclustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an aws cluster",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var aclustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an aws cluster",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var aconfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for aws",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}
