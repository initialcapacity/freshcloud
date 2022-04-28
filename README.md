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
```

Source the file.

```bash
source .env_development
```

Run via go run (for now).

```base
go run cmd/freshctl.go
```
