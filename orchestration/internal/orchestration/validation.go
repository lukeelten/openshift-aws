package orchestration

import "path/filepath"
import "../util"

func prettyPrint(dir string) string {
	return filepath.Clean(dir) + "/"
}

func checkTerraformDir (dir string) bool {
	return util.FileExists(dir + "provider.tf") && util.FileExists(dir + "variables.tf")
}

func checkInstallerDir (dir string) bool {
	return util.FileExists(dir + "playbooks/prerequisites.yml") && util.FileExists(dir + "playbooks/deploy_cluster.yml")
}

func checkOutputDir (dir string) bool {
	return util.IsWritable(dir)
}

func checkPlaybookDir(dir string) bool {
	return util.FileExists(dir + "ebs.xml") && util.FileExists(dir + "post-config.yml")
}

func checkTemplatesDir(dir string) bool {
	return util.FileExists(dir + "ssh.tmpl") && util.FileExists(dir + "inventory.tmpl")
}