#!/bin/bash
yum -y update
yum -y install centos-release-openshift-origin37 epel-release
yum -y install NetworkManager nfs-utils python36 python36-tools nano python-passlib python2-passlib
systemctl enable NetworkManager
systemctl start NetworkManager
reboot