#!/usr/bin/env bash
#
# Install Ollama

source env.sh

mkdir -p build/ollama/

function write_ollama_manifest() {

   cat <<EOF> build/ollama/ollama.${SUB_DOMAIN}.yml
---
apiVersion: v1
kind: Namespace
metadata:
  name: ollama
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ollama-models-pvc
  namespace: ollama
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 100Gi  # Adjust size as needed
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: ollama
  name: ollama
  namespace: ollama
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ollama
  template:
    metadata:
      labels:
        app: ollama
    spec:
      containers:
      - image: ollama/ollama
        name: ollama
        imagePullPolicy: Always
        ports:
        - name: http
          containerPort: 11434
          protocol: TCP
        volumeMounts:
        - name: models-volume
          mountPath: "/models"
        env:
        - name: OLLAMA_MODELS
          value: "/models"
        resources:
          requests:
            nvidia.com/gpu: 1 # Requesting 1 GPU
          limits:
            nvidia.com/gpu: 1 # Setting a limit of 1 GPU
      volumes:
      - name: models-volume
        persistentVolumeClaim:
          claimName: ollama-models-pvc
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: ollama
  name: ollama
  namespace: ollama
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 11434
  selector:
    app: ollama
  sessionAffinity: None
  type: ClusterIP
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ollama
  namespace: ollama
  labels:
    app: ollama
  annotations:
    cert-manager.io/cluster-issuer: letsencrypt-prod
    kubernetes.io/ingress.class: contour
    ingress.kubernetes.io/force-ssl-redirect: "true"
    projectcontour.io/websocket-routes: "/"
    projectcontour.io/response-timeout: 10m
    kubernetes.io/tls-acme: "true"
spec:
  rules:
  - host: ollama.${SUB_DOMAIN}
    http:
      paths:
      - backend:
          service:
            name: ollama
            port:
              number: 80
        pathType: ImplementationSpecific
  tls:
  - hosts:
    - ollama.${SUB_DOMAIN}
    secretName: ollama-cert

EOF

kubectl apply -f build/ollama/ollama.${SUB_DOMAIN}.yml
}

write_ollama_manifest
