FROM golang:1 AS build-env

COPY orchestration /app
WORKDIR /app


RUN go get -t -d ./...
RUN go test ./...
RUN go build ./cmd/openshift-aws


# -----------------------------------------------------------------------------------


FROM centos:7

USER root
WORKDIR /root

# Install OpenShift client (oc)
RUN yum -y install centos-release-openshift-origin39 && yum -y install origin-clients && rm -rf /var/cache/yum

# Install ansible tool
RUN yum -y install epel-release && yum -y install ansible unzip python-passlib python2-passlib && rm -rf /var/cache/yum

# Install Java JRE
RUN yum -y install java-1.8.0-openjdk-headless && rm -rf /var/cache/yum

# Install terraform
RUN curl https://releases.hashicorp.com/terraform/0.11.10/terraform_0.11.10_linux_amd64.zip -o terraform.zip \
        && unzip terraform.zip \
        && mv terraform /usr/bin \
        && chmod 755 /usr/bin/terraform \
        && rm terraform.zip

# Create Directories
RUN mkdir -p /app/generated && mkdir -p /root/.aws

COPY dockerentry.sh /usr/bin/
COPY --from=build-env /app/openshift-aws /usr/bin/

COPY openshift-ansible /app/openshift-ansible

COPY terraform /app/terraform

WORKDIR /app/terraform
RUN terraform init && chmod 755 /usr/bin/openshift-aws /usr/bin/dockerentry.sh

WORKDIR /app
COPY templates /app/templates
COPY playbooks /app/playbooks
COPY config.yaml /app/config.yaml
