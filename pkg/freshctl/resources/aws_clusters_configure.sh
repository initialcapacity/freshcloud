aws eks --region {{index . "AWS_REGION" }} update-kubeconfig --name {{index . "AWS_CLUSTER_NAME" }}