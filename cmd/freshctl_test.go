package main

import (
	"bytes"
	"fmt"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/cmds"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestExec(t *testing.T) {
	main()
}

func setup() {
	_ = os.Setenv("GCP_PROJECT_ID", "aProject")
	_ = os.Setenv("GCP_ZONE", "aZone")
	_ = os.Setenv("GCP_CLUSTER_NAME", "aClusterName")
}

func TestCommands(t *testing.T) {
	setup()

	var buf bytes.Buffer

	fresh := cmds.Fresh()
	fresh.SetOut(&buf)

	fresh.SetArgs([]string{"version"})
	_ = fresh.Execute()

	version, _ := io.ReadAll(&buf)
	assert.Equal(t, "freshcloud[version]\nfreshctl version 0.1\n\n", string(version))

	fresh.SetArgs([]string{"services"})
	assert.NoError(t, fresh.Execute())

	fresh.SetArgs([]string{"clusters"})
	assert.NoError(t, fresh.Execute())

	fresh.SetArgs([]string{"configure"})
	assert.NoError(t, fresh.Execute())
}

func TestCommands_withFlags(t *testing.T) {
	setup()

	var buf bytes.Buffer

	fresh := cmds.Fresh()
	fresh.SetOut(&buf)

	_ = fresh.Flags().Set("execute", "true")
	fresh.SetArgs([]string{"version"})
	assert.NoError(t, fresh.Execute())
	d, _ := io.ReadAll(&buf)
	fmt.Println(string(d))
	assert.NotContains(t, string(d), "executing command.")
}
