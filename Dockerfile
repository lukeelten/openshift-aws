FROM centos:7

USER root
WORKDIR /root

# Install OpenShift client (oc)
RUN yum -y install centos-release-openshift-origin37 && yum -y install origin-clients && rm -rf /var/cache/yum

# Install ansible tool
RUN yum -y install epel-release && yum -y install ansible unzip python-passlib python2-passlib && rm -rf /var/cache/yum

# Install Java JRE
RUN yum -y install java-1.8.0-openjdk-headless && rm -rf /var/cache/yum

# Install terraform
RUN curl https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip -o terraform.zip \
        && unzip terraform.zip \
        && mv terraform /usr/bin \
        && chmod 755 /usr/bin/terraform \
        && rm terraform.zip

# Create Directories
RUN mkdir -p /app/generated && mkdir -p /root/.aws

ADD orchestration/openshift-aws /usr/bin/openshift-aws
ADD . /app/

WORKDIR /app/terraform
RUN terraform init

WORKDIR /app