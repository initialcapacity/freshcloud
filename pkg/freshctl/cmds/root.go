package cmds

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
)

var resourcesLocation string
var outOrStderr io.Writer
var outOrStderrOverride bytes.Buffer
var execute bool
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
		freshCmd := fmt.Sprintf("Running freshcloud[%v]", cmd.Name())
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), freshCmd+"\n")
		_, _ = color.New(color.FgGreen).Fprintf(cmd.OutOrStderr(), fmt.Sprintf("%v\n", args[1]))
		if cmd.Name() != "version" {
			if execute {
				command := exec.Command("/bin/bash", args...)
				command.Stdout = os.Stdout
				command.Stderr = os.Stderr
				command.Stdin = os.Stdin
				err := command.Run()
				if err != nil {
					_, _ = color.New(color.FgRed).Printf("Unable to run %v command %v\n", freshCmd, err)
				} else {
					_, _ = color.New(color.FgGreen).Printf("Success! running %v command\n", freshCmd)
				}
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&execute, "execute", "e", false, "execute the command")
}

func Fresh() *cobra.Command {
	if found := os.Getenv("FRESH_RESOURCES"); found != "" {
		resourcesLocation = found
		log.Printf("Found resources location %s", resourcesLocation)
	} else {
		log.Println("Unable to find resources location, using embedded resources.")
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	return rootCmd
}

func MakeEnvironmentMap(envs []string) map[string]string {
	env := make(map[string]string)
	for _, v := range envs {
		for i := 0; i < len(v); i++ {
			if v[i] == '=' {
				env[v[0:i]] = v[i+1:]
			}
		}
	}
	return env
}

func requiredString(env map[string]string, required ...string) map[string]string {
	for _, r := range required {
		if env[r] == "" {
			panic(fmt.Sprintf("missing required argument %v %v", r, env[r]))
		}
	}
	return env
}

func writeCommands(w io.Writer, cmds []string) {
	for _, c := range cmds {
		_, _ = fmt.Fprintf(w, c+"\n")
	}
}
