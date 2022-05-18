package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/support"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(resourcesCmd)
	resourcesCmd.AddCommand(copyResourcesCmd)
}

var resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "Manage resources",
}

var copyResourcesCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy resources to a local directory",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()))
		for _, s := range support.CopyResourcesCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
