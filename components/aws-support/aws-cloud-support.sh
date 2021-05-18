#!/usr/bin/env bash
#
# Create a k8s cluster in EKS

source ../../.env_development.sh

function aws_create_k8s_cluster() {

cat << EOF | eksctl create cluster -f -
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: ${K8S_CLUSTER_NAME}
  region: ${REGION}
nodeGroups:
  - name: controllers
    labels: { role: controllers }
    instanceType: m5.large
    desiredCapacity: 2
    iam:
      withAddonPolicies:
        certManager: true
        albIngress: true
    taints:
      controllers: "true:NoSchedule"
managedNodeGroups:
  - name: workers
EOF
}

function aws_delete_k8s_cluster() {
  eksctl delete cluster --name ${K8S_CLUSTER_NAME} --region ${REGION}
}

if [ "$1" == 'delete' ]; then
  aws_delete_k8s_cluster
else
  aws_create_k8s_cluster
fi
