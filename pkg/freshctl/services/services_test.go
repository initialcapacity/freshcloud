package services_test

import (
	"github.com/initialcapacity/freshcloud/pkg/freshctl/services"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestInstallContour(t *testing.T) {
	env := map[string]string{
		"DOMAIN": "aDomain",
	}
	contourCmd := services.InstallContourCmd(resourcesLocation(), env)
	expected := `kubectl create namespace projectcontour
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install ingress bitnami/contour -n projectcontour --version 7.3.8
if [ $? != 0 ]; then
  echo "Failed to install Contour. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n projectcontour
sleep 10 # waiting for an ip address
load_balancer=$(kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}')
echo "Create a DNS A for *.aDomain to $load_balancer"
echo "Remove contour by running - kubectl delete ns projectcontour"`
	assert.Equal(t, expected, contourCmd[0])
}

func TestInstallCertManager(t *testing.T) {
	env := map[string]string{
		"EMAIL_ADDRESS": "anEmail",
	}
	certManagerCmd := services.InstallCertManagerCmd(resourcesLocation(), env)
	expected := `kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.7.1 --set installCRDs=true
if [ $? != 0 ]; then
  echo "Failed to install Cert-Manager. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n cert-manager
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
	assert.Equal(t, expected, certManagerCmd[0])
}

func TestInstallHarbor(t *testing.T) {
	env := map[string]string{
		"DOMAIN":        "aDomain",
		"EMAIL_ADDRESS": "anEmail",
		"PASSWORD":      "aPassword",
	}
	cmd := services.InstallHarborCmd(resourcesLocation(), env)
	expected := `mkdir -p .freshcloud
cat <<EOF > .freshcloud/harbor-values.yaml
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
helm install harbor bitnami/harbor -f .freshcloud/harbor-values.yaml -n harbor --version 11.2.4
if [ $? != 0 ]; then
  echo "Failed to install Harbor. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n harbor
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
	env := map[string]string{
		"DOMAIN":   "aDomain",
		"PASSWORD": "aPassword",
	}
	cmd := services.InstallConcourseCmd(resourcesLocation(), env)
	expected := `mkdir -p .freshcloud
cat <<EOF > .freshcloud/concourse-values.yaml
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
      projectcontour.io/response-timeout: 10m
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
helm install concourse concourse/concourse -f .freshcloud/concourse-values.yaml -n concourse
if [ $? != 0 ]; then
  echo "Failed to install Concourse. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n concourse
echo "Remove concourse by running - kubectl delete ns concourse"`
	assert.Equal(t, expected, cmd[0])
}

func TestInstallKpack(t *testing.T) {
	env := map[string]string{
		"DOMAIN":   "aDomain",
		"PASSWORD": "aPassword",
	}
	cmd := services.InstallKpackCmd(resourcesLocation(), env)
	expected := `kubectl create namespace kpack
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
kubectl apply -f https://github.com/pivotal/kpack/releases/download/v0.5.1/release-0.5.1.yaml
if [ $? != 0 ]; then
  echo "Failed to install Kpack. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n kpack
REGISTRY="registry.aDomain"
cat <<EOF | kubectl apply -f -
apiVersion: kpack.io/v1alpha1
kind: ClusterStack
metadata:
  name: base
spec:
  id: "heroku-20"
  buildImage:
    image: "heroku/pack:20-build"
  runImage:
    image: "heroku/pack:20"
EOF
cat <<EOF | kubectl apply -f -
apiVersion: kpack.io/v1alpha1
kind: ClusterStore
metadata:
  name: default
spec:
  sources:
  - image: heroku/buildpacks:20
EOF
kubectl create secret docker-registry ${REGISTRY} \
  --docker-username=admin \
  --docker-password=aPassword \
  --docker-server=https://${REGISTRY}/ \
  --namespace default
echo "Remove kpack by running - kubectl delete ns kpack"`
	assert.Equal(t, expected, cmd[0])
}

func resourcesLocation() string {
	_, file, _, _ := runtime.Caller(0)
	resourcesLocation := filepath.Join(file, "../../resources")
	return resourcesLocation
}
