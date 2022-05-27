### Install azure-cli on OSX

```console
  $ brew install azure-cli
```

### Configure azure-cli access

```console
  $ az login
```

### Create and then edit .env_development.sh and enter the below values

```console
  $ touch .env_development.sh
  $ vim .env_development.sh

# Deployment defaults
  export DOMAIN='cluster-name.domain.cc'      # Enter the DNS subdomain
  export EMAIL_ADDRESS='nobody@gmail.com'     # E-Mail for CERT registration confirmation
  export PASSWD='my-pass'                     # Password that will be used throughout the project

# Azure Resource Group 
  export K8S_CLUSTER_NAME='cluster-name'
  export REGION='westus'
```

### Delete existing Azure Resource Group

```console
  $ bash azure-cloud-support.sh delete
```
