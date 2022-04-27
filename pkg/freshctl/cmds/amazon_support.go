package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	clustersCmd.AddCommand(awsClustersCmd)
	awsClustersCmd.AddCommand(aservicesCmd)
	awsClustersCmd.AddCommand(aclustersCreateCmd)
	awsClustersCmd.AddCommand(aclustersDeleteCmd)
	awsClustersCmd.AddCommand(aconfigureCmd)
}

var awsClustersCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage aws clusters",
}

var aservicesCmd = &cobra.Command{
	Use:   "enable-services",
	Short: "Enable amazon web API services",
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
