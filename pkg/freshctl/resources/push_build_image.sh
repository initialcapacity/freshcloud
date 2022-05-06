docker build --platform linux/amd64 --rm -t registry.{{index . "REGISTRY_DOMAIN"}}/concourse-images/cluster-mgmt:latest \
  https://raw.githubusercontent.com/initialcapacity/freshcloud/main/bin/Dockerfile
docker login -u admin -p {{index . "REGISTRY_PASSWORD"}} https://registry.{{index . "REGISTRY_DOMAIN"}}
docker push registry.{{index . "REGISTRY_DOMAIN"}}/concourse-images/cluster-mgmt:latest