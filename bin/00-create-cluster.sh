#!/usr/bin/env bash

source ../.env_development.sh

if [ -z $1 ]; then
  echo "$0 <gcp|aws|azure>"
fi

if [ "$1" = 'gcp' ]; then
  cd ../components/google-cloud-support/ && ./google-cloud-support.sh
elif [ "$1" = 'aws' ]; then
  cd ../components/aws-support/ && ./aws-cloud-support.sh
elif [ "$1" = 'azure' ]; then
  cd ../components/azure-support/ && ./azure-cloud-support.sh
fi
