package hades_core

import (
	"context"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetPodsUsingCtxNs(ctx context.Context, kubeCtxNs zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.PodList, error) {
	log.Ctx(ctx).Debug().Msg("GetPodsUsingCtxNs")
	pods, err := h.GetPods(ctx, kubeCtxNs.Namespace, metav1.ListOptions{})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetPodsUsingCtxNs: GetPods")
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

func (h *Hades) GetPods(ctx context.Context, ns string, opts metav1.ListOptions) (*v1.PodList, error) {
	return h.kc.CoreV1().Pods(ns).List(ctx, opts)
}
