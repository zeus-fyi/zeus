package hades_core

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Hades struct {
	kc        *kubernetes.Clientset
	cfgAccess clientcmd.ConfigAccess
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

	cfgOveride := &clientcmd.ConfigOverrides{}
	if len(context) > 0 {
		cfgOveride = &clientcmd.ConfigOverrides{
			CurrentContext: context}
	}

	cc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: h.CfgPath},
		cfgOveride)
	h.cfgAccess = cc.ConfigAccess()
	h.clientCfg, err = cc.ClientConfig()
	if err != nil {
		log.Panic().Msg("Hades: SetContext, failed to set ClientConfig")
		panic(err)
	}
	h.SetClient(h.clientCfg)
}

func (h *Hades) SetClient(config *rest.Config) {
	var err error
	h.kc, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic().Msg("Hades: SetClient, failed to set client")
		panic(err)
	}
}

func (h *Hades) ConnectToK8s() {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	h.CfgPath = filepath.Join(home, ".kube", "config")
	h.SetContext("")
}

func (h *Hades) ConnectToK8sFromConfig(dir string) {
	h.CfgPath = dir
	h.SetContext("")
}

func (h *Hades) DefaultK8sCfgPath() string {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}
	return filepath.Join(home, ".kube", "config")
}
