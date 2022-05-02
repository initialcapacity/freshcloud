package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/applications"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(applicationsCmd)
	applicationsCmd.AddCommand(pushCmd)
	applicationsCmd.AddCommand(deployCmd)
}

var applicationsCmd = &cobra.Command{
	Use:   "applications",
	Short: "Manage applications",
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an image to the registry",
	Run: func(cmd *cobra.Command, args []string) {

		domain := must("REGISTRY_DOMAIN")
		password := must("REGISTRY_PASSWORD")
		app := must("APP_NAME")
		image := must("APP_IMAGE_NAME")

		for _, s := range applications.PushImageCmd(resourcesDirectory, domain, password, app, image) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an app to a cluster",
	Run: func(cmd *cobra.Command, args []string) {

		domain := must("REGISTRY_DOMAIN")
		password := must("REGISTRY_PASSWORD")
		appDomain := must("DOMAIN")
		app := must("APP_NAME")
		image := must("APP_IMAGE_NAME")
		path := must("APP_CONFIGURATION_PATH")

		for _, s := range applications.DeployAppCmd(resourcesDirectory, domain, password, app, appDomain, image, path) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
