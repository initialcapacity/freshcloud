### Install aws cli

```console
  $ brew install aws
```

### Configure awscli access
```console
  $ aws configure
```

### Create and then edit .env_development.sh and enter the below values
```console
  $ touch .env_development.sh
  $ vim .env_development.sh

# Deployment defaults
  export DOMAIN='cluster-name.domain.cc'      # Enter the DNS subdomain
  export EMAIL_ADDRESS='nobody@gmail.com'     # E-Mail for CERT registration confirmation
  export PASSWD='my-pass'                     # Password that will be used throughout the project

# AWS cluster Info
  export K8S_CLUSTER_NAME='cluster-name'
  export REGION='us-west-2'
  export NODE_AZ='us-west-2a'
  export AWS_ACCESS_KEY_ID='xxx'
  export AWS_SECRET_ACCESS_KEY='xxx'
  export AWS_SESSION_TOKEN='xxx'       # Optional, if using temporary security credentials
```

### Install tanzu cli

Follow these [instructions](https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.3/vmware-tanzu-kubernetes-grid-13/GUID-install-cli.html) to download and install the tanzu cli bundle.


### Create a management cluster
```console
  $ bash tanzu-aws-support.sh mgmt
```

### Login to management cluster
```console
  $ tanzu login
```

### Create a cluster
```console
  $ bash tanzu-aws-support.sh
```

### Login to workload cluster by getting the kubeconfig
```console
  $ tanzu cluster kubeconfig get freshcloud --admin
```

### Delete existing cluster
```console
  $ bash tanzu-aws-support.sh delete
```

### Delete existing management cluster
```console
  $ bash tanzu-aws-support.sh mgmt delete
```
