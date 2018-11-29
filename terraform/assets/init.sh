#!/bin/bash
yum -y update
yum -y install centos-release-openshift-origin39 epel-release firewalld
yum -y install NetworkManager openshift-clients nfs-utils python36 python36-tools nano python-passlib python2-passlib java-1.8.0-openjdk-headless docker
systemctl enable NetworkManager
systemctl start NetworkManager
systemctl enable docker
systemctl start docker
reboot