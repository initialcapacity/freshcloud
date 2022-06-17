mkdir -p .freshcloud

helm repo add gitlab https://charts.gitlab.io/
helm repo update

helm upgrade --install --create-namespace --namespace gitlab gitlab gitlab \
        --repo https://charts.gitlab.io/ \
        --set global.hosts.domain={{index . "DOMAIN"}} \
        --set certmanager.install=false \
        --set global.ingress.configureCertmanager=false \
        --set nginx-ingress.enabled=false \
        --set global.ingress.class=contour\
        --set global.ingress.provider=contour \
        --set global.ingress.annotations."cert-manager\.io/cluster-issuer"=letsencrypt-prod \
        --set global.ingress.annotations."kubernetes.io/ingress.class"=contour \
        --set global.ingress.annotations."ingress.kubernetes.io/force-ssl-redirect"=true \
        --set global.ingress.annotations."projectcontour.io/websocket-routes"='/' \
        --set global.ingress.annotations."projectcontour.io/response-timeout"='10m' \
        --set global.ingress.annotations."kubernetes.io/tls-acme"='true' \
        --set gitlab-runner.runners.privileged=true \
        --set gitlab.webservice.ingress.tls.secretName=gitlab-webservice-tls \
        --set registry.ingress.tls.secretName=gitlab-registry-tls \
        --set minio.ingress.tls.secretName=gitlab-minio-tls \
        --set gitlab.webservice.replicaCount=2
