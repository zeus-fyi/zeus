package hades_core

import (
	"os"
	"path/filepath"

	monitoringclient "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Hades struct {
	kc        *kubernetes.Clientset
	cfgAccess clientcmd.ConfigAccess
	mc        *monitoringclient.Clientset
	kcCfg     clientcmd.ClientConfig
	clientCfg *rest.Config

	CfgPath string
}

type FilterOpts struct {
	DoesNotInclude []string
}

func (h *Hades) GetContexts() (map[string]*clientcmdapi.Context, error) {
	startingConfig, err := h.cfgAccess.GetStartingConfig()
	return startingConfig.Contexts, err
}

func (h *Hades) SetContext(context string) {
	var err error
	rc, err := h.kcCfg.RawConfig()
	if err != nil {
		log.Err(err)
		panic(err)
	}
	cc := clientcmd.NewNonInteractiveClientConfig(rc, context, nil, h.cfgAccess)
	h.cfgAccess = cc.ConfigAccess()
	h.clientCfg, err = cc.ClientConfig()
	if err != nil {
		log.Err(err)
		panic(err)
	}
	h.SetClient(h.clientCfg)
	mclient, err := monitoringclient.NewForConfig(h.clientCfg)
	if err != nil {
		log.Panic().Msg("Zeus: NewForConfig, failed to set client config")
		panic(err)
	}
	h.mc = mclient
}

func (h *Hades) SetClient(config *rest.Config) {
	var err error
	h.kc, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic().Msg("Hades: SetClient, failed to set client")
		panic(err)
	}
}

func (h *Hades) ConnectToK8sFromConfig(dir string) {
	h.CfgPath = dir
}

func (h *Hades) ConnectToK8sFromInMemFsCfgPath(fs memfs.MemFS) {
	log.Info().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath starting")
	var err error
	b, err := fs.ReadFile("/.kube/config")
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to read inmemfs kube config")
		panic(err)
	}
	cc, err := clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to set context")
		panic(err)
	}
	h.kcCfg = cc
	h.cfgAccess = cc.ConfigAccess()
	h.clientCfg, err = cc.ClientConfig()
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to set client config")
		panic(err)
	}
	h.SetClient(h.clientCfg)
	mclient, err := monitoringclient.NewForConfig(h.clientCfg)
	if err != nil {
		log.Panic().Msg("Zeus: NewForConfig, failed to set client config")
		panic(err)
	}
	h.mc = mclient
	log.Info().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath complete")
}

func (h *Hades) ConnectToK8s() {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}
	b, err := os.ReadFile(filepath.Join(home, ".kube", "config"))
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to read inmemfs kube config")
		panic(err)
	}
	cc, err := clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to set context")
		panic(err)
	}
	h.kcCfg = cc
	h.cfgAccess = cc.ConfigAccess()
	h.clientCfg, err = cc.ClientConfig()
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to set client config")
		panic(err)
	}
	h.SetClient(h.clientCfg)
	mclient, err := monitoringclient.NewForConfig(h.clientCfg)
	if err != nil {
		log.Panic().Msg("Zeus: NewForConfig, failed to set client config")
		panic(err)
	}
	h.mc = mclient
	log.Info().Msg("Zeus: DefaultK8sCfgPath complete")
	h.CfgPath = filepath.Join(home, ".kube", "config")
}

func (h *Hades) DefaultK8sCfgPath() string {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}
	b, err := os.ReadFile(filepath.Join(home, ".kube", "config"))
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to read inmemfs kube config")
		panic(err)
	}
	cc, err := clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to set context")
		panic(err)
	}
	h.kcCfg = cc
	h.cfgAccess = cc.ConfigAccess()
	h.clientCfg, err = cc.ClientConfig()
	if err != nil {
		log.Panic().Msg("Zeus: ConnectToK8sFromInMemFsCfgPath, failed to set client config")
		panic(err)
	}
	h.SetClient(h.clientCfg)
	mclient, err := monitoringclient.NewForConfig(h.clientCfg)
	if err != nil {
		log.Panic().Msg("Zeus: NewForConfig, failed to set client config")
		panic(err)
	}
	h.mc = mclient
	log.Info().Msg("Zeus: DefaultK8sCfgPath complete")
	return filepath.Join(home, ".kube", "config")
}
