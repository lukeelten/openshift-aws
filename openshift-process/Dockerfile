FROM centos:latest

USER root

# Needs to be two separate install commands
RUN yum -y install centos-release-openshift-origin37 && yum -y install origin-clients && rm -rf /var/cache/yum
RUN yum -y install epel-release && yum -y install python-pip && pip -q install awscli && rm -rf /var/cache/yum
