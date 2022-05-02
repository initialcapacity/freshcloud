mkdir -p .freshcloud
sha=$(curl -s --user "admin:{{.Password}}" -X GET \
  "https://registry.{{.RegistryDomain}}/api/v2.0/projects/{{.App}}/repositories/{{.Image}}/artifacts" \
  | jq -r '.[].digest'|head -1)
export IMAGE="registry.{{.RegistryDomain}}/{{.App}}/{{.Image}}@${sha}"
echo "Found image ${IMAGE}"
kubectl create namespace {{.App}}
envsubst < {{.AppConfigurationPath}} > .freshcloud/{{.App}}.yaml
kubectl apply -f .freshcloud/{{.App}}.yaml
echo "Deploy {{.App}} to https://{{.App}}.{{.AppDomain}}"
echo "Remove the app by running - kubectl delete ns {{.App}}"