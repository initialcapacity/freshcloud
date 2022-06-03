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
	_ = os.Setenv("GCP_SERVICE_ACCOUNT_JSON", "aJsonFile")

	_ = os.Setenv("AZURE_RESOURCE_GROUP", "aGroup")
	_ = os.Setenv("AZURE_LOCATION", "aLocation")

	_ = os.Setenv("AWS_REGION", "aLocation")

	_ = os.Setenv("K8S_CLUSTER_NAME", "aCluster")

	_ = os.Setenv("DOMAIN", "aDomain")
	_ = os.Setenv("EMAIL_ADDRESS", "anEmail")
	_ = os.Setenv("PASSWORD", "aPassword")

	_ = os.Setenv("REGISTRY_DOMAIN", "aRegistryDomain")
	_ = os.Setenv("REGISTRY_PASSWORD", "aRegistryPassword")
	_ = os.Setenv("REGISTRY_CLUSTER_NAME", "aClusterName")

	_ = os.Setenv("APP_NAME", "anAppName")
	_ = os.Setenv("APP_IMAGE_NAME", "anImageName")
	_ = os.Setenv("APP_CONFIGURATION_PATH", "aPath")
	_ = os.Setenv("APP_PIPELINE_PATH", "aPipelinePath")
	_ = os.Setenv("APP_PIPELINE_CONFIGURATION_PATH", "aPipelineConfigPath")
}

func TestCommands(t *testing.T) {
	setup()

	var buf bytes.Buffer

	fresh := cmds.Fresh()
	fresh.SetOut(&buf)

	fresh.SetArgs([]string{"version"})
	_ = fresh.Execute()

	version, _ := io.ReadAll(&buf)
	assert.Equal(t, "Running freshcloud[version]\nfreshctl version 0.1\n", string(version))

	clusterCommands := map[string][]string{
		"gservices":  {"clusters", "gcp", "enable-services"},
		"gcreate":    {"clusters", "gcp", "create"},
		"lcreate":    {"clusters", "gcp", "list"},
		"gconfigure": {"clusters", "gcp", "configure"},
		"gdelete":    {"clusters", "gcp", "delete"},

		"azcreate":         {"clusters", "az", "create"},
		"azcreateresource": {"clusters", "az", "create-resource-group"},
		"azdeleteresource": {"clusters", "az", "delete-resource-group"},

		"acreate":    {"clusters", "aws", "create"},
		"aconfigure": {"clusters", "aws", "configure"},
		"adelete":    {"clusters", "aws", "delete"},

		"contour":     {"services", "add", "contour"},
		"certmanager": {"services", "add", "cert-manager"},
		"harbor":      {"services", "add", "harbor"},
		"concourse":   {"services", "add", "concourse"},
		"kpack":       {"services", "add", "kpack"},

		"push":   {"applications", "push"},
		"deploy": {"applications", "deploy"},

		"ppush":   {"pipelines", "push-build-image"},
		"pdeploy": {"pipelines", "deploy"},
		"pdelete": {"pipelines", "delete"},
	}
	for _, value := range clusterCommands {
		fresh.SetArgs(value)
		_ = fresh.Execute()
		d, _ := io.ReadAll(&buf)
		assert.Contains(t, string(d),
			fmt.Sprintf("Running freshcloud[%v]", value[len(value)-1]),
			fmt.Sprintf("failed on %v", value),
		)
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
	assert.NotContains(t, string(d), "executing command.")
}
