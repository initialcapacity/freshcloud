az aks create --resource-group {{index . "AZURE_RESOURCE_GROUP"}} --name {{index . "AZURE_CLUSTER_NAME"}} \
  --node-count 1 --generate-ssh-keys