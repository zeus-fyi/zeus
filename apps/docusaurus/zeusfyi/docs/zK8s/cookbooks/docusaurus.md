---
sidebar_position: 5
displayed_sidebar: zK8s
---

# Docusaurus #

## Medium Tutorial

[How to deploy Docusaurus on Kubernetes](https://medium.com/zeusfyi/hosted-docusaurus-in-5-minutes-and-under-10-month-af999d7ef90a)

## Relevant Code

#### ```apps/docusaurus ```

Docusaurus is a static-site generator. It builds a single-page application with fast client-side navigation,
leveraging the full power of React to make your site interactive. It provides out-of-the-box documentation features
but can be used to create any kind of site (personal website, product, blog, marketing landing pages, etc).

#### ```cookbooks/docusaurus ```

Contains full Kubernetes infra setup for Docusaurus, which is the site generator used to build our docs site. Which you
can deploy on our platform via UI, API, or any place that supports Kubernetes.

#### ```docker/docusaurus ```

You can use this docker setup as a template to build your own docusaurus site.

#### ```.github/workflows/docusaurus.yml ```

You can use this docker setup as a template to build your own docusaurus CI using GitHub actions.

## Docusaurus Function

Docusaurus Cluster Generation Function
You simply call this function, and it’ll generate a Deployment, Service, and Ingress file that also prints out a copy
of the generated workloads and then uploads them to the registered cluster class, in this case that is
docusaurus-template.
Just update the docker image to your own, and it’s literally that simple.

```go
func CreateDocusaurusDeployment(zc zeus_client.ZeusClient, nc bool) error {
wd := zeus_cluster_config_drivers.WorkloadDefinition{
WorkloadName: "docusaurus-template",
ReplicaCount: 1,
Containers: zk8s_templates.Containers{
docusaurusTemplate: zk8s_templates.Container{
ImagePullPolicy: "Always",
DockerImage: zk8s_templates.DockerImage{
ImageName: "docker.io/zeusfyi/docusaurus-template:latest",
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
if nc {
gcd := zeus_cluster_config_drivers.CreateGeneratedClusterClassCreationRequest(cd)
gcdExp := DocusaurusClusterDefinition.BuildClusterDefinitions()
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
```
