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
sleep 5
kubectl wait --for=condition=Ready pods --timeout=900s --all -n projectcontour
sleep 5
LB=$(kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}')
echo "Create a DNS A for *.aDomain to $LB"
echo "Remove contour by running - kubectl delete ns projectcontour"`
	assert.Equal(t, expected, clusterCmd[0])
}

func TestInstallCertManager(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	clusterCmd := services.InstallCertManagerCmd(resourcesDirectory, "anEmail")
	expected := `kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.7.1 --set installCRDs=true
if [ $? != 0 ]; then
    echo "Failed to install Cert-Manager. Bummer"
    exit 1
fi
sleep 5
kubectl wait --for=condition=Ready pods --timeout=900s --all -n cert-manager
sleep 5
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
EOF
echo "Remove cert-manager by running - kubectl delete ns cert-manager"`
	assert.Equal(t, expected, clusterCmd[0])
}

func TestInstallHarbor(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	cmd := services.InstallHarborCmd(resourcesDirectory, "aDomain", "aEmail", "aPassword")
	expected := `cat <<EOF > harbor-values.yaml
harborAdminPassword: aPassword
service:
    type: ClusterIP
    tls:
        enabled: true
        existingSecret: harbor-tls-prod
        notaryExistingSecret: notary-tls-prod
ingress:
    enabled: true
    hosts:
        core: registry.aDomain
        notary: notary.aDomain
    annotations:
        cert-manager.io/cluster-issuer: letsencrypt-prod     # use letsencrypt-prod as the cluster issuer for TLS certs
        ingress.kubernetes.io/force-ssl-redirect: "true"     # force https, even if http is requested
        kubernetes.io/ingress.class: contour                 # using Contour for ingress
        kubernetes.io/tls-acme: "true"                       # using ACME certificates for TLS
externalURL: https://registry.aDomain
portal:
    tls:
        existingSecret: harbor-tls-prod
EOF
kubectl create namespace harbor
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install harbor bitnami/harbor -f harbor-values.yaml -n harbor --version 11.2.4
if [ $? != 0 ]; then
    echo "Failed to install Harbor. Bummer"
    exit 1
fi
sleep 5
kubectl wait --for=condition=Ready pods --timeout=900s --all -n harbor
sleep 5
rm -f harbor-values.yaml
for REPO in {concourse-images,kpack}; do
    echo "Creating: ${REPO} in Harbor."
    curl --user "admin:aPassword" -X POST \
        https://registry.aDomain/api/v2.0/projects \
        -H "Content-type: application/json" --data \
        '{ "project_name": "'${REPO}'",
        "metadata": {
        "auto_scan": "true",
        "enable_content_trust": "false",
        "prevent_vul": "false",
        "public": "true",
        "reuse_sys_cve_whitelist": "true",
        "severity": "high" }
        }'
done
cat << EOF
echo "Remove harbor by running - kubectl delete ns harbor"
url: https://registry.aDomain
username: admin
password: aPassword
EOF`
	assert.Equal(t, expected, cmd[0])
}

func TestInstallConcourse(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	resourcesDirectory := filepath.Join(file, "../../resources")
	cmd := services.InstallConcourseCmd(resourcesDirectory, "aDomain", "aPassword")
	expected := `cat <<EOF > concourse-values.yaml
concourse:
  worker:
    replicaCount: 4
  web:
    externalUrl: https://ci.aDomain
    auth:
      mainTeam:
        localUser: "admin"
secrets:
  localUsers: "admin:aPassword"
web:
  env:
  ingress:
    enabled: true
    annotations:
      cert-manager.io/cluster-issuer: letsencrypt-prod
      kubernetes.io/ingress.class: contour
      ingress.kubernetes.io/force-ssl-redirect: "true"
      projectcontour.io/websocket-routes: "/"
      kubernetes.io/tls-acme: "true"
    hosts:
      - ci.aDomain
    tls:
      - hosts:
          - ci.aDomain
        secretName: concourse-cert
EOF
kubectl create namespace concourse
helm repo add concourse https://concourse-charts.storage.googleapis.com/
helm install concourse concourse/concourse -f concourse-values.yaml -n concourse
if [ $? != 0 ]; then
    echo "Failed to install Concourse. Bummer"
    exit 1
fi
sleep 5
kubectl wait --for=condition=Ready pods --timeout=900s --all -n concourse
sleep 5
rm -f concourse-values.yaml
echo "Remove concourse by running - kubectl delete ns concourse"`
	assert.Equal(t, expected, cmd[0])
}

func TestInstallKpack(t *testing.T) {
	cmd := services.InstallKpackCmd()
	assert.Equal(t, "echo todo", cmd[0])
}
