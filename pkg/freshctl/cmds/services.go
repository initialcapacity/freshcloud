package cmds

import (
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/services"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(servicesCmd)
	servicesCmd.AddCommand(servicesAddCmd)
	servicesAddCmd.AddCommand(contourCmd)
	servicesAddCmd.AddCommand(certManagerCmd)
	servicesAddCmd.AddCommand(harborCmd)
	servicesAddCmd.AddCommand(concourseCmd)
	servicesAddCmd.AddCommand(kpackCmd)
	servicesCmd.AddCommand(servicesRemoveCmd)
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Manage cluster services",
}

var servicesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add services",
}

var servicesRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove services",
}

var contourCmd = &cobra.Command{
	Use:   "contour",
	Short: "Install contour",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"DOMAIN",
			"PASSWORD",
		)
		for _, s := range services.InstallContourCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var certManagerCmd = &cobra.Command{
	Use:   "cert-manager",
	Short: "Install cert-manager",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"EMAIL_ADDRESS",
		)
		for _, s := range services.InstallCertManagerCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var harborCmd = &cobra.Command{
	Use:   "harbor",
	Short: "Install harbor",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"DOMAIN",
			"EMAIL_ADDRESS",
			"PASSWORD",
		)
		for _, s := range services.InstallHarborCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var concourseCmd = &cobra.Command{
	Use:   "concourse",
	Short: "Install concourse",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"DOMAIN",
			"PASSWORD",
		)
		for _, s := range services.InstallConcourseCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}

var kpackCmd = &cobra.Command{
	Use:   "kpack",
	Short: "Install kpack",
	Run: func(cmd *cobra.Command, args []string) {
		env := requiredString(MakeEnvironmentMap(os.Environ()),
			"DOMAIN",
			"PASSWORD",
		)
		for _, s := range services.InstallKpackCmd(resourcesLocation, env) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%s", s))
		}
	},
}
