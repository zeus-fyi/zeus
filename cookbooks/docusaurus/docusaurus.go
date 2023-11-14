package docusaurus_cookbooks

import (
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var (
	DocusaurusClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "docusaurus",
		CloudCtxNs:       DocusaurusCloudCtxNs,
		ComponentBases:   DocusaurusComponentBases,
	}
	DocusaurusCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "nyc1",
		Context:       "do-nyc1-do-nyc1-zeus-demo",
		Namespace:     "docusaurus",
		Env:           "production",
	}
	DocusaurusComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"docusaurus": DocusaurusComponentBase,
	}
	DocusaurusComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"docusaurus": Docusaurus,
		},
	}
	Docusaurus = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: DocusaurusChartPath,
	}
)
