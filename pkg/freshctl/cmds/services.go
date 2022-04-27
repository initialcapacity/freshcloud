package cmds

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(servicesCmd)
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Enable cloud infrastructure.",
}
