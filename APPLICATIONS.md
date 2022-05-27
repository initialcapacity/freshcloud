# Fresh Cloud

## Application clusters

Create a service account for you applications. _Note_, for the moment, the below requires application clusters
to reside in the same google cloud project as the management cluster.

```base
freshctl clusters gcp create-service-account
```

Similar to your managment cluster, create a `.env_app.sh` file for you application.

* `REGISTRY_*` variables should match the management cluster above.
* `GCP_SERVICE_ACCOUNT_JSON` should be located in your lcoal `.freshcloud` directory.
* `APP_CONFIGURATION_PATH` should contain kubernetes objects.
* `APP_PIPELINE_*_PATH` should contain concourse pipeline files.

See our [Kotlin-ktor-starter](https://github.com/initialcapacity/kotlin-ktor-starter/tree/main/deployments) for an
example.

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
freshctl clusters gcp create
freshctl services contour
```

Create a DNS entry for your load balancer similar to your management cluster.

```bash
freshctl services cert-manager
freshctl pipelines push-build-image
freshctl pipelines deploy
```

Cleaning up

```bash
docker rm -f $(docker ps -a -q)
docker volume rm -f $(docker volume ls -q)
docker system prune -a -f
```
