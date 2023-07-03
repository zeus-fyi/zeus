package web3signer_cookbooks

import (
	"context"

	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
)

func (t *Web3SignerCookbookTestSuite) TestClusterAPIDeploy() {
	t.TestUploadWeb3SignerAPIChart()
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DeployCluster(ctx, Web3SignerExternalAPIClusterDefinition)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *Web3SignerCookbookTestSuite) TestClusterAPIDestroy() {
	ctx := context.Background()
	knsReq := DeployWeb3SignerExternalAPIKnsReq
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *Web3SignerCookbookTestSuite) TestCreateAPIClusterBase() {
	ctx := context.Background()
	basesInsert := []string{Web3SignerExternalAPIClusterBaseName, Web3SignerExternalAPIClusterIngressBaseName}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   Web3SignerExternalAPIClusterClassName,
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *Web3SignerCookbookTestSuite) TestCreateAPIClusterSkeletonBases() {
	ctx := context.Background()

	cc := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  Web3SignerExternalAPIClusterClassName,
		ComponentBaseName: Web3SignerExternalAPIClusterBaseName,
		SkeletonBaseNames: []string{Web3SignerExternalAPIClusterSkeletonBaseName},
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

	// ingress
	cc = zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  Web3SignerExternalAPIClusterClassName,
		ComponentBaseName: Web3SignerExternalAPIClusterIngressBaseName,
		SkeletonBaseNames: []string{Web3SignerExternalAPIClusterIngressSkeletonBaseName},
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

// TODO split this into ingress, and core workfload
func (t *Web3SignerCookbookTestSuite) TestUploadWeb3SignerAPIChart() {
	ctx := context.Background()

	ingr := topology_workloads.NewTopologyBaseInfraWorkload()
	err := Web3SignerIngressChartPath.WalkAndApplyFuncToFileType(".yaml", ingr.DecodeK8sWorkload)
	t.Require().Nil(err)

	// sets the custom ingress config
	infCfg := zeus_topology_config_drivers.IngressDriver{NginxAuthURL: t.AuthURL}
	customIngTc := zeus_topology_config_drivers.TopologyConfigDriver{
		IngressDriver: &infCfg,
	}
	customIngTc.SetCustomConfig(&ingr)
	// prints the new customized workload, and then uploads this new customized version
	// switches the print path, since it defaults to the DirOut path
	// TODO use function to do this custom config path switching and refactor
	tmp := Web3SignerIngressChartPath.DirOut
	Web3SignerIngressChartPath.DirOut = "./ethereum/web3signers/infra/custom_ingress"
	err = ingr.PrintWorkload(Web3SignerIngressChartPath)
	t.Require().Nil(err)
	Web3SignerIngressChartPath.DirOut = tmp

	// read the new config ingress
	Web3SignerIngressChartPath.DirIn = "./ethereum/web3signers/infra/custom_ingress"
	err = Web3SignerIngressChartPath.WalkAndApplyFuncToFileType(".yaml", ingr.DecodeK8sWorkload)
	t.Require().Nil(err)

	resp, err := t.ZeusTestClient.UploadChart(ctx, Web3SignerIngressChartPath, Web3SignerIngressChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployWeb3SignerExternalAPIKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployWeb3SignerExternalAPIKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(Web3SignerIngressChartPath)
	t.Require().Nil(err)

	// part2, core workload \\

	// gets base chart Web3SignerChartPath, then processes it and dumps to the custom config folder
	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err = Web3SignerChartPath.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)

	stsCfg := GetWeb3SignerAPIStatefulSetConfig(web3signerDockerImage)
	svcCfg := GetWeb3SignerAPIServiceConfig()
	tc := zeus_topology_config_drivers.TopologyConfigDriver{
		StatefulSetDriver: &stsCfg,
		ServiceDriver:     &svcCfg,
	}

	// prints the new customized workload, and then uploads this new customized version
	// switches the print path, since it defaults to the DirOut path
	tmp = Web3SignerAPIChartChartPath.DirOut
	Web3SignerAPIChartChartPath.DirOut = Web3SignerAPIChartChartPath.DirIn
	tc.SetCustomConfig(&inf)
	err = inf.PrintWorkload(Web3SignerAPIChartChartPath)
	t.Require().Nil(err)
	Web3SignerAPIChartChartPath.DirOut = tmp

	resp, err = t.ZeusTestClient.UploadChart(ctx, Web3SignerAPIChartChartPath, Web3SignerAPIChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployWeb3SignerExternalAPIKnsReq.TopologyID = resp.TopologyID
	tar = zeus_req_types.TopologyRequest{TopologyID: DeployWeb3SignerExternalAPIKnsReq.TopologyID}
	chartResp, err = t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(Web3SignerAPIChartChartPath)
	t.Require().Nil(err)
}
