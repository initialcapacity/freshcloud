#!/usr/bin/env bash
#
# Install Cert-Manager

source ../.env_development.sh
source ../components/kubernetes-support/kubectl-support.sh


function helm_install_cert-manager() {
  kubectl create namespace cert-manager
  helm repo add jetstack https://charts.jetstack.io
  helm repo update
  helm install cert-manager jetstack/cert-manager --namespace cert-manager \
   --version v1.8.0 --set installCRDs=true

  if [ $? != 0 ]; then
   echo "Failed to install Cert-Manager. Bummer"
   exit 1
  fi
}

function install_ClusterIssuer() {

  cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-staging
spec:
  acme:
    email: $EMAIL_ADDRESS
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
    email: $EMAIL_ADDRESS
    privateKeySecretRef:
      name: letsencrypt-prod
    server: https://acme-v02.api.letsencrypt.org/directory
    solvers:
    - http01:
        ingress:
          class: contour
EOF

}

#helm_install_cert-manager
#wait_for_ready cert-manager
#sleep 10; 
install_ClusterIssuer
