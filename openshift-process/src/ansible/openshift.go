package ansible

import (
	"aws"
	"settings"
)

const FSID = "fs-38827261"
const EFS_KEY_ID = "AKIAITWNTFTUS4RYRIXQ"
const EFS_KEY_SECRET = "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"

func GenerateOpenshiftInventory(filename string) *Inventory {
	masters := aws.MasterNodes()
	infra := aws.InfraNodes()
	apps := aws.AppNodes()
	bastion := aws.BastionNode()

	//apiLb := aws.GetInternalLB()

	defaultSubdomain := "apps.cc-openshift.de"
	externalMasterHostname := "master.cc-openshift.de"
	internalMasterHostname := "internal-api.cc-openshift.de"

	sshConfig := settings.NewSshConfig("ssh.cfg")
	bastionConfig := settings.NewHostConfig(bastion.ExternalDns)
	bastionConfig.AddVar("Hostname", bastion.ExternalDns)
	bastionConfig.AddVar("User", "centos")
	bastionConfig.AddVar("ControlMaster", "auto")
	bastionConfig.AddVar("ControlPersist", "5m")
	bastionConfig.AddVar("ControlPath", "~/.ssh/ansible-%r@%h:%p")
	bastionConfig.AddVar("StrictHostKeyChecking", "no")
	bastionConfig.AddVar("ProxyCommand", "none")
	bastionConfig.AddVar("ForwardAgent", "yes")
	sshConfig.AddHost(bastionConfig)

	nodeConfig := settings.NewHostConfig("10.10.*")
	nodeConfig.AddVar("ProxyCommand", "ssh -o StrictHostKeyChecking=no -W %h:%p centos@" + bastion.ExternalDns)
	nodeConfig.AddVar("StrictHostKeyChecking", "no")
	sshConfig.AddHost(nodeConfig)
	sshConfig.Write()

	inventory := NewInventory(filename)
	inventory.AddSection("OSEv3:children", []string{"masters", "nodes", "etcd"})

	var vars []string
	vars = append(vars, "ansible_user=centos", "ansible_become=true", "deployment_type=origin")
	vars = append(vars, "ansible_ssh_common_args='-F ssh.cfg -o StrictHostKeyChecking=no -o ControlMaster=auto -o ControlPersist=30m'")
	vars = append(vars, "openshift_release=v3.7.1", "openshift_image_tag=v3.7.1")
	vars = append(vars, "openshift_router_selector='region=infra'", "openshift_registry_selector='region=infra'")
	vars = append(vars, "openshift_master_cluster_method=native")
	vars = append(vars, "openshift_master_default_subdomain='" + defaultSubdomain + "'")
	vars = append(vars, "openshift_clock_enable=true", "openshift_use_dnsmasq=true", "os_firewall_use_firewalld=true")

	// persistence
	/*
	vars = append(vars, "openshift_provisioners_install_provisioners=true")
	vars = append(vars, "openshift_provisioners_efs=true")
	vars = append(vars, "openshift_provisioners_efs_region='eu-central-1'")
	vars = append(vars, "#openshift_provisioners_efs_nodeselector='region=infra'")
	vars = append(vars, "openshift_provisioners_efs_aws_access_key_id='" + EFS_KEY_ID + "'")
	vars = append(vars, "openshift_provisioners_efs_aws_secret_access_key='" + EFS_KEY_SECRET + "'")
	vars = append(vars, "openshift_provisioners_efs_fsid='" + FSID + "'")
	vars = append(vars, "openshift_provisioners_efs_path=/persistentvolumes")
	vars = append(vars, "openshift_provisioners_image_version=v3.7")

	if !settings.ActiveSettings.ActivateTSB {
		vars = append(vars, "openshift_enable_service_catalog=true")
	}
	*/

	vars = append(vars, "openshift_master_cluster_hostname='" + internalMasterHostname + "'", "openshift_master_cluster_public_hostname='" + externalMasterHostname + "'")
	vars = append(vars, "openshift_disable_check=docker_storage,memory_availability,package_version")
	vars = append(vars, "openshift_master_identity_providers=[{'name': 'htpasswd_auth', 'login': 'true', 'challenge': 'true', 'kind': 'HTPasswdPasswordIdentityProvider', 'filename': '/etc/origin/master/htpasswd'}]")
	vars = append(vars, "openshift_master_htpasswd_users={'admin': '$apr1$zgSjCrLt$1KSuj66CggeWSv.D.BXOA1', 'user': '$apr1$.gw8w9i1$ln9bfTRiD6OwuNTG5LvW50'}")

	inventory.AddSection("OSEv3:vars", vars)

	nodesSection := generateNodeLines(masters, false, false)
	nodesSection = append(nodesSection, generateNodeLines(infra, true, true)...)
	nodesSection = append(nodesSection, generateNodeLines(apps, false, true)...)


	inventory.AddSection("masters", generateNodeLines(masters, false, false))
	inventory.AddSection("etcd", extractNodeIps(masters, false))
	inventory.AddSection("nodes", nodesSection)

	return inventory
}

func generateNodeLines(nodes []aws.NodeInfo, infra bool, schedulable bool) []string {
	var lines []string

	for _,node := range nodes {
		lines = append(lines, generateNodeLine(node, infra, schedulable))
	}

	return lines
}

func generateNodeLine(node aws.NodeInfo, infra bool, schedulable bool) string {
	var s string
	extra := " openshift_schedulable="

	if schedulable {
		extra += "true"
	} else {
		extra += "false"
	}

	extra += " openshift_node_labels=\"{'region':'"

	if infra {
		extra += "infra"
	} else {
		extra += "primary"
	}

	extra += "','zone':'" + node.Zone + "'}\""

	s += node.InternalIp + extra
	s += " openshift_ip=" + node.InternalIp
	s += " openshift_hostname=" + node.InternalDns

	return s
}

func extractNodeIps(nodes []aws.NodeInfo, public bool) []string {
	var ips []string

	for _,node := range nodes {
		if public {
			ips = append(ips, node.ExternalIp)
		} else {
			ips = append(ips, node.InternalIp)
		}
	}

	return ips
}