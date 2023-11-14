package zk8s_templates

type Containers map[string]Container

type Container struct {
	IsInitContainer bool        `json:"isInitContainer"`
	ImagePullPolicy string      `json:"imagePullPolicy,omitempty"`
	DockerImage     DockerImage `json:"dockerImage"`
}

type DockerImage struct {
	ImageName            string               `json:"imageName"`
	Cmd                  string               `json:"cmd"`
	Args                 string               `json:"args"`
	ResourceRequirements ResourceRequirements `json:"resourceRequirements,omitempty"`
	Ports                []Port               `json:"ports,omitempty"`
	VolumeMounts         []VolumeMount        `json:"volumeMounts,omitempty"`
}
