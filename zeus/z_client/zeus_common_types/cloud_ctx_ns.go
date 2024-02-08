package zeus_common_types

type CloudCtxNs struct {
	ClusterCfgID int `json:"clusterCfgID,omitempty"`

	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	Context       string `json:"context"`
	Namespace     string `json:"namespace"`
	Alias         string `json:"alias,omitempty"`
	Env           string `json:"env,omitempty"`
}

func (c *CloudCtxNs) CheckIfEmpty() bool {
	if c.CloudProvider == "" {
		return true
	}
	if c.Region == "" {
		return true
	}
	if c.Context == "" {
		return true
	}
	if c.Namespace == "" {
		return true
	}
	return false
}
