#!/usr/bin/env bash
#
# Create a k8s cluster in TKG for AWS
# See https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.3/vmware-tanzu-kubernetes-grid-13/GUID-tanzu-k8s-clusters-aws.html

# shellcheck disable=SC1091
source ../../.env_development.sh

function tanzu_aws_create_k8s_mgmt_cluster() {
  cat > ../../mgmt-cluster.yaml <<EOF
CLUSTER_NAME: tanzu-mgmt
CLUSTER_PLAN: dev
INFRASTRUCTURE_PROVIDER: aws
ENABLE_CEIP_PARTICIPATION: false
ENABLE_AUDIT_LOGGING: false
CONTROL_PLANE_MACHINE_TYPE: t3.large
NODE_MACHINE_TYPE: m5.large
AWS_REGION: us-east-2
AWS_NODE_AZ: "us-east-2a"
AWS_SSH_KEY_NAME: default
EOF
  tanzu management-cluster create --file ../../mgmt-cluster.yaml
}

function tanzu_aws_delete_k8s_mgmt_cluster() {
  tanzu management-cluster delete tanzu-mgmt
}

function tanzu_aws_create_k8s_cluster() {

  cat > ../../freshcloud-cluster.yaml <<EOF
CLUSTER_NAME: freshcloud
CLUSTER_PLAN: dev
NAMESPACE: default
CNI: antrea
CONTROL_PLANE_MACHINE_TYPE: t3.large
NODE_MACHINE_TYPE: m5.large
CONTROL_PLANE_MACHINE_COUNT: 1
WORKER_MACHINE_COUNT: 3
AWS_REGION: $REGION
AWS_NODE_AZ: $NODE_AZ
AWS_ACCESS_KEY_ID: <encoded:$(echo -ne "$AWS_ACCESS_KEY_ID" | base64)>
AWS_SECRET_ACCESS_KEY: <encoded:$(echo -ne "$AWS_SECRET_ACCESS_KEY" | base64)>
AWS_SSH_KEY_NAME: default
BASTION_HOST_ENABLED: true
EOF
  tanzu cluster create -f ../../freshcloud-cluster.yaml
}

function tanzu_aws_delete_k8s_cluster() {
  tanzu cluster delete freshcloud
}

if [ -z "$AWS_ACCESS_KEY_ID" ] || [ -z "$AWS_SECRET_ACCESS_KEY" ]; then
  printf "Please set the following environment variables in .env_development.sh under root directory:\n"
  printf "AWS_ACCESS_KEY_ID\n"
  printf "AWS_SECRET_ACCESS_KEY\n"
  exit 1
fi

if [ "$1" == 'mgmt' ]; then
  if [ "$2" == 'delete' ]; then
    tanzu_aws_delete_k8s_mgmt_cluster
  else
    tanzu_aws_create_k8s_mgmt_cluster
  fi
  exit 0
fi

if [ "$1" == 'delete' ]; then
  tanzu_aws_delete_k8s_cluster
else
  tanzu_aws_create_k8s_cluster
fi
