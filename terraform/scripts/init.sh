#!/bin/bash
yum -y update
yum -y install NetworkManager
systemctl enable NetworkManager
systemctl start NetworkManager
reboot