package hades_core

import (
	"context"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetPod(ctx context.Context, name string, kns zeus_common_types.CloudCtxNs) (*v1.Pod, error) {
	h.SetContext(kns.Context)
	log.Ctx(ctx).Debug().Msg("GetPod")
	p, err := h.kc.CoreV1().Pods(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
	return p, err
}

func (h *Hades) GetPods(ctx context.Context, kns zeus_common_types.CloudCtxNs, opts metav1.ListOptions) (*v1.PodList, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().Pods(kns.Namespace).List(ctx, opts)
}

func (h *Hades) GetPodsUsingCtxNs(ctx context.Context, kubeCtxNs zeus_common_types.CloudCtxNs, logOpts *v1.PodLogOptions, filter *strings_filter.FilterOpts) (*v1.PodList, error) {
	log.Ctx(ctx).Debug().Msg("GetPodsUsingCtxNs")
	h.SetContext(kubeCtxNs.Context)

	if logOpts == nil {
		logOpts = &v1.PodLogOptions{}
	}
	pods, err := h.GetPods(ctx, kubeCtxNs, metav1.ListOptions{})
	if err != nil {
		return pods, err
	}

	if filter != nil {
		filteredPods := v1.PodList{}
		for _, pod := range pods.Items {
			if strings_filter.FilterStringWithOpts(pod.GetName(), filter) {
				filteredPods.Items = append(filteredPods.Items, pod)
			}
		}
		return &filteredPods, nil
	}

	return pods, err
}
