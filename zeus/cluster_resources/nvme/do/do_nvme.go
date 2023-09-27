package do_nvme

const (
	DoStorageClass = "fast-disks"
	DoNvmePath     = "/mnt/fast-disks"
)

func AddDoNvmeLabels(labels map[string]string) map[string]string {
	labels["fast-disk-node"] = "pv-raid"
	return labels
}
