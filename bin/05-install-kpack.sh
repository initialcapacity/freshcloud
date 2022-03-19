#!/usr/bin/env bash
#
# Install KPack

source ../.env_development.sh
source ../components/kubernetes-support/kubectl-support.sh

function check_docker_service() {

  docker info &> /dev/null
  if [ $? != 0 ]; then
    echo "Please start the docker service so we can build containers."
    exit 127
  fi
}

function kube_install_kpack() {

   kubectl apply -f https://github.com/pivotal/kpack/releases/download/v0.4.2/release-0.4.2.yaml
}

function build_docker_container() {

 echo "Building images: kubectl-docker:latest"
 docker login -u admin -p ${PASSWD} https://registry.${DOMAIN}

 # Container: pipeline talks to k8s
 docker build --platform linux/amd64 --rm -t registry.${DOMAIN}/concourse-images/kubectl-docker:latest .
 docker push registry.${DOMAIN}/concourse-images/kubectl-docker:latest

 # Container: pipeline talks to kpack
 docker pull gcr.io/cf-build-service-public/concourse-kpack-resource:1.0
 docker tag gcr.io/cf-build-service-public/concourse-kpack-resource:1.0 registry.${DOMAIN}/concourse-images/concourse-kpack-resource:1.0
 docker push registry.${DOMAIN}/concourse-images/concourse-kpack-resource:1.0
}

function create_cluster_stack_kpack() {

  cat <<EOF | kubectl apply -f -
apiVersion: kpack.io/v1alpha1
kind: ClusterStack
metadata:
  name: base
spec:
  id: "heroku-20"
  buildImage:
    image: "heroku/pack:20-build"
  runImage:
    image: "heroku/pack:20"
EOF
}


function create_cluster_store_kpack() {

  cat <<EOF | kubectl apply -f -
apiVersion: kpack.io/v1alpha1
kind: ClusterStore
metadata:
  name: default
spec:
  sources:
  - image: heroku/buildpacks:20
EOF
}

check_docker_service
build_docker_container
kube_install_kpack
wait_for_ready kpack
create_cluster_stack_kpack
create_cluster_store_kpack

