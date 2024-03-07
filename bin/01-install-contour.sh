#!/usr/bin/env bash
#
# Install Contour

source env.sh

function helm_install_contour() {
  kubectl create namespace ${INGRESS_NS}
  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo update
  helm install ingress bitnami/contour -n ${INGRESS_NS}
  if [ $? != 0 ]; then
   echo "Failed to install Contour. Bummer"
   exit 1
  fi
}

function get_load_balancer_ip() {

  echo "Waiting to get the load-balancer IP."
  while true; do
    if [ -z "$LB" ]; then
      LB=$(kubectl describe svc ingress-contour-envoy --namespace ${INGRESS_NS} | grep Ingress | awk '{print $3}')
        sleep 3;
    else
      echo "Create a DNS A for *.$SUB_DOMAIN to $LB"
      break
    fi
  done
}

helm_install_contour
get_load_balancer_ip
