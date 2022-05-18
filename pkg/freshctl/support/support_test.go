package support_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/support"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCopyResourcesCmd(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesLocation := filepath.Join(file, "../../resources")
	env := map[string]string{}
	copyCmd := support.CopyResourcesCmd(resourcesLocation, env)
	expected := `resources=( 'applications_deploy_app.sh'
'applications_push_image.sh'
'google_clusters_configure.sh'
'google_clusters_create.sh'
'google_clusters_create_service_account.sh'
'google_clusters_delete.sh'
'google_clusters_enable_services.sh'
'google_clusters_list.sh'
'pipelines_delete_pipeline.sh'
'pipelines_deploy_pipeline.sh'
'pipelines_push_build_image.sh'
'services_install_cert_manager.sh'
'services_install_concourse.sh'
'services_install_contour.sh'
'services_install_harbor.sh'
'services_install_kpack.sh' )
mkdir -p local_resources
for resource in ${resources[@]}; do
  curl -L https://raw.githubusercontent.com/initialcapacity/freshcloud/main/pkg/freshctl/resources/${resource} -o local_resources/${resource}
done`
	assert.Equal(t, expected, copyCmd[0])
}
