#!/usr/bin/env bash
#
# Create or Delete a jumpbox in GCP.

source ../../.env_development.sh

NETWORK_NAME="fc-net"
SUBNET_NAME="fc-subnet"
CIDR='172.16.0.0/24'
VMIP='172.16.0.10'

function create_stack() {

  VM_NAME=$1

  gcloud --quiet compute networks create ${NETWORK_NAME} --subnet-mode custom

  gcloud --quiet compute networks subnets create ${SUBNET_NAME} \
    --network ${NETWORK_NAME} \
    --range ${CIDR}

  gcloud --quiet compute firewall-rules create ${NETWORK_NAME}-allow-internal \
    --allow tcp,udp,icmp \
    --network ${NETWORK_NAME} \
    --source-ranges ${CIDR}

  gcloud --quiet compute firewall-rules create ${NETWORK_NAME}-allow-external \
    --allow tcp:22,icmp \
    --network ${NETWORK_NAME} \
    --source-ranges 0.0.0.0/0

  gcloud --quiet compute firewall-rules list --filter="network:${NETWORK_NAME}"

  gcloud --quiet compute addresses create ${NETWORK_NAME} \
    --region ${REGION}

  gcloud --quiet compute addresses list --filter="name=('${NETWORK_NAME}')"

  gcloud --quiet compute instances create $VM_NAME \
    --async \
    --boot-disk-size 200GB \
    --zone ${ZONE} \
    --image-family ubuntu-1804-lts \
    --image-project ubuntu-os-cloud \
    --machine-type n1-standard-1 \
    --private-network-ip ${VMIP} \
    --scopes compute-rw,storage-ro,service-management,service-control,logging-write,monitoring \
    --subnet ${SUBNET_NAME} \
    --tags ${NETWORK_NAME}

  sleep 15
  echo "gcloud compute ssh --zone ${ZONE} $VM_NAME"
  # notes
  # gcloud compute scp --recurse ~/.ssh/ jumpbox-00:~/
}

function delete_stack() {

  VM_NAME=$1

  gcloud --quiet compute instances delete --zone ${ZONE} $VM_NAME
  gcloud compute firewall-rules delete ${NETWORK_NAME}-allow-internal
  gcloud --quiet compute firewall-rules delete ${NETWORK_NAME}-allow-external
  gcloud --quiet compute addresses delete ${NETWORK_NAME}
  gcloud --quiet compute networks subnets delete ${SUBNET_NAME}
  gcloud --quiet compute networks delete ${NETWORK_NAME}
}

if [ "$1" = '--create' ]; then
  create_stack $2
elif [ "$1" = '--delete' ]; then
  delete_stack $2
else
  echo -e "Create or delete a jumpbox in GCP.\n"
  echo -e "$0 --create <vm-name>
             --delete <vm-name>\n"
fi

