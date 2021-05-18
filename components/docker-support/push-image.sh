#!/usr/bin/env bash
#
# Push an application image to our registry.

source ../../.env_development.sh

if [ -z $1 ]; then
  echo "$0 <app-name> <image-name>"
  exit 1
fi

APP_NAME=$1
IMAGE=$2

curl --user "admin:${PASSWD}" -X POST \
         https://registry.${DOMAIN}/api/v2.0/projects \
         -H "Content-type: application/json" --data \
         '{ "project_name": "'${APP_NAME}'", "metadata":
          { "auto_scan": "true", "enable_content_trust":
            "false", "prevent_vul": "false", "public":
            "true", "reuse_sys_cve_whitelist": "true",
            "severity": "high" }
          }'

docker login -u admin -p ${PASSWD} https://registry.${DOMAIN}
docker tag $IMAGE registry.${DOMAIN}/${APP_NAME}/${IMAGE}
docker push registry.${DOMAIN}/${APP_NAME}/${IMAGE}
