package sui_cookbooks

// TODO: use this to override local-pv per cloud provider

func ConfigureCloudProvider(cp string) {
	switch cp {
	case "aws":
		// TODO set storage class override
	case "gcp":
	case "do":
	case "ovh":
		// TODO when nvme is available in public cloud (OvhCloud US)
	}
}
