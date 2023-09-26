package aws_nvme

const (
	AwsStorageClass = "fast-disks"
	AwsNvmePath     = "/mnt/fast-disks"
)

func AddAwsEksNvmeLabels(labels map[string]string) map[string]string {
	labels["fast-disk-node"] = "pv-raid"
	return labels
}
