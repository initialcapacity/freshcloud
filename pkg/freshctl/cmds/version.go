package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/templatesupport"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version number",
	Run: func(cmd *cobra.Command, args []string) {
		_ = os.Setenv("Version", "0.1")
		for _, s := range []string{templatesupport.Parse(resourcesLocation, "version", MakeEnvironmentMap(os.Environ()))} {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
