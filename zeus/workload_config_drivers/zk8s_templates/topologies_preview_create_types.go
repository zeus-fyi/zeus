package zk8s_templates

type ConfigMap map[string]string

type Ingress struct {
	AuthServerURL string `json:"authServerURL"`
	Host          string `json:"host"`
}

type IngressPath struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
}

type IngressPaths map[string]IngressPath

type Deployment struct {
	ReplicaCount int `json:"replicaCount"`
}

type StatefulSet struct {
	ReplicaCount int           `json:"replicaCount"`
	PVCTemplates []PVCTemplate `json:"pvcTemplates"`
}

type PVCTemplate struct {
	Name               string `json:"name"`
	AccessMode         string `json:"accessMode"`
	StorageSizeRequest string `json:"storageSizeRequest"`
}

type ResourceRequirements struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type Port struct {
	Name               string        `json:"name"`
	Number             string        `json:"number"`
	Protocol           string        `json:"protocol"`
	IngressEnabledPort bool          `json:"ingressEnabledPort"`
	ProbeSettings      ProbeSettings `json:"probeSettings,omitempty"`
}

type VolumeMount struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
}

type ProbeSettings struct {
	UseForLivenessProbe  bool `json:"useForLivenessProbe,omitempty"`
	UseForReadinessProbe bool `json:"useForReadinessProbe,omitempty"`
	UseTcpSocket         bool `json:"useTcpSocket,omitempty"`
}
