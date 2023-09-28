package sui_cookbooks

import (
	aws_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/aws"
	do_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/do"
	gcp_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/gcp"
)

func ConfigureCloudProviderStorageClass(cp string) string {
	switch cp {
	case "aws":
		return aws_nvme.AwsStorageClass
	case "gcp":
		return gcp_nvme.GcpStorageClass
	case "do":
		return do_nvme.DoStorageClass
	case "ovh":
		return ""
	// TODO when nvme is available in public cloud (OvhCloud US)
	default:
		return ""
	}
}
