package docusaurus_cookbooks

import (
	"context"
	"fmt"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/zk8s_templates"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

const (
	docusaurus         = "docusaurus"
	docusaurusTemplate = "docusaurus-template"
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

func CreateDocusaurusDeployment(ctx context.Context, zc zeus_client.ZeusClient, createClass bool) error {
	dockerImage := "docker.io/zeusfyi/docusaurus-template:latest"
	wd := zeus_cluster_config_drivers.WorkloadDefinition{
		WorkloadName: docusaurusTemplate,
		ReplicaCount: 1,
		Containers: zk8s_templates.Containers{
			docusaurusTemplate: zk8s_templates.Container{
				IsInitContainer: false,
				ImagePullPolicy: "Always",
				DockerImage: zk8s_templates.DockerImage{
					ImageName: dockerImage,
					ResourceRequirements: zk8s_templates.ResourceRequirements{
						CPU:    "100m",
						Memory: "500Mi",
					},
					Ports: []zk8s_templates.Port{
						{
							Name:               "http",
							Number:             "3000",
							Protocol:           "TCP",
							IngressEnabledPort: true,
							ProbeSettings: zk8s_templates.ProbeSettings{
								UseForLivenessProbe:  true,
								UseForReadinessProbe: true,
								UseTcpSocket:         true,
							},
						},
					},
				},
			},
		},
		FilePath: filepaths.Path{
			DirOut: "./docusaurus/outputs",
			FnIn:   docusaurusTemplate,
		},
	}
	cd, err := zeus_cluster_config_drivers.GenerateDeploymentCluster(ctx, wd)
	if err != nil {
		return err
	}
	cd.IngressPaths = map[string]zk8s_templates.IngressPath{
		wd.WorkloadName: {
			Path:     "/",
			PathType: "ImplementationSpecific",
		},
	}

	prt, err := zeus_cluster_config_drivers.PreviewTemplateGeneration(ctx, cd)
	if err != nil {
		return err
	}

	if createClass {
		gcd := zeus_cluster_config_drivers.CreateGeneratedClusterClassCreationRequest(cd)
		fmt.Println(gcd)
		gcdExp := DocusaurusClusterDefinition.BuildClusterDefinitions()
		fmt.Println(gcdExp)
		err = gcd.CreateClusterClassDefinitions(ctx, zc)
		if err != nil {
			return err
		}
	}

	_, err = prt.UploadChartsFromClusterDefinition(ctx, zc, true)
	if err != nil {
		return err
	}
	return nil
}
