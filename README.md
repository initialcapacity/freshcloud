[![Build results](https://github.com/initialcapacity/freshcloud/workflows/build/badge.svg)](https://github.com/initialcapacity/freshcloud/actions)
[![codecov](https://codecov.io/gh/initialcapacity/freshcloud/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/freshcloud)
[![CodeQL](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml/badge.svg)](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml)

# Fresh Cloud

## A modern application platform.

The main FreshCloud article with step-by-step instructions can be found here [freshcloud.com](https://www.freshcloud.com).

## About the golang binary

Fresh Cloud now includes a golang binary for managing infrastructure and application clusters.

The current `freshclt` binary support Google's Cloud Platform.

## Google Cloud

To get started, configure the Google Cloud CLI for your project.

```bash
gcloud config set project ${GCP_PROJECT_ID}
```

Check the project.

```bash
gcloud projects describe ${GCP_PROJECT_ID}
```

Update components as needed.

```bash
gcloud components update
```

## Management cluster

Create a `.env_infra.sh` file similar to the below.

```bash
export GCP_PROJECT_ID=aProjectId
export GCP_ZONE=aZone
export GCP_CLUSTER_NAME=aClusterName

export DOMAIN=aDomain
export EMAIL_ADDERESS=anEmail
export PASSWORD=aPassword
```

Next, source the file.

```bash
source .env_infra.sh
```

Then, run each command via go run (for now) to create a fresh cloud management cluster. Adding the `-e` flag will
execute the command.

```base
go run cmd/freshctl.go clusters gcp enable-services
go run cmd/freshctl.go clusters gcp create
go run cmd/freshctl.go clusters gcp list
go run cmd/freshctl.go services contour
go run cmd/freshctl.go services cert-manager
go run cmd/freshctl.go services harbor
go run cmd/freshctl.go services concourse
go run cmd/freshctl.go services kpack
```

## Application cluster

Similar to the above, create a `.env_app.sh` file for you application.

```bash
export GCP_PROJECT_ID=aProjectId
export GCP_ZONE=aZone
export GCP_CLUSTER_NAME=anAppClusterName

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
go run cmd/freshctl.go clusters gcp list

go run cmd/freshctl.go services contour
go run cmd/freshctl.go services cert-manager

go run cmd/freshctl.go applications push
go run cmd/freshctl.go applications deploy

go run cmd/freshctl.go pipelines push-build-image
go run cmd/freshctl.go pipelines deploy
```

That's a wrap for now.
