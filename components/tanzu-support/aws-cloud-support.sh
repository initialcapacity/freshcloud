#!/usr/bin/env bash
#
# Create a k8s cluster in TKG for AWS
# See https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.3/vmware-tanzu-kubernetes-grid-13/GUID-tanzu-k8s-clusters-aws.html

source ../../.env_development.sh

function tanzu_aws_create_k8s_mgmt_cluster() {
  echo "Not yet implemented"
}

function tanzu_aws_create_k8s_cluster() {

cat << EOF | tanzu cluster create -f -
CLUSTER_NAME: freshcloud
CLUSTER_PLAN: dev
NAMESPACE: default
CNI: antrea
CONTROL_PLANE_MACHINE_TYPE: t3.large
NODE_MACHINE_TYPE: m5.large
CONTROL_PLANE_MACHINE_COUNT: 1
WORKER_MACHINE_COUNT: 3
AWS_REGION: us-east-2
AWS_NODE_AZ: "us-east-2a"
AWS_ACCESS_KEY_ID: $AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY: $AWS_SECRET_ACCESS_KEY
AWS_SSH_KEY_NAME: default
BASTION_HOST_ENABLED: false
ENABLE_AUTOSCALER: false
AUTOSCALER_MAX_NODES_TOTAL: "10"
AUTOSCALER_SCALE_DOWN_DELAY_AFTER_ADD: "10m"
AUTOSCALER_SCALE_DOWN_DELAY_AFTER_DELETE: "10s"
AUTOSCALER_SCALE_DOWN_DELAY_AFTER_FAILURE: "3m"
AUTOSCALER_SCALE_DOWN_UNNEEDED_TIME: "10m"
AUTOSCALER_MAX_NODE_PROVISION_TIME: "15m"
EOF
}

function tanzu_aws_delete_k8s_cluster() {
  tanzu cluster delete freshcloud
}

if [ -z "$AWS_ACCESS_KEY_ID" ] || [ -z "$AWS_SECRET_ACCESS_KEY" ]; then
  printf "Please set the following environment variables:\n"
  printf "\tAWS_ACCESS_KEY_ID\n"
  printf "\tAWS_SECRET_ACCESS_KEY\n"
fi

if [ "$1" == 'delete' ]; then
  tanzu_aws_delete_k8s_cluster
else
  tanzu_aws_create_k8s_cluster
fi
