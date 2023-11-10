package zk8s_templates

import (
	"context"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type TemplateProcessorTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *TemplateProcessorTestSuite) TestGeneratePreview() {
	ctx := context.Background()

	req := Cluster{
		ClusterName:    "avaxNodeTest",
		ComponentBases: make(map[string]SkeletonBases),
		IngressSettings: Ingress{
			AuthServerURL: "aegis.zeus.fyi",
			Host:          "host.zeus.fyi",
		},
		IngressPaths: make(map[string]IngressPath),
	}

	m := make(map[string]string)
	m["start.sh"] = "#!/bin/sh\n    exec /avalanchego/build/avalanchego --db-dir=/data --http-host=0.0.0.0"
	sb := SkeletonBase{
		AddStatefulSet:    true,
		AddDeployment:     false,
		AddConfigMap:      true,
		AddService:        true,
		AddIngress:        true,
		AddServiceMonitor: false,
		ConfigMap:         m,
		StatefulSet: StatefulSet{
			ReplicaCount: 1,
			PVCTemplates: []PVCTemplate{{
				Name:               "avax-client-storage",
				AccessMode:         "ReadWriteOnce",
				StorageSizeRequest: "2Ti",
			}},
		},
		Containers: make(map[string]Container),
	}
	c := Container{
		IsInitContainer: false,
		DockerImage: DockerImage{
			ImageName: "avaplatform/avalanchego:v1.9.10",
			Cmd:       "/bin/sh",
			Args:      "-c,/scripts/start.sh",
			ResourceRequirements: ResourceRequirements{
				CPU:    "6",
				Memory: "12Gi",
			},
			Ports: []Port{
				{
					Name:               "p2p-tcp",
					Number:             "9651",
					Protocol:           "TCP",
					IngressEnabledPort: false,
				}, {
					Name:               "http-api",
					Number:             "9650",
					Protocol:           "TCP",
					IngressEnabledPort: true,
				}, {
					Name:               "metrics",
					Number:             "9090",
					Protocol:           "TCP",
					IngressEnabledPort: false,
				},
			},
			VolumeMounts: []VolumeMount{{
				Name:      "avax-client-storage",
				MountPath: "/data",
			}},
		},
	}

	sb.Containers["avax-client"] = c
	req.ComponentBases["avaxClients"] = make(map[string]SkeletonBase)
	req.ComponentBases["avaxClients"]["avaxClients"] = sb
	req.IngressPaths["avaxClients"] = IngressPath{
		Path:     "/",
		PathType: "ImplementationSpecific",
	}

	pcg, err := GenerateSkeletonBaseChartsPreview(ctx, req)
	t.Assert().NoError(err)
	t.Assert().NotEmpty(pcg)
}

func forceDirToCallerLocation() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}

func TestTemplateProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateProcessorTestSuite))
}
