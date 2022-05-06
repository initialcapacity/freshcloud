kubectl create namespace projectcontour
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install ingress bitnami/contour -n projectcontour --version 7.3.8
if [ $? != 0 ]; then
  echo "Failed to install Contour. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n projectcontour
sleep 5
LB=$(kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}')
echo "Create a DNS A for *.{{.Domain}} to $LB"
echo "Remove contour by running - kubectl delete ns projectcontour"