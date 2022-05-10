[![Build results](https://github.com/initialcapacity/freshcloud/workflows/build/badge.svg)](https://github.com/initialcapacity/freshcloud/actions)
[![codecov](https://codecov.io/gh/initialcapacity/freshcloud/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/freshcloud)
[![CodeQL](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml/badge.svg)](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml)

# Fresh Cloud

## A modern application platform.

The Fresh Cloud article with step-by-step instructions in *Bash* can be found here [freshcloud.com](https://www.freshcloud.com).

## Getting started

Fresh Cloud now includes a golang binary for managing services and applications on Kubernetes.
The current `freshctl` binary supports Google's Cloud Platform.

Download and install the
[latest darwin-arm64 release](https://github.com/initialcapacity/freshcloud/releases/latest/download/freshctl-darwin-arm64)
of the binary. Additional golang operating systems and architectures can be found in the
GitHub action [build artifacts](https://github.com/initialcapacity/freshcloud/actions/workflows/build.yml).

```bash
curl -L https://github.com/initialcapacity/freshcloud/releases/download/v202205100303/freshctl-darwin-arm64 \
  -o /usr/local/bin/freshctl
chmod 755 /usr/local/bin/freshctl
```

Install the following prerequisites.
* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Google Cloud](https://cloud.google.com/sdk)

## Management cluster

Configure your google cloud project.

```bash
gcloud config set project ${GCP_PROJECT_ID}
```

Ensure the project was set correctly.

```bash
gcloud projects describe ${GCP_PROJECT_ID}
```

Create a directory for you project.

```bash
cd /home/{user}/workspace/
mkdir freshcloud-example
cd freshcloud-example
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

Then, run each command to create a fresh cloud management cluster.

_Note_ adding the `-e` flag will execute the command.

```base
freshctl clusters gcp enable-services
freshctl clusters gcp create
freshctl clusters gcp list
freshctl services contour
```

Create a DNS entry for your load balancer. As needed, re-run the below command to show your ip address.

```base
kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}'
```

Continue installing management cluster services.

```base
freshctl services cert-manager
freshctl services harbor
freshctl services concourse
freshctl services kpack
```

Confirm the management cluster services are deployed.
* [Harbor](https://registry.{your-domain})
* [Concourse](https://ci.{your-domain})

That's a wrap for now.
Continue to our [Applications](APPLICATIONS.md) page to learn how to deploy application pipelines to your new cluster.
