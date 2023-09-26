## GKE NVMe Setup

[https://cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/local-ssd](https://cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/local-ssd)

Note: Local SSDs require machine type n1-standard-1 or larger; the default machine type, e2-medium is not supported. You
can learn more about machine types in the Compute Engine documentation.

To create a node pool with local NVMe SSDs for raw block access, run the following command:

```shell
gcloud container node-pools create POOL_NAME \
    --cluster CLUSTER_NAME \
    --local-nvme-ssd-block count=NUMBER_OF_DISKS
```

### POOL_NAME

the name of your new node pool.

### CLUSTER_NAME

the name of the cluster.

### NUMBER_OF_DISKS

the number of local SSD disks to provision on each node. The maximum number of disks varies by machine type and region.
For C3 machine types with local SSD (Preview), you must use the number of local SSDs that correspond to the machine
type. For more information, see Supported disk types for C3.

## Local PersistentVolumes

Local SSDs can be specified as PersistentVolumes.

You can create PersistentVolumes from local SSDs by manually creating a PersistentVolume, or by running the local volume
static provisioner.

Local PersistentVolume objects are not automatically cleaned up when a node is deleted, upgraded, repaired, or scaled
down. We recommend you periodically scan and delete stale Local PersistentVolume objects associated with deleted nodes.