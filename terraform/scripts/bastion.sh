#!/bin/bash
yum -y update
yum -y install centos-release-openshift-origin37
yum -y install origin-clients
reboot