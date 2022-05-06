gcloud container clusters get-credentials '{{index . "REGISTRY_CLUSTER_NAME"}}' \
  --project '{{index . "GCP_PROJECT_ID"}}' \
  --zone '{{index . "GCP_ZONE"}}' \
  --quiet
kubectl create namespace {{index . "APP_NAME"}}
kubectl create secret docker-registry {{index . "APP_NAME"}}-registry-credentials \
  --docker-username=admin \
  --docker-password={{index . "REGISTRY_PASSWORD"}} \
  --docker-server=registry.{{index . "REGISTRY_DOMAIN"}} \
  --namespace {{index . "APP_NAME"}}

envsubst < {{index . "APP_PIPELINE_CONFIGURATION_PATH"}} > .freshcloud/{{index . "APP_NAME"}}-pipeline-configuration.yaml
kubectl apply -f .freshcloud/{{index . "APP_NAME"}}-pipeline-configuration.yaml

envsubst < {{index . "APP_PIPELINE_PATH"}} > .freshcloud/{{index . "APP_NAME"}}-pipeline.yaml

NAME=$(kubectl get secrets -n {{index . "APP_NAME"}} |grep {{index . "APP_NAME"}}-service-account-token|awk '{print $1}')
CA=$(kubectl get secret/${NAME} -n {{index . "APP_NAME"}} -o jsonpath='{.data.ca\.crt}')
TOKEN=$(kubectl get secret/${NAME} -n {{index . "APP_NAME"}} -o jsonpath='{.data.token}' | base64 --decode)
SERVER=$(kubectl cluster-info|head -1|awk '{print $NF}'|sed -r "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g")

cat <<EOF > .freshcloud/{{index . "APP_NAME"}}-pipeline-parameters.yaml
service-account-key: random-string
domain: {{index . "REGISTRY_DOMAIN"}}
kubeconfig: |
  apiVersion: v1
  kind: Config
  clusters:
  - name: {{index . "REGISTRY_CLUSTER_NAME"}}
    cluster:
      certificate-authority-data: ${CA}
      server: ${SERVER}
  contexts:
  - name: default-context
    context:
      cluster: {{index . "REGISTRY_CLUSTER_NAME"}}
      namespace: {{index . "APP_NAME"}}
      user: default-user
  current-context: default-context
  users:
  - name: default-user
    user:
      token: ${TOKEN}
EOF

fly login -c https://ci.{{index . "REGISTRY_DOMAIN"}} -u admin -p {{index . "REGISTRY_PASSWORD"}} -t {{index . "REGISTRY_CLUSTER_NAME"}}
echo y | fly -t {{index . "REGISTRY_CLUSTER_NAME"}} set-pipeline -p build-{{index . "APP_NAME"}} \
  -c .freshcloud/{{index . "APP_NAME"}}-pipeline.yaml \
  -l .freshcloud/{{index . "APP_NAME"}}-pipeline-parameters.yaml
fly -t {{index . "REGISTRY_CLUSTER_NAME"}} unpause-pipeline -p build-{{index . "APP_NAME"}}