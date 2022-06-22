package cmds

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clustersCmd)
}

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Deploy and manage kubernetes clusters",
}
