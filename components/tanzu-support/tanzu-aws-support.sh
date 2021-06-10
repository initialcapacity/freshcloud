#!/usr/bin/env bash
#
# Create a k8s cluster in TKG for AWS
# See https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.3/vmware-tanzu-kubernetes-grid-13/GUID-tanzu-k8s-clusters-aws.html

# shellcheck disable=SC1091
source ../../.env_development.sh

function tanzu_aws_create_k8s_mgmt_cluster() {
  local temp_dir=${1:-/tmp}
  local cluster_name=${2:-tanzu-mgmt}
  cat > "$temp_dir/mgmt-cluster.yaml" <<EOF
CLUSTER_NAME: $cluster_name
CLUSTER_PLAN: dev
INFRASTRUCTURE_PROVIDER: aws
ENABLE_CEIP_PARTICIPATION: false
ENABLE_AUDIT_LOGGING: false
CONTROL_PLANE_MACHINE_TYPE: t3.large
NODE_MACHINE_TYPE: m5.large
AWS_REGION: us-east-2
AWS_NODE_AZ: "us-east-2a"
AWS_ACCESS_KEY_ID: <encoded:$(echo -ne "$AWS_ACCESS_KEY_ID" | base64)>
AWS_SECRET_ACCESS_KEY: <encoded:$(echo -ne "$AWS_SECRET_ACCESS_KEY" | base64)>
AWS_SSH_KEY_NAME: default
BASTION_HOST_ENABLED: false
AWS_PRIVATE_NODE_CIDR: 10.0.0.0/24
AWS_PUBLIC_NODE_CIDR: 10.0.1.0/24
AWS_VPC_CIDR: 10.0.0.0/16
SERVICE_CIDR: 100.64.0.0/13
CLUSTER_CIDR: 100.96.0.0/11
EOF
  tanzu management-cluster create --file "$temp_dir/mgmt-cluster.yaml"
  printf "Management cluster file saved to: %s\n" "$temp_dir/mgmt-cluster.yaml"
}

function tanzu_aws_delete_k8s_mgmt_cluster() {
  local cluster_name=${1:-tanzu-mgmt}
  tanzu management-cluster delete "$cluster_name"
}

function tanzu_aws_create_k8s_cluster() {
  local temp_dir=${1:-/tmp}
  local cluster_name=${2:-freshcloud}
  cat > "$temp_dir/cluster.yaml" <<EOF
CLUSTER_NAME: $cluster_name
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
BASTION_HOST_ENABLED: false
EOF
  tanzu cluster create -f "$temp_dir/cluster.yaml"
  printf "Workload cluster file saved to: %s\n" "$temp_dir/cluster.yaml"
}

function tanzu_aws_delete_k8s_cluster() {
  local cluster_name=${1:-freshcloud}
  tanzu cluster delete "$cluster_name"
}

if [ -z "$AWS_ACCESS_KEY_ID" ] || [ -z "$AWS_SECRET_ACCESS_KEY" ]; then
  printf "Please set the following environment variables in .env_development.sh under root directory:\n"
  printf "AWS_ACCESS_KEY_ID\n"
  printf "AWS_SECRET_ACCESS_KEY\n"
  exit 1
fi

temp_dir=$(mktemp -d -t cluster-XXXXXXXXXX)

if [ "$1" == 'mgmt' ]; then
  if [ "$2" == 'delete' ]; then
    tanzu_aws_delete_k8s_mgmt_cluster "$3"
  else
    tanzu_aws_create_k8s_mgmt_cluster "$temp_dir" "$2"
  fi
  exit 0
fi

if [ "$1" == 'delete' ]; then
    tanzu_aws_delete_k8s_cluster "$2"
else
  tanzu_aws_create_k8s_cluster "$temp_dir" "$1"
fi
