package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(servicesCmd)
	servicesCmd.AddCommand(contourCmd)
	servicesCmd.AddCommand(certManagerCmd)
	servicesCmd.AddCommand(harborCmd)
	servicesCmd.AddCommand(concourseCmd)
	servicesCmd.AddCommand(kpackCmd)
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Manage cluster services",
}

var contourCmd = &cobra.Command{
	Use:   "contour",
	Short: "Install contour",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var certManagerCmd = &cobra.Command{
	Use:   "certmanager",
	Short: "Install cert-manager",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var harborCmd = &cobra.Command{
	Use:   "harbor",
	Short: "Install harbor",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var concourseCmd = &cobra.Command{
	Use:   "concourse",
	Short: "Install concourse",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}

var kpackCmd = &cobra.Command{
	Use:   "kpack",
	Short: "Install kpack",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "todo")
	},
}
