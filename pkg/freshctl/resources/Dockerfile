# Container to run in our pipeline with kubectl installed.
FROM golang:alpine

ARG KUBE_VERSION="v1.18.10"
ARG CLOUD_SDK_VERSION="384.0.1"

RUN apk add --update ca-certificates \
 && apk add --update -t deps curl \
 && apk add --update gettext \
 && apk add --update jq \
 && apk add --update git \
 && curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBE_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
 && chmod +x /usr/local/bin/kubectl \
 && apk del --purge deps \
 && rm /var/cache/apk/*

ENV PATH=/go/google-cloud-sdk/bin:$PATH

RUN apk add --no-cache curl python3  py3-pip py3-crcmod bash libc6-compat openssh-client git gnupg \
 && curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-${CLOUD_SDK_VERSION}-linux-x86_64.tar.gz \
 && tar xzf google-cloud-sdk-${CLOUD_SDK_VERSION}-linux-x86_64.tar.gz \
 && rm google-cloud-sdk-${CLOUD_SDK_VERSION}-linux-x86_64.tar.gz

RUN pip3 --no-cache-dir install --no-input --upgrade awscli