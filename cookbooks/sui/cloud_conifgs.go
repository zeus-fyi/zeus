package sui_cookbooks

// TODO: use this to override local-pv per cloud provider

func ConfigureCloudProvider(cp string) {
	switch cp {
	case "aws":
		// TODO set storage class override
		//aws_nvme.AwsStorageClass
	case "gcp":
		// gcp_nvme.GcpStorageClass
	case "do":
		// do_nvme.DoStorageClass
	case "ovh":
		// TODO when nvme is available in public cloud (OvhCloud US)
	}
}
