FROM centos:7

USER root
WORKDIR /root

# Install OpenShift client (oc)
RUN yum -y install centos-release-openshift-origin37 && yum -y install origin-clients && rm -rf /var/cache/yum

# Install AWS CLI tool (not needed yet)
#RUN yum -y install epel-release && yum -y install python-pip && pip -q install awscli && rm -rf /var/cache/yum

# Install ansible tool
RUN yum -y install epel-release && yum -y install ansible unzip && rm -rf /var/cache/yum

# Install terraform
RUN curl https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip -o terraform.zip \
        && unzip terraform.zip \
        && mv terraform /usr/bin \
        && rm terraform.zip

# Create Directories
RUN mkdir -p /app/openshift-process/generated

# Copy files
COPY openshift-ansible /app/openshift-ansible
COPY terraform /app/terraform
COPY openshift-process/templates /app/openshift-process/templates
COPY openshift-process/playbooks /app/openshift-process/playbooks
COPY openshift-process/openshift-process /app/openshift-process

WORKDIR /app/openshift-process

CMD ["sh"]
