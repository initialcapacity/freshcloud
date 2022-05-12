package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/applications"
	"github.com/spf13/cobra"
	"os"
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
	Short: "Push an application image to the registry",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"REGISTRY_DOMAIN",
			"REGISTRY_PASSWORD",
			"APP_NAME",
			"APP_IMAGE_NAME",
		)
		for _, s := range applications.PushImageCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an application to a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"REGISTRY_DOMAIN",
			"REGISTRY_PASSWORD",
			"DOMAIN",
			"APP_NAME",
			"APP_IMAGE_NAME",
			"APP_CONFIGURATION_PATH",
		)
		for _, s := range applications.DeployAppCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
