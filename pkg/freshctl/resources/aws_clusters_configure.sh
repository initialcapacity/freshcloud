aws eks --region {{index . "AWS_REGION" }} update-kubeconfig --name {{index . "K8S_CLUSTER_NAME" }}