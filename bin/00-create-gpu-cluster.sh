#!/usr/bin/env bash

# Load variables from env.sh
source env.sh

# Check required variables
if [[ -z "$ZONE" || -z "$CLUSTER_NAME" || -z "$MACHINE_TYPE" ]]; then
    echo "Error: Required environment variables ZONE, CLUSTER_NAME, or MACHINE_TYPE are not set."
    exit 1
fi

# Set default compute zone
gcloud config set compute/zone "$ZONE"

# Create the Kubernetes cluster without GPUs in the default node pool
echo "Creating Kubernetes cluster named $CLUSTER_NAME..."
if gcloud container clusters create "$CLUSTER_NAME" \
    --zone "$ZONE" \
    --num-nodes $NUM_NODES \
    --machine-type "$MACHINE_TYPE" \
    --enable-autoupgrade \
    --enable-autorepair \
    --scopes cloud-platform; then
    echo "Cluster $CLUSTER_NAME created successfully."
else
    echo "Failed to create cluster $CLUSTER_NAME."
    exit 1
fi

# Add a node pool with GPU capabilities
echo "Adding a GPU-enabled node pool to the cluster..."
if gcloud container node-pools create "gpu-pool" \
    --cluster "$CLUSTER_NAME" \
    --zone "$ZONE" \
    --machine-type "$MACHINE_TYPE" \
    --accelerator "type=nvidia-tesla-t4,count=1" \
    --num-nodes 1 \
    --enable-autoupgrade \
    --enable-autorepair; then
    echo "GPU-enabled node pool added successfully."
    kubectl apply -f \
    https://raw.githubusercontent.com/GoogleCloudPlatform/container-engine-accelerators/master/nvidia-driver-installer/cos/daemonset-preloaded.yaml
else
    echo "Failed to add a GPU-enabled node pool to the cluster $CLUSTER_NAME."
    exit 1
fi

# Getting credentials for kubectl
echo "Getting credentials for cluster $CLUSTER_NAME..."
if gcloud container clusters get-credentials "$CLUSTER_NAME"; then
    echo "Credentials for cluster $CLUSTER_NAME obtained successfully."
else
    echo "Failed to get credentials for cluster $CLUSTER_NAME."
    exit 1
fi

# List GPU nodes
kubectl get nodes "-o=custom-columns=NAME:.metadata.name,GPU:.status.allocatable.nvidia\.com/gpu"
