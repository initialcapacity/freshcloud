package main

import (
	"bytes"
	"github.com/initialcapacity/freshcloud/pkg/freshctl/cmds"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestExec(t *testing.T) {
	main()
}

func TestVersion(t *testing.T) {
	_ = os.Setenv("GCP_PROJECT_ID", "aProject")
	_ = os.Setenv("GCP_ZONE", "aZone")
	_ = os.Setenv("GCP_CLUSTER_NAME", "aClusterName")

	var buf bytes.Buffer

	fresh := cmds.Fresh()
	fresh.SetOut(&buf)

	fresh.SetArgs([]string{"version"})
	_ = fresh.Execute()

	version, _ := io.ReadAll(&buf)
	assert.Equal(t, "freshctl version 0.1\n", string(version))

	fresh.SetArgs([]string{"cluster"})
	assert.NoError(t, fresh.Execute())

	fresh.SetArgs([]string{"configure"})
	assert.NoError(t, fresh.Execute())
}
