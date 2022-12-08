package zeus_endpoints

// InfraCreateV1Path uploads and saves a kubernetes app workload
const InfraCreateV1Path = "/v1/infra/create"

// InfraCreateClassV1Path creates a class definition
const InfraCreateClassV1Path = "/v1/infra/class/create"

// InfraAddBasesToClassV1Path adds base relationship to class
const InfraAddBasesToClassV1Path = "/v1/infra/class/bases/create"

// InfraAddSkeletonBasesToBaseClassV1Path adds skeleton base relationship to base class
const InfraAddSkeletonBasesToBaseClassV1Path = "/v1/infra/class/skeleton/bases/create"

// InfraReadChartV1Path reads the chart workload you uploaded
const InfraReadChartV1Path = "/v1/infra/read/chart"

// InfraReadTopologyV1Path reads the metadata for your uploaded topologies
const InfraReadTopologyV1Path = "/v1/infra/read/topologies"

// deploy infra, distributed systems api endpoints

// DeployStatusV1Path gets the topology deployment status updates
const DeployStatusV1Path = "/v1/deploy/status"

// DeployTopologyV1Path deploys topology
const DeployTopologyV1Path = "/v1/deploy"

// DeployClusterTopologyV1Path deploys a cluster topology
const DeployClusterTopologyV1Path = "/v1/deploy/cluster"

// ReplaceTopologyV1Path replaces topology at specified location with a temporary override that's local to that location
const ReplaceTopologyV1Path = "/v1/deploy/replace"

// DestroyDeployInfraV1Path destroys topology, in other words uninstalls the app
const DestroyDeployInfraV1Path = "/v1/deploy/destroy"

// live kubernetes actions requests

// ReadWorkloadV1Path reads all the statefulsets, services, ingresses, deployments, configmaps, and pods in a namespace.
const ReadWorkloadV1Path = "/v1/workload/read"

const PodsActionV1Path = "/v1/pods"
