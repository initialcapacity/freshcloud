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

Create a service account for you applications. _Note_, for the moment, the below requires application clusters
to reside in the same google cloud project as the management cluster.

```base
go run cmd/freshctl.go clusters gcp create-service-account
```

Confirm the management cluster services are deployed.
* [Harbor](https://registry.{your-domain})
* [Concourse](https://ci.{your-domain})

## Application cluster

Similar to the above, create a `.env_app.sh` file for you application.
* `REGISTRY_*` variables should match the management cluster above. 
* `GCP_SERVICE_ACCOUNT_JSON` should be located in your lcoal `.freshcloud` directory.
* `APP_CONFIGURATION_PATH` should contain kubernetes objects. 
* `APP_PIPELINE_*_PATH` should contain concourse pipeline files. 

See our [Kotlin-ktor-starter](https://github.com/initialcapacity/kotlin-ktor-starter/tree/main/deployments) for an example. 

```bash
export GCP_PROJECT_ID=aProjectId
export GCP_ZONE=aZone
export GCP_CLUSTER_NAME=anAppClusterName
export GCP_SERVICE_ACCOUNT_JSON=.freshcloud/your-service-account.json

export REGISTRY_DOMAIN=aRegistryDomain
export REGISTRY_PASSWORD=aRegistryPassword
export REGISTRY_CLUSTER_NAME=aRegistryClusterName

export DOMAIN=aDomain
export EMAIL_ADDERESS=anEmail

export APP_NAME=anAppName
export APP_IMAGE_NAME=anAppImageName
export APP_CONFIGURATION_PATH=anApplicationConfigurationPath
export APP_PIPELINE_CONFIGURATION_PATH=anAppPipelineConfigurationPath
export APP_PIPELINE_PATH=anAppPipelinePath
```

Next, source the file.

```bash
source .env_app.sh
```

Run each command to deploy your application. Adding the `-e` flag will execute the command.

```bash
go run cmd/freshctl.go clusters gcp create
go run cmd/freshctl.go services contour
```

Create a DNS entry for your load balancer similar to your management cluster.

```bash
go run cmd/freshctl.go services cert-manager
go run cmd/freshctl.go pipelines push-build-image
go run cmd/freshctl.go pipelines deploy
```

Cleaning up

```bash
docker rm -f $(docker ps -a -q)
docker volume rm -f $(docker volume ls -q)
docker system prune -a -f
```

That's a wrap for now.
