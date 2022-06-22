#!/usr/bin/env bash
#
# Create a k8s cluster in GCP.

source ../../.env_development.sh

if [ -z "${GCP_PROJECT_NAME}" ]; then
    echo "Please setup .env_development.sh"
    exit 1
fi

function gcloud_enable_services() {

  # This is needed if it's a new gcloud project.
  G_APIS=( dns sqladmin compute storage-component datastore
	   cloudresourcemanager container )

  for service in ${G_APIS[@]}; do
    gcloud services enable ${service}.googleapis.com
  done
}


function gcp_create_k8s_cluster() {

  gcloud container clusters create \
      --project "${GCP_PROJECT_NAME}" \
      --machine-type "n1-standard-2" \
      ${K8S_CLUSTER_NAME} --zone ${ZONE} --num-nodes 6

}

function old_create() {

  gcloud beta container --project "${GCP_PROJECT_NAME}" \
      clusters create "${K8S_CLUSTER_NAME}" \
      --zone "${ZONE}" \
      --no-enable-basic-auth \
      --machine-type "n1-standard-2" \
      --image-type "COS" \
      --disk-type "pd-standard" \
      --disk-size "100" \
      --metadata disable-legacy-endpoints=true \
      --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
      --num-nodes "5" \
      --logging=SYSTEM,WORKLOAD --monitoring=SYSTEM --enable-ip-alias \
      --network "projects/${GCP_PROJECT_NAME}/global/networks/default" \
      --subnetwork "projects/${GCP_PROJECT_NAME}/regions/us-central1/subnetworks/default" \
      --default-max-pods-per-node "110" \
      --no-enable-master-authorized-networks \
      --addons HorizontalPodAutoscaling,HttpLoadBalancing \
      --enable-autoupgrade \
      --enable-autorepair
}

function configure_kubectl() {

  gcloud container clusters get-credentials ${K8S_CLUSTER_NAME} --zone ${ZONE}  --project ${GCP_PROJECT_NAME}
}

function gcp_delete_k8s_cluster() {

  echo "Deleting: ${K8S_CLUSTER_NAME} in ${ZONE}"
  gcloud container clusters delete ${K8S_CLUSTER_NAME} --zone ${ZONE}  --project ${GCP_PROJECT_NAME}
}


if [ "$1" == 'delete' ]; then
  gcp_delete_k8s_cluster
elif [ "$1" == 'auth' ]; then
  configure_kubectl
elif [ "$1" == 'enable-services' ]; then
  gcloud_enable_services
else
  echo "Creating: ${K8S_CLUSTER_NAME} --zone ${ZONE}"
  gcp_create_k8s_cluster
  configure_kubectl
fi
