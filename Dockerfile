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
        && rm terraform.zip

# Create Directories
RUN mkdir -p /app && mkdir -p /root/.ssh && mkdir -p /root/.aws && mkdir -p /app/generated

WORKDIR /app
ADD orchestration/openshift-aws /app/openshift-aws
ADD playbooks terraform templates openshift-ansible config.default.yaml /app/

ENTRYPOINT [ "/app/openshift-aws" ]
