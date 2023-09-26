## AWS NVMe Setup

[https://aws.amazon.com/blogs/containers/eks-persistent-volumes-for-instance-store](https://aws.amazon.com/blogs/containers/eks-persistent-volumes-for-instance-store)

Note: launch templates are needed to bootstrap k8s nodes on creation
https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html

Note: you may need to request a service quota limit increase for vCPUs first.

You will to replace the cluster name, and security group id, etc on the pv-raid-nodegroup.yaml and the
pv-nvme-nodegroup.yaml