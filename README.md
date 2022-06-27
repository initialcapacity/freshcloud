[![Build results](https://github.com/initialcapacity/freshcloud/workflows/build/badge.svg)](https://github.com/initialcapacity/freshcloud/actions)
[![codecov](https://codecov.io/gh/initialcapacity/freshcloud/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/freshcloud)
[![CodeQL](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml/badge.svg)](https://github.com/initialcapacity/freshcloud/actions/workflows/codeql.yml)

# Fresh Cloud

## A modern application platform.

Fresh Cloud is a lightweight golang binary or command line interface (CLI) that simplifies the creation of a modern
application platform on Kubernetes. The `freshctl` CLI is packed with a default, yet highly customizable set of scripts
for managing services, application pipelines and production applications on Kubernetes.

The original Fresh Cloud article with step-by-step instructions in Bash can be found
here [freshcloud.com](https://www.freshcloud.com).

## Getting started

Download and install the
[latest darwin-arm64 release](https://github.com/initialcapacity/freshcloud/releases/latest/download/freshctl-darwin-arm64)
of the binary. Additional golang operating systems and architectures can be found in the
GitHub action [build artifacts](https://github.com/initialcapacity/freshcloud/actions/workflows/build.yml).

```bash
curl -L https://github.com/initialcapacity/freshcloud/releases/latest/download/freshctl-darwin-arm64 -o /usr/local/bin/freshctl
chmod 755 /usr/local/bin/freshctl
```

Install the following prerequisites.

* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Google Cloud](https://cloud.google.com/sdk)

## Management cluster

The management cluster runs the image registry and continuous integration and delivery pipeline.
By default, Fresh Cloud uses [Harbor](https://goharbor.io) for the image registry
and [Concourse](https://concourse-ci.org) for the CI/CD pipeline.

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

export K8S_CLUSTER_NAME=aClusterName

export DOMAIN=aDomain
export EMAIL_ADDRESS=anEmail
export PASSWORD=aPassword
```

Next, source environment the file.

```bash
source .env_infra.sh
```

Then, run each command to create a fresh cloud management cluster.

_Note_ adding the `--execute` flag will execute the command below.

```base
freshctl clusters gcp enable-services
freshctl clusters gcp create
freshctl clusters gcp list
freshctl services add contour
```

Create a DNS entry for your load balancer. As needed, re-run the below command to show your ip address.

```base
kubectl describe svc ingress-contour-envoy --namespace projectcontour | grep Ingress | awk '{print $3}'
```

Continue installing management cluster services.

```base
freshctl services add cert-manager
freshctl services add harbor
freshctl services add concourse
freshctl services add kpack
```

Confirm the management cluster services are deployed.

* Harbor https://registry.{your-domain}
* Concourse https://ci.{your-domain}

That's a wrap for now.

Continue to our [Applications](APPLICATIONS.md) page to learn how to deploy application pipelines to your new cluster.
