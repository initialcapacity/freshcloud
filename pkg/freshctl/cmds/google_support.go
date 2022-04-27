package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/googlecloudsupport"
	"github.com/spf13/cobra"
)

func init() {
	servicesCmd.AddCommand(googleServicesCmd)
	googleServicesCmd.AddCommand(gservicesCmd)

	clustersCmd.AddCommand(googleClusterCmd)
	googleClusterCmd.AddCommand(gclustersCreateCmd)
	googleClusterCmd.AddCommand(gclustersDeleteCmd)
	googleClusterCmd.AddCommand(gconfigureCmd)
}

var googleServicesCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Manage gcp services.",
}

var googleClusterCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Manage gcp infrastructure.",
}

var gservicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Enable google cloud services",
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range googlecloudsupport.EnableServicesCmd() {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var gclustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")

		_, _ = fmt.Fprintf(cmd.OutOrStderr(), googlecloudsupport.CreateClustersCmd(resourcesDirectory, projectID, zone, clusterName))
	},
}

var gclustersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a google cloud cluster",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")

		_, _ = fmt.Fprintf(cmd.OutOrStderr(), googlecloudsupport.DeleteClustersCmd(projectID, zone, clusterName))
	},
}

var gconfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure kubectl for google cloud",
	Run: func(cmd *cobra.Command, args []string) {

		projectID := must("GCP_PROJECT_ID")
		zone := must("GCP_ZONE")
		clusterName := must("GCP_CLUSTER_NAME")

		_, _ = fmt.Fprintf(cmd.OutOrStderr(), googlecloudsupport.ConfigureCmd(projectID, zone, clusterName))
	},
}
