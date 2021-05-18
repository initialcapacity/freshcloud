#!/usr/bin/env bash
#
# Install Concourse

source ../.env_development.sh
source ../components/kubernetes-support/kubectl-support.sh

function create_concourse_values() {

  cat <<EOF > concourse-values.yaml
concourse:
  web:
    externalUrl: https://ci.$DOMAIN
    auth:
      mainTeam:
        localUser: "admin"
secrets:
  localUsers: "admin:$PASSWD"
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
      - ci.$DOMAIN
    tls:
      - hosts:
          - ci.$DOMAIN
        secretName: concourse-cert
    hosts:
      - ci.$DOMAIN
EOF
}

function helm_install_concourse() {

  helm repo add concourse https://concourse-charts.storage.googleapis.com/
  kubectl create namespace concourse
  helm install concourse concourse/concourse -f concourse-values.yaml -n concourse
  rm -f concourse-values.yaml
}

create_concourse_values
helm_install_concourse
wait_for_ready concourse


cat << EOF
url: https://ci.$DOMAIN
username: admin
password: $PASSWD
EOF
