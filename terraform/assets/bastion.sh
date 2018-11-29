#!/bin/bash
yum -y update
yum -y install centos-release-openshift-origin39 epel-release
yum -y install origin-clients nfs-utils python36 python36-tools nano python-passlib python2-passlib
reboot