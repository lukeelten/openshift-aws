# Automated OpenShift Installation for AWS




## Prerequisites
* Docker Container Engine (__Version: >= 17.09__)

### Supported OS
* Linux
* MacOS

It should work properly with Docker for Windows, nevertheless it is not supported or tested.


# Quick Start



# Configuration



# Available Scripts

There are three bash scripts available which make life easier.


## Building the application
All parts of the application can be build using a provided script.
The following dependencies are required:
* Docker
* Git

There are no more dependencies, compilation will take place inside a docker image.
```bash
./build.sh
```

## Creating a cluster




```bash
./run.sh -name="Test Cluster"
```

### Required Parameters


### AWS Keys
| Parameter   | Description |
|-------------|----------------------------------------------------------------------|
| -aws-key    | AWS access key id. If empty the credentials used for AWS CLI will be loaded |
| -aws-secret | AWS secret key. If empty the credentials used for AWS CLI will be loaded |



### Optional Parameters

| Parameter   | Description |
|-------------|----------------------------------------------------------------------|
| -skip-terraform | Skip terraform invocation. This should be used if there is an already existing infrastructure.                                                                                                                     |
| -skip-config    | Skip generation of configuration files. This is useful when there are already configuration files and only the installer should be run. Nevertheless recreating the configuration files would not harm the system. |
| -skip-pre       | Skip prerequisites playbooks of OpenShift installer. This is useful when restarting the installer due to a previous failure.                                                                                       |
| -verbose        | Enables Ansible -vvv verbose mode                                                                                                                                                                                  |                                                                                                                                                                                                                 |


## Destroying a cluster

If you want to delete a previously created cluster, simply run
```bash
./destroy.sh
```

It is important that the files generated during creation are still available.
At least the Terraform state (terraform.tfstate) and the Terraform configuration (configuration.tfvars) must be available.

__Attention:__ After a successful destruction the terraform state and the old SSH key will be removed automatically.
