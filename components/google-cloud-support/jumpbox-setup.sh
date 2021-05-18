#!/usr/bin/env bash
#
# Setup an Ubuntu jumpbox with the tools we need to operate this repository.

if [ "$EUID" -ne 0 ]; then
  echo "Please run as: sudo $0"
  exit 1
fi

if [ -z ${GCP_PROJECT_NAME} ]; then
  echo "Please configure and source ../bin/env.sh"
  exit 1
fi

# Setup gcloud
gcloud auth login
gcloud config set project ${GCP_PROJECT_NAME}

# Setup Docker
snap install docker
addgroup --system docker
adduser $USER docker
newgrp docker
snap disable docker
snap enable docker

# Install utils
snap install kubectl --classic
sudo snap install helm --classic
snap install yq
add-apt-repository ppa:cncf-buildpacks/pack-cli
apt-get update
apt-get install pack-cli

wget https://github.com/vmware-tanzu/kpack-cli/releases/download/v0.2.0/kp-linux-0.2.0 -O /usr/local/bin/kp
chmod 755 /usr/local/bin/kp

# Need to install fly to talk to Concourse.
