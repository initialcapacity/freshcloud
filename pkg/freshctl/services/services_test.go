package services_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/services"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestName(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := services.InstallContourCmd(resourcesDirectory, "aDomain")
	expected := `kubectl create namespace projectcontour
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install ingress bitnami/contour -n projectcontour
if [ $? != 0 ]; then
    echo "Failed to install Contour. Bummer"
    exit 1
fi
echo "Waiting to get the load-balancer IP."
while true; do
if [ -z "$LB" ]; then
    LB=$(kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}')
    sleep 3;
    else
    echo "Create a DNS A for *.aDomain to $LB"
    break
fi
done
`
	assert.Equal(t, expected, clusterCmd[0])
}
