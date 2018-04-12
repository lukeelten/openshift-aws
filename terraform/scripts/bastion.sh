#!/bin/bash
yum -y update
yum -y install centos-release-openshift-origin37 epel-release
yum -y install origin-clients nfs-utils python36 python36-tools nano
reboot