package services_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/services"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestInstallContour(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := services.InstallContourCmd(resourcesDirectory, "aDomain")
	expected := `kubectl create namespace projectcontour
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install ingress bitnami/contour -n projectcontour --version 7.3.8
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
        echo "Remove contour by running - kubectl delete ns projectcontour"
    break
fi
done`
	assert.Equal(t, expected, clusterCmd[0])
}

func TestInstallCertManager(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := services.InstallCertManagerCmd(resourcesDirectory, "anEmail")
	expected := `function wait_for_cert_manager() {
    echo "Waiting for pods in cert-manager to become ready."
    while true; do
        STATUS=$(kubectl get pods -n cert-manager | egrep -v 'Running|NAME|Completed')
        if [ -z "$STATUS" ]; then
            break
        fi
    done
    echo "All pods are running."
}
kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager --namespace cert-manager \
    --version v1.0.2 --set installCRDs=true
if [ $? != 0 ]; then
    echo "Failed to install Cert-Manager. Bummer"
    exit 1
fi
wait_for_cert_manager
sleep 10
cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
    name: letsencrypt-staging
spec:
    acme:
        email: anEmail
        privateKeySecretRef:
            name: letsencrypt-staging
        server: https://acme-staging-v02.api.letsencrypt.org/directory
        solvers:
            - http01:
                ingress:
                    class: contour
EOF
cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
    name: letsencrypt-prod
spec:
    acme:
        email: anEmail
        privateKeySecretRef:
            name: letsencrypt-prod
        server: https://acme-v02.api.letsencrypt.org/directory
        solvers:
            - http01:
                ingress:
                    class: contour
EOF`
	assert.Equal(t, expected, clusterCmd[0])
}

func TestInstallHarbor(t *testing.T) {
	cmd := services.InstallHarborCmd()
	assert.Equal(t, "echo todo", cmd[0])
}

func TestInstallConcourse(t *testing.T) {
	cmd := services.InstallConcourseCmd()
	assert.Equal(t, "echo todo", cmd[0])
}

func TestInstallKpack(t *testing.T) {
	cmd := services.InstallKpackCmd()
	assert.Equal(t, "echo todo", cmd[0])
}
