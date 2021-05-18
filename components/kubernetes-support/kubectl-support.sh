#!/usr/bin/env bash

function wait_for_ready() {

  NS=$1
  echo "Waiting for pods in $NS to become ready."
  while true; do
    STATUS=$(kubectl get pods -n $NS | egrep -v 'Running|NAME|Completed')
    if [ -z "$STATUS" ]; then
      break
    fi
  done
  echo "All pods are running."
}