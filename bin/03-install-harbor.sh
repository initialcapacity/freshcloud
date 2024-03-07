#!/usr/bin/env bash
#
# Install Harbor

source env.sh
source utils.sh

function create_harbor_values() {
  kubectl create namespace harbor

  cat <<EOF > harbor-values.yaml
global:
  imagePullSecrets:
    - name: regcred

clusterDomain: cluster.local

adminPassword: "${HARBOR_PASSWD}"
externalURL: https://registry.${SUB_DOMAIN}
proxy:
  noProxy: 127.0.0.1,localhost,.local,.internal
  components:
    - core
    - jobservice
    - trivy
logLevel: debug

ipFamily:
  ipv6:
    enabled: true
  ipv4:
    enabled: true

exposureType: ingress

service:
  type: LoadBalancer
  ports:
    http: 80
    https: 443
    notary: 4443
  externalTrafficPolicy: Cluster

ingress:
  core:
    hostname: registry.${SUB_DOMAIN}
    annotations:
      ingress.kubernetes.io/force-ssl-redirect: 'true'
      ingress.kubernetes.io/proxy-body-size: '0'
      nginx.ingress.kubernetes.io/ssl-redirect: 'true'
      nginx.ingress.kubernetes.io/proxy-body-size: '0'
      cert-manager.io/cluster-issuer: letsencrypt-prod
    tls: true
  notary:
    hostname: notary.${SUB_DOMAIN}
    annotations:
      ingress.kubernetes.io/force-ssl-redirect: 'true'
      ingress.kubernetes.io/proxy-body-size: '0'
      nginx.ingress.kubernetes.io/ssl-redirect: 'true'
      nginx.ingress.kubernetes.io/proxy-body-size: '0'
      cert-manager.io/cluster-issuer: letsencrypt-prod
    tls: true

persistence:
  enabled: true
  persistentVolumeClaim:
    registry:
      accessModes:
        - ReadWriteOnce
      size: 5Gi
    jobservice:
      accessModes:
        - ReadWriteOnce
      size: 1Gi
    chartmuseum:
      accessModes:
        - ReadWriteOnce
      size: 5Gi
    trivy:
      accessModes:
        - ReadWriteOnce
      size: 5Gi
  imageChartStorage:
    type: filesystem
    filesystem:
      rootdirectory: /storage

trivy:
  enabled: true
  vulnType: 'os,library'
  severity: 'UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL'
  resources:
    requests:
      cpu: 200m
      memory: 512Mi
    limits:
      cpu: 1
      memory: 1Gi

postgresql:
  enabled: true
  image:
    tag: 13.9.0-debian-11-r13
  auth:
    postgresPassword: ${HARBOR_PASSWD}
  primary:
    extendedConfiguration: |
      max_connections = 1024
    initdb:
      scripts:
        initial-notaryserver.sql: |
          CREATE DATABASE notaryserver;
          CREATE USER server;
          alter user server with encrypted password 'password';
          GRANT ALL PRIVILEGES ON DATABASE notaryserver TO server;
        initial-notarysigner.sql: |
          CREATE DATABASE notarysigner;
          CREATE USER signer;
          alter user signer with encrypted password 'password';
          GRANT ALL PRIVILEGES ON DATABASE notarysigner TO signer;
        initial-registry.sql: |
          CREATE DATABASE registry ENCODING 'UTF8';
          \c registry;
          CREATE TABLE schema_migrations(version bigint not null primary key, dirty boolean not null);

redis:
  enabled: true
  auth:
    password: ${HARBOR_PASSWD}
EOF
}

function helm_install_harbor() {

  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo update

  helm install harbor bitnami/harbor -f harbor-values.yaml -n harbor --version 19.8.2
  rm -f harbor-values.yaml
}

function create_harbor_repos() {

   for REGISTRY in {concourse-images,kpack}; do

     echo "Creating: ${REGISTRY} in Harbor."

     curl --user "admin:${HARBOR_PASSWD}" -X POST \
         https://registry.${SUB_DOMAIN}/api/v2.0/projects \
         -H "Content-type: application/json" --data \
         '{ "project_name": "'${REGISTRY}'",
          "metadata": {
          "auto_scan": "true",
          "enable_content_trust": "false",
          "prevent_vul": "false",
          "public": "true",
          "reuse_sys_cve_whitelist": "true",
          "severity": "high" }
          }'

    done

}

create_harbor_values
helm_install_harbor
wait_for_ready harbor
echo "Sleeping while Harbor auth is setup."
sleep 120
create_harbor_repos

cat << EOF
url: https://registry.$SUB_DOMAIN
username: admin
password: ${HARBOR_PASSWD}
EOF
