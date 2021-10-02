#!/bin/bash
#
# Update Route53 with our LB addy

source ../../.env_development.sh

if [ -z ${ZONE_ID} ]; then
  echo "Please set ZONE_ID. in .env_development.sh"
  exit 127
fi

LB=$(kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}')

echo "Setting: *.${DOMAIN} to ${LB}"
aws route53 change-resource-record-sets --hosted-zone-id ${ZONE_ID} \
  --change-batch '{ "Comment": "CREATE/DELETE/UPSERT a record ", "Changes": [{ "Action": "UPSERT", "ResourceRecordSet": { "Name": "'*.${DOMAIN}'", "Type": "A", "TTL": 300, "ResourceRecords": [{ "Value": "'${LB}'"}] }}] }' 
sleep 30

# -- Test
VAR=$(host app.${DOMAIN} | awk '{print $NF}')

if [ "${VAR}" != "${LB}" ]; then
  echo "Update failed."
else
  echo "DNS looks good."
fi
