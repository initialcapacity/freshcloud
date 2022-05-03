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
	Short: "Push an image to the registry",
	Run: func(cmd *cobra.Command, args []string) {
		env := map[string]string{
			"REGISTRY_DOMAIN":   requiredEnv("REGISTRY_DOMAIN"),
			"REGISTRY_PASSWORD": requiredEnv("REGISTRY_PASSWORD"),
			"APP_NAME":          requiredEnv("APP_NAME"),
			"APP_IMAGE_NAME":    requiredEnv("APP_IMAGE_NAME"),
		}
		for _, s := range applications.PushImageCmd(resourcesDirectory, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an app to a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"REGISTRY_DOMAIN", "REGISTRY_PASSWORD", "APP_NAME", "DOMAIN", "APP_IMAGE_NAME", "APP_CONFIGURATION_PATH")
		for _, s := range applications.DeployAppCmd(resourcesDirectory, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
