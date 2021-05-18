### Install aws and eksctl 
```console
  $ brew install aws
  $ brew tap weaveworks/tap
  $ brew install weaveworks/tap/eksctl
  $ eksctl version
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
```

### Delete existing cluster
```console
  $ bash aws-cloud-support.sh delete
```
