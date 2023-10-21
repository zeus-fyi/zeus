## AWS NVMe Setup

[https://aws.amazon.com/blogs/containers/eks-persistent-volumes-for-instance-store](https://aws.amazon.com/blogs/containers/eks-persistent-volumes-for-instance-store)

Note: launch templates are needed to bootstrap k8s nodes on creation
https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html

Note: you may need to request a service quota limit increase for vCPUs first.

You will to replace the cluster name, and security group id, etc on the pv-raid-nodegroup.yaml and the
pv-nvme-nodegroup.yaml

Once you've installed the local pv provisioner. Launch your node group like this:

```shell
eksctl create nodegroup --config-file=pv-raid-nodegroup.yaml
```

To remove a node group

```shell
eksctl delete nodegroup --cluster=eksworkshop-eksctl --name=eks-pv-raid-ng
```