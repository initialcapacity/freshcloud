package cmds

import (
	"github.com/spf13/cobra"
	"path/filepath"
	"runtime"
)

var rootCmd = &cobra.Command{Use: "freshctl"}
var resourcesDirectory string

func Fresh() *cobra.Command {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory = filepath.Join(file, "../../resources")
	return rootCmd
}
