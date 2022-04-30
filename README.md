[![Build results](https://github.com/initialcapacity/freshcloud/workflows/build/badge.svg)](https://github.com/initialcapacity/freshcloud/actions)
[![codecov](https://codecov.io/gh/initialcapacity/freshcloud/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/freshcloud)

# Fresh Cloud

## A modern application platform.

README.md notes can be found here [freshcloud.com](https://www.freshcloud.com).

# freshctl go-binary (wip)

Create a `.env_development.sh` file similar to the below.

```bash
export GCP_PROJECT_ID=aProjectId
export GCP_ZONE=aZone
export GCP_CLUSTER_NAME=aClusterName
export DOMAIN=aDomain
export EMAIL_ADDERESS=anEmail
```

Source the file.

```bash
source .env_development
```

Configure the Google Cloud CLI for your project.

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

Run via go run (for now).

```base
go run cmd/freshctl.go clusters gcp enable-services
go run cmd/freshctl.go clusters gcp create
go run cmd/freshctl.go clusters gcp list
go run cmd/freshctl.go services contour
go run cmd/freshctl.go services cert-manager
```
