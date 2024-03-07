#!/usr/bin/env bash
#
# Install Concourse

source env.sh
source utils.sh

function create_concourse_values() {

  cat <<EOF > concourse-values.yaml
concourse:
  worker:
    runtime: containerd
  web:
    externalUrl: https://ci.$SUB_DOMAIN
    auth:
      mainTeam:
        localUser: "admin"
secrets:
  localUsers: "admin:$CONCOURSE_PASSWD"
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
      - ci.$SUB_DOMAIN
    tls:
      - hosts:
          - ci.$SUB_DOMAIN
        secretName: concourse-cert
    hosts:
      - ci.$SUB_DOMAIN
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
url: https://ci.$SUB_DOMAIN
username: admin
password: $CONCOURSE_PASSWD
EOF
