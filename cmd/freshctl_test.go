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

	clusterCommands := map[string][]string{
		"gservices":  {"services", "gcp", "enable"},
		"gcreate":    {"clusters", "gcp", "create"},
		"gconfigure": {"clusters", "gcp", "configure"},
		"gdelete":    {"clusters", "gcp", "delete"},

		"aservices":  {"services", "aws", "enable"},
		"acreate":    {"clusters", "aws", "create"},
		"aconfigure": {"clusters", "aws", "configure"},
		"adelete":    {"clusters", "aws", "delete"},
	}
	for _, v := range clusterCommands {
		fresh.SetArgs(v)
		assert.NoError(t, fresh.Execute())
	}
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
