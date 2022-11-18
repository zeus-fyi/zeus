package zeus_common_types

type CloudCtxNs struct {
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	Context       string `json:"context"`
	Namespace     string `json:"namespace"`
	Env           string `json:"env"`
}
