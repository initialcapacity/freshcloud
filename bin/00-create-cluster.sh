#!/usr/bin/env bash

# Load variables from env.sh
source env.sh

# Set default compute zone
gcloud config set compute/zone $ZONE

# Create the Kubernetes cluster with auto-upgrade enabled for nodes
echo "Creating Kubernetes cluster named $CLUSTER_NAME with node auto-upgrades enabled..."
gcloud container clusters create $CLUSTER_NAME \
    --zone $ZONE \
    --num-nodes $NUM_NODES \
    --machine-type $MACHINE_TYPE \
    --enable-autoupgrade

# Getting credentials for kubectl
echo "Getting credentials for cluster $CLUSTER_NAME..."
gcloud container clusters get-credentials $CLUSTER_NAME

echo "Cluster $CLUSTER_NAME created with node auto-upgrades enabled."

