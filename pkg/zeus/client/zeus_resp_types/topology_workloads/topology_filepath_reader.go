package topology_workloads

import yaml_fileio "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/yaml"

func (t *TopologyBaseInfraWorkload) DecodeK8sWorkload(filepath string) error {
	b, err := yaml_fileio.ReadYamlConfig(filepath)
	if err != nil {
		return err
	}
	err = t.DecodeBytes(b)
	return err
}
