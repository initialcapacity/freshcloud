#!/usr/bin/env bash

__DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" || exit; pwd)"

die() {
    2>&1 echo "$@"
    exit 1
}

# shellcheck disable=SC1091
source "${__DIR}/../.env_development.sh" || die "Could not find '.env_development.sh' in root directory"

if [ -z "$1" ]; then
  echo "$0 <gcp|aws|azure|tanzu>"
fi

if [ "$1" = 'gcp' ]; then
  cd "${__DIR}"/../components/google-cloud-support/ && ./google-cloud-support.sh
elif [ "$1" = 'aws' ]; then
  cd "${__DIR}"/../components/aws-support/ && ./aws-cloud-support.sh
elif [ "$1" = 'azure' ]; then
  cd "${__DIR}"/../components/azure-support/ && ./azure-cloud-support.sh
elif [ "$1" = 'tanzu' ]; then
  cd "${__DIR}"/../components/tanzu-support/ || die "Could not cd into ../components/tanzu-support from $PWD"
  if [ "$2" == 'mgmt' ]; then
    ./tanzu-aws-support.sh 'mgmt'
  else
    ./tanzu-aws-support.sh
  fi
fi
