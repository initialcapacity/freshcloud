#!/usr/bin/env bash
#
# Create a k8s cluster in AKS

source ../../.env_development.sh

function aks_create_rg() {
  echo "Creating Resource Group: ${K8S_CLUSTER_NAME}-rg"
  az group create --name ${K8S_CLUSTER_NAME}-rg --location ${REGION}
}

function aks_create_cluster() {
  echo "Creating AKS: ${K8S_CLUSTER_NAME}"
  az aks create --resource-group ${K8S_CLUSTER_NAME}-rg --name ${K8S_CLUSTER_NAME} \
    --node-count 1 --generate-ssh-keys
}

function aks_cluster_auth() {
  
  echo "Getting AKS Auth: ${K8S_CLUSTER_NAME}"
  az aks get-credentials --name ${K8S_CLUSTER_NAME} --overwrite-existing --resource-group ${K8S_CLUSTER_NAME}-rg
}

function aks_rg_delete() {
  echo "Deleting AKS Resource Group: ${K8S_CLUSTER_NAME}-rg"
  az group delete --name ${K8S_CLUSTER_NAME}-rg --yes --no-wait
}

if [ "$1" == 'delete' ]; then
  aks_rg_delete
else
  aks_create_rg
  aks_create_cluster
  aks_cluster_auth
fi
