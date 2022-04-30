package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/services"
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
		domain := must("DOMAIN")
		for _, s := range services.InstallContourCmd(resourcesDirectory, domain) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var certManagerCmd = &cobra.Command{
	Use:   "cert-manager",
	Short: "Install cert-manager",
	Run: func(cmd *cobra.Command, args []string) {
		emailAddress := must("EMAIL_ADDRESS")
		for _, s := range services.InstallCertManagerCmd(resourcesDirectory, emailAddress) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var harborCmd = &cobra.Command{
	Use:   "harbor",
	Short: "Install harbor",
	Run: func(cmd *cobra.Command, args []string) {
		domain := must("DOMAIN")
		emailAddress := must("EMAIL_ADDRESS")
		password := must("PASSWORD")
		for _, s := range services.InstallHarborCmd(resourcesDirectory, domain, emailAddress, password) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
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
