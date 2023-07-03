package zeus_common_types

type CloudCtxNs struct {
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	Context       string `json:"context"`
	Namespace     string `json:"namespace"`
	Alias         string `json:"alias,omitempty"`
	Env           string `json:"env,omitempty"`
}
