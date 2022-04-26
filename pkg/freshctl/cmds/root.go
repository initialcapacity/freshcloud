package cmds

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"os/exec"
	"path/filepath"
	"runtime"
)

var outOrStderr io.Writer
var outOrStderrOverride bytes.Buffer

var rootCmd = &cobra.Command{
	Use: "freshctl",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		outOrStderr = cmd.OutOrStderr()
		outOrStderrOverride = bytes.Buffer{}
		cmd.SetOut(&outOrStderrOverride)
		cmd.SetErr(&outOrStderrOverride)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		arg, _ := io.ReadAll(&outOrStderrOverride)
		args = append(args, "-c")
		args = append(args, string(arg))
		cmd.SetOut(outOrStderr)
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), "freshcloud[%v]\n", cmd.Name())
		_, _ = color.New(color.FgGreen).Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%v\n\n", args[1]))
		if cmd.Name() != "version" {
			if execute {
				var commandOut bytes.Buffer
				command := exec.Command("/bin/bash", args...)
				command.Stdout = &commandOut
				command.Stderr = &commandOut

				err := command.Run()
				if err != nil {
					_, _ = fmt.Printf("Unable to run cmd %v\n.", err)
				}
				_, _ = fmt.Fprintf(cmd.OutOrStderr(), commandOut.String())
			}
		}
	},
}

var resourcesDirectory string

func Fresh() *cobra.Command {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory = filepath.Join(file, "../../resources")
	return rootCmd
}
