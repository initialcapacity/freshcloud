[![Build results](https://github.com/initialcapacity/freshcloud/workflows/build/badge.svg)](https://github.com/initialcapacity/freshcloud/actions)
[![codecov](https://codecov.io/gh/initialcapacity/freshcloud/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/freshcloud)
[![CodeQL](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml/badge.svg)](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml)

# Fresh Cloud

## A modern application platform.

The FreshCloud article with step-by-step instructions in Bash can be found here [freshcloud.com](https://www.freshcloud.com).

## About the golang binary

Fresh Cloud now includes a golang binary for managing infrastructure and applications on Kubernetes.

The current `freshclt` binary support Google's Cloud Platform.

## Getting started

Download the codebase from GitHub to your local machine and install the following prerequisites.

* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Google Cloud](https://cloud.google.com/sdk)

```bash
cd /home/user/workspace/
git clone git@github.com:initialcapacity/freshcloud.git
cd freshcloud
```

Optionally, build and install the Fresh Cloud command line interface; `freshctl`.

```bash
go build cmd/freshctl.go
go install cmd/freshctl.go
```

Check the `freshctl` installation path as needed.

```bash
go list -f '{{.Target}}' cmd/freshctl.go
```

## Management cluster

Configure your google cloud project.

```bash
gcloud config set project ${GCP_PROJECT_ID}
```

Ensure the project was set correctly.

```bash
gcloud projects describe ${GCP_PROJECT_ID}
```

Create a `.env_infra.sh` file similar to the below.

```bash
export GCP_PROJECT_ID=aProjectId
export GCP_ZONE=aZone
export GCP_CLUSTER_NAME=aClusterName
export DOMAIN=aDomain
export EMAIL_ADDERESS=anEmail
export PASSWORD=aPassword
```

Next, source environment the file.

```bash
source .env_infra.sh
```

Then, run each command via go run or the installed binary, `~/go/bin/freshctl`, to create a fresh cloud management cluster.

_Note_ adding the `-e` flag will execute the command.

```base
go run cmd/freshctl.go clusters gcp enable-services
go run cmd/freshctl.go clusters gcp create
go run cmd/freshctl.go clusters gcp list
go run cmd/freshctl.go services contour
```

Create a DNS entry for your load balancer. As needed, re-run the below command to show your ip address.

```base
kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}'
```

Continue installing management cluster services.

```base
go run cmd/freshctl.go services cert-manager
go run cmd/freshctl.go services harbor
go run cmd/freshctl.go services concourse
go run cmd/freshctl.go services kpack
```

Confirm the management cluster services are deployed.
* [Harbor](https://registry.{your-domain})
* [Concourse](https://ci.{your-domain})

That's a wrap for now.
Continue to our [Applications](README_APPS.md) to learn how to deploy application pipelines to your new cluster.
