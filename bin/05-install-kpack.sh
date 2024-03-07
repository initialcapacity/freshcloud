#!/usr/bin/env bash
#
# Install KPack

source env.sh
source utils.sh

function check_docker_service() {

  docker info &> /dev/null
  if [ $? != 0 ]; then
    echo "Please start the docker service so we can build containers."
    exit 127
  fi
}

function kube_install_kpack() {

   kubectl apply -f https://github.com/buildpacks-community/kpack/releases/download/v0.13.2/release-0.13.2.yaml
}

function pull_pipeline_images() {

 echo "Pulling pipeline images"
 docker login -u admin -p ${HARBOR_PASSWD} https://registry.${SUB_DOMAIN}

 # Kubectl container image
 docker pull quay.io/tfgco/kubectl
 docker tag quay.io/tfgco/kubectl registry.${SUB_DOMAIN}/concourse-images/kubectl-docker:latest 
 docker push registry.${SUB_DOMAIN}/concourse-images/kubectl-docker:latest

 # Container: pipeline talks to kpack
 docker pull gcr.io/cf-build-service-public/concourse-kpack-resource:1.0
 docker tag gcr.io/cf-build-service-public/concourse-kpack-resource:1.0 registry.${SUB_DOMAIN}/concourse-images/concourse-kpack-resource:1.0
 docker push registry.${SUB_DOMAIN}/concourse-images/concourse-kpack-resource:1.0
}

function create_cluster_stack_kpack() {

  cat <<EOF | kubectl apply -f -
apiVersion: kpack.io/v1alpha2
kind: ClusterStack
metadata:
  name: base
spec:
  id: "io.buildpacks.stacks.jammy"
  buildImage:
    image: "paketobuildpacks/build-jammy-base"
  runImage:
    image: "paketobuildpacks/run-jammy-base"
EOF
}


function create_cluster_store_kpack() {

  cat <<EOF | kubectl apply -f -
apiVersion: kpack.io/v1alpha2
kind: ClusterStore
metadata:
  name: default
spec:
  sources:
  - image: gcr.io/paketo-buildpacks/java
  - image: gcr.io/paketo-buildpacks/nodejs
  - image: gcr.io/paketo-buildpacks/go
EOF
}

check_docker_service
kube_install_kpack
pull_pipeline_images
wait_for_ready kpack
create_cluster_stack_kpack
create_cluster_store_kpack
