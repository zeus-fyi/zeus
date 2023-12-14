---
sidebar_position: 7
displayed_sidebar: zK8s
---

# Sui

## Step 1 — Understanding the App Hardware Requirements

Suggested minimum hardware to run a Sui Full node on mainnet:

    CPUs: 8 physical cores / 16 vCPUs
    RAM: 128 GB
    Storage (SSD): 3.7TB+ NVMe drive

## Step 2— App Configuration

We keep a reference folder in cookbooks/sui/node/sui_config

    archival-fallback.yaml
    fullnode.yaml
    p2p-mainnet
    p2p-testnet
    Generating node Kubernetes configs

…for every combination of

- Network (devnet, testnet, mainnet)
- Cloud + Local NVMe on K8s (AWS, GCP, DigitalOcean)
- Archival Fallback
- Snapshot Integration
- Authenticated API setup for https RPC urls.
- With sidecar Hercules
- With service monitor for observability with Prometheus
- Which generates 288 distinct Kubernetes configs across every option possible

These templated configs then are added or dynamically overwritten to build configs for every network type & many configuration variations for full nodes.

```go
type SuiConfigOpts struct {
 DownloadSnapshot bool   `json:"downloadSnapshot"`
 Network          string `json:"network"`

 CloudProvider string `json:"cloudProvider"`
 WithLocalNvme bool   `json:"withLocalNvme"`

 WithIngress          bool `json:"withIngress"`
 WithServiceMonitor   bool `json:"withServiceMonitor"`
 WithArchivalFallback bool `json:"withArchivalFallback"`
 WithHercules         bool `json:"withHercules"`
}
```

### Which overrides config values like this

```go
 dataDir := "/data"
 switch cfg.CloudProvider {
 case "aws":
    dataDir = aws_nvme.AwsNvmePath
 case "gcp":
    dataDir = gcp_nvme.GcpNvmePath
 case "do":
    dataDir = do_nvme.DoNvmePath
 }
 if !cfg.WithLocalNvme {
    dataDir = "/data"
 }
 var storageClassName *string
 if cfg.WithLocalNvme {
    storageClassName = aws.String(zeus_nvme.ConfigureCloudProviderStorageClass(cfg.CloudProvider))
 }
 if !cfg.WithArchivalFallback {
    entryPointScript = "noFallBackEntrypoint.sh"
 }
 var envAddOns []v1Core.EnvVar
 if cfg.WithArchivalFallback {
    s3AccessKey := zeus_topology_config_drivers.MakeSecretEnvVar("AWS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID", "aws-credentials")
    s3SecretKey := zeus_topology_config_drivers.MakeSecretEnvVar("AWS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "aws-credentials")
    envAddOns = []v1Core.EnvVar{s3AccessKey, s3SecretKey}
 }
```

### Why generate configs this way?

It’s by far the most efficient way to generate matrix type application configurations, eg ones with many different config variations, and the only real way to actually unit test Kubernetes configs are correct using strongly typed Go code.

Meaning the odds of mistakes goes down to near zero.
Meaning you can support 10x the options with 1/10th the work.

```go
func (t *SuiCookbookTestSuite) TestSuiTestnetCfg() {
 cfg := SuiConfigOpts{
  WithLocalNvme:    true,
  DownloadSnapshot: false,
  WithIngress:      false,
  CloudProvider:    "do",
  Network:          testnet,
 }
 t.generateFromConfigDriverBuilder(cfg)

 p := SuiMasterChartPath
 p.DirIn = "./sui/node/custom_sui"
 inf := topology_workloads.NewTopologyBaseInfraWorkload()
 err := p.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
 t.Require().Nil(err)
 t.Nil(inf.Ingress)
 t.Nil(inf.Deployment)
 t.NotNil(inf.StatefulSet)
 t.NotNil(inf.Service)
 t.NotNil(inf.ConfigMap)

 t.NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates)
 t.Len(inf.StatefulSet.Spec.VolumeClaimTemplates, 1)
 t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests)
 t.Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage(), resource.MustParse(testnetDiskSize))
 t.Require().NotNil(inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName)
 t.Require().Equal(*inf.StatefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName, do_nvme.DoStorageClass)
}
```

## Kubernetes Core Node Components

### StatefulSet

#### Init Snapshots

We have the startup command to download the Sui snapshot if desired, and gets the correct genesis blob files downloaded as well for devnet, testnet, and mainnet. You’ll have to adjust the below command if you want to download a specific epoch.

```go
cmd := exec.Command(
  "aws",
  "s3",
  "cp",
  s3,
  dbPathExt,
  "--recursive",
  "--no-sign-request",
 )
 log.Info().Msgf("SuiDownloadSnapshotS3: downloading snapshot using aws cli cmd:", cmd.String())

 cmd.Stdout = &out
 cmd.Stderr = &stderr
 err := cmd.Run()
 if err != nil {
  log.Warn().Err(err).Str("stdout", out.String()).Str("stderr", stderr.String()).Msg("error downloading snapshot from S3")
  log.Err(err).Str("stdout", out.String()).Str("stderr", stderr.String()).Msg("error downloading snapshot from S3")
  return fmt.Errorf("error downloading snapshot from S3: %v", err)
 }
```

We then have this stage which downloads the necessary information if missing before starting the node.

```go
- name: init-snapshots
  image: "zeusfyi/snapshots:latest"
  imagePullPolicy: Always
  command: [ "/bin/sh" ]
  args: [ "-c","/scripts/download.sh" ]
  env:
    - name: AWS_DEFAULT_REGION
      value: "us-east-1"
  resources:
    { }
  volumeMounts:
    - name: cm-sui
      mountPath: "/scripts"
    - name: sui-client-storage
      mountPath: "/data"
```


### Resources

If you want to change these parameters, we’d advise you to add some margin so that it doesn’t run into a scheduling issue.
eg. A server may say 128GB, but it may need to run other background software and only 126GB is available for instance, so setting a request for 128GB will prevent your app from ever using that machine.

```yaml
resources:
  limits:
    cpu: "15"
    memory: "110Gi"
  requests:
    cpu: "15"
    memory: "110Gi"
```


### Disk Selection

Disk Selection
We use zK8s to let us easily switch out the storage class to the desired one. Otherwise you’ll need to adjust the below, which is setup to use the cloud provider standard block storage option.

```yaml
  volumeClaimTemplates:
    - metadata:
        name: sui-client-storage
      spec:
        accessModes:
          - ReadWriteOnce
        storageClassName: fast-disks
        resources:
          requests:
            storage: "4Ti"
```

### Config Map Components

```yaml
  download.sh: |-
    #!/bin/sh
    exec snapshots --downloadURL="" --protocol="sui" --network="mainnet" --workload-type="full" --dataDir="/data"
```

### Copying secrets into in-memory file system

```yaml
entrypoint.sh: |-
    #!/bin/sh
    cp /scripts/fullnode.yaml /secrets/fullnode.yaml
    sed -i "s/<AWS_ACCESS_KEY_ID>/$AWS_ACCESS_KEY_ID/g" /secrets/fullnode.yaml
    sed -i "s/<AWS_SECRET_ACCESS_KEY>/$AWS_SECRET_ACCESS_KEY/g" /secrets/fullnode.yaml
    exec sui-node --config-path /secrets/fullnode.yaml
```

### Node Startup Config

We also use zK8s to build the final node configs, since it adds on the relevant p2p table by network parameter, updates the db paths if needed, and appends the archival fallback option on.

```yaml
  fullnode.yaml: |-
    # Update this value to the location you want Sui to store its database
    db-path: "/data"

    network-address: "/dns/0.0.0.0/tcp/8080/http"
    metrics-address: "0.0.0.0:9184"
    # this address is also used for web socket connections
    json-rpc-address: "0.0.0.0:9000"
    enable-event-processing: true

    genesis:
      # Update this to the location of where the genesis file is stored
      genesis-file-location: "/data/genesis.blob"
    authority-store-pruning-config:
      # Number of epoch dbs to keep 
      # Not relevant for object pruning
      num-latest-epoch-dbs-to-retain: 3
      # The amount of time, in seconds, between running the object pruning task.
      # Not relevant for object pruning
      epoch-db-pruning-period-secs: 3600
      # Number of epochs to wait before performing object pruning.
      num-epochs-to-retain: 0
      # Advanced setting: Maximum number of checkpoints to prune in a batch. The default
      # settings are appropriate for most use cases.
      max-checkpoints-in-batch: 10
      # Advanced setting: Maximum number of transactions in one batch of pruning run. The default
      # settings are appropriate for most use cases.
      max-transactions-in-batch: 1000
      # Not documented anywhere but necessary for pruning to work
      use-range-deletion: true
      # Used for tx pruning
      # Number of epochs to wait before performing transaction pruning.
      # When this is N (where N >= 2), Sui prunes transactions and effects from 
      # checkpoints up to the `current - N` epoch. Sui never prunes transactions and effects from the current and
      # immediately prior epoch. N = 2 is a recommended setting for Sui Validator nodes and Sui Full nodes that don't 
      # serve RPC requests.
      #num-epochs-to-retain-for-checkpoints: 2
      # Ensures that individual database files periodically go through the compaction process.
      # This helps reclaim disk space and avoid fragmentation issues
      #periodic-compaction-threshold-days: 1
```

### zK8s Cluster Definition

```go
var (
 suiNodeDefinition = zeus_cluster_config_drivers.ClusterDefinition{
    ClusterClassName: Sui,
    ComponentBases:   suiComponentBases,
 }
 suiComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
    Sui: suiMasterComponentBase,
 }
 suiMasterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
  SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
     Sui: suiSkeletonBaseConfig,
  },
 }
 suiSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
   SkeletonBaseNameChartPath: SuiMasterChartPath,
 }
 SuiMasterChartPath = filepaths.Path{
    DirIn:       "./sui/node/infra",
    DirOut:      "./sui/output",
    FnIn:        Sui, // filename for your gzip workload
 }
```

### Config Overrides

Using the zK8s style of coding Kubernetes apps lets us create a wide variation in application configuration choices from a relatively simple place. You can easily manage a single page that distills and condenses all the most changes into a handful of top-level parameters, which we created a Sui specific one for toggling add-ons to the base config using SuiConfigOpts.

How does this compare to your current workflows? Let us know!

```go
const (
    // networks
    mainnet = "mainnet"
    testnet = "testnet"
    devnet  = "devnet"

    // docker image references
    dockerImage        = "mysten/sui-node:mainnet"
    dockerImageTestnet = "mysten/sui-node:testnet"
    dockerImageDevnet  = "mysten/sui-node:devnet"

    hercules = "hercules"

    // mainnet workload compute resources
    cpuCores   = "16"
    memorySize = "128Gi"
    // mainnet workload disk sizes
    mainnetDiskSize = "4Ti"

    // testnet compute resources
    cpuCoresTestnet   = "7500m"
    memorySizeTestnet = "63Gi"
    // testnet workload disk sizes
    testnetDiskSize = "3Ti"
    devnetDiskSize  = "2Ti"

    // workload label, name, or k8s references
    suiDiskName  = "sui-client-storage"
    suiConfigMap = "cm-sui"

    // workload type, impacts where/if for snapshot downloads
    suiNodeConfig      = "full"
    suiValidatorConfig = "validator"
    suiNoSnapshot      = "min"

    SuiRpcPortName = "http-rpc"

    DownloadMainnet = "downloadMainnetNode"
)

type SuiConfigOpts struct {
    DownloadSnapshot bool   `json:"downloadSnapshot"`
    Network          string `json:"network"`

    CloudProvider string `json:"cloudProvider"`
    WithLocalNvme bool   `json:"withLocalNvme"`

    WithIngress          bool `json:"withIngress"`
    WithServiceMonitor   bool `json:"withServiceMonitor"`
    WithArchivalFallback bool `json:"withArchivalFallback"`
    WithHercules         bool `json:"withHercules"`
}

func GetSuiClientNetworkConfigBase(cfg SuiConfigOpts) zeus_cluster_config_drivers.ComponentBaseDefinition {
    diskSize := mainnetDiskSize
    cpuSize := cpuCores
    memSize := memorySize
    dockerImageSui := dockerImage
    entryPointScript := "entrypoint.sh"
    switch cfg.Network {
    case mainnet:
       cpuSize = cpuCores
       memSize = memorySize
       diskSize = mainnetDiskSize
       dockerImageSui = dockerImage
    case testnet:
       diskSize = testnetDiskSize
       cpuSize = cpuCoresTestnet
       memSize = memorySizeTestnet
       dockerImageSui = dockerImageTestnet
    case devnet:
       diskSize = devnetDiskSize
       cpuSize = cpuCoresTestnet
       memSize = memorySizeTestnet
       dockerImageSui = dockerImageDevnet
       entryPointScript = "noFallBackEntrypoint.sh"
    }

    sd := &config_overrides.ServiceDriver{}
    if cfg.WithIngress {
       sd.AddNginxTargetPort("nginx", SuiRpcPortName)
    }

    dataDir := "/data"
    switch cfg.CloudProvider {
    case "aws":
       dataDir = aws_nvme.AwsNvmePath
    case "gcp":
       dataDir = gcp_nvme.GcpNvmePath
    case "do":
       dataDir = do_nvme.DoNvmePath
    }
    if !cfg.WithLocalNvme {
       dataDir = "/data"
    }
    var storageClassName *string
    if cfg.WithLocalNvme {
       storageClassName = aws.String(zeus_nvme.ConfigureCloudProviderStorageClass(cfg.CloudProvider))
    }
    if !cfg.WithArchivalFallback {
       entryPointScript = "noFallBackEntrypoint.sh"
    }
    var envAddOns []v1Core.EnvVar
    if cfg.WithArchivalFallback {
       s3AccessKey := config_overrides.MakeSecretEnvVar("AWS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID", "aws-credentials")
       s3SecretKey := config_overrides.MakeSecretEnvVar("AWS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "aws-credentials")
       envAddOns = []v1Core.EnvVar{s3AccessKey, s3SecretKey}
    }

    wkType := "full"
    if !cfg.DownloadSnapshot {
       wkType = "min"
    }
    downloadCmd := fmt.Sprintf("#!/bin/sh\nexec snapshots --downloadURL=\"\" --protocol=\"sui\" --network=\"%s\" --workload-type=\"%s\" --dataDir=\"%s\"", cfg.Network, wkType, dataDir)
    sbCfg := zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
       SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
       SkeletonBaseNameChartPath: SuiMasterChartPath,
       TopologyConfigDriver: &config_overrides.TopologyConfigDriver{
          ConfigMapDriver: &config_overrides.ConfigMapDriver{
             ConfigMap: v1Core.ConfigMap{
                ObjectMeta: metav1.ObjectMeta{Name: suiConfigMap},
                Data: map[string]string{
                   "fullnode.yaml": OverrideNodeConfigDataDir(dataDir, cfg),
                   "download.sh":   downloadCmd,
                },
             },
          },
          ServiceDriver: sd,
          StatefulSetDriver: &config_overrides.StatefulSetDriver{
             ContainerDrivers: map[string]config_overrides.ContainerDriver{
                Sui: {
                   Container: v1Core.Container{
                      Name:      Sui,
                      Image:     dockerImageSui,
                      Resources: config_overrides.CreateComputeResourceRequirementsLimit(cpuSize, memSize),
                      VolumeMounts: []v1Core.VolumeMount{{
                         Name:      suiDiskName,
                         MountPath: dataDir,
                      }},
                      Command: []string{fmt.Sprintf("/scripts/%s", entryPointScript)},
                   },
                   AppendEnvVars: envAddOns,
                },
                "hercules": {
                   Container: v1Core.Container{
                      Name: "hercules",
                      VolumeMounts: []v1Core.VolumeMount{{
                         Name:      suiDiskName,
                         MountPath: dataDir,
                      }},
                   },
                   IsDeleteContainer: !cfg.WithHercules,
                },
                "init-snapshots": {
                   Container: v1Core.Container{
                      Name: "init-snapshots",
                      VolumeMounts: []v1Core.VolumeMount{{
                         Name:      suiDiskName,
                         MountPath: dataDir,
                      }},
                   },
                   IsInitContainer: true,
                },
                "init-chown-data": {
                   Container: v1Core.Container{
                      Name:    "init-chown-data",
                      Command: []string{"chown", "-R", "10001:10001", dataDir},
                      VolumeMounts: []v1Core.VolumeMount{{
                         Name:      suiDiskName,
                         MountPath: dataDir,
                      }},
                   },
                   IsInitContainer: true,
                },
             },
             PVCDriver: &config_overrides.PersistentVolumeClaimsConfigDriver{
                PersistentVolumeClaimDrivers: map[string]v1Core.PersistentVolumeClaim{
                   suiDiskName: {
                      ObjectMeta: metav1.ObjectMeta{Name: suiDiskName},
                      Spec: v1Core.PersistentVolumeClaimSpec{
                         Resources:        config_overrides.CreateDiskResourceRequirementsLimit(diskSize),
                         StorageClassName: storageClassName,
                      },
                   },
                }},
          },
       }}
    suiCompBase := zeus_cluster_config_drivers.ComponentBaseDefinition{
       SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
          Sui: sbCfg,
       },
    }
    return suiCompBase
}
```