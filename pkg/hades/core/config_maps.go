package hades_core

import (
	"context"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetConfigMapListWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.ConfigMapList, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().ConfigMaps(kns.Namespace).List(ctx, metav1.ListOptions{})
}

func (h *Hades) GetConfigMapWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.ConfigMap, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().ConfigMaps(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
}

func (h *Hades) CreateConfigMapWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, cm *v1.ConfigMap, filter *strings_filter.FilterOpts) (*v1.ConfigMap, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().ConfigMaps(kns.Namespace).Create(ctx, cm, metav1.CreateOptions{})
}

func (h *Hades) DeleteConfigMapWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	h.SetContext(kns.Context)
	err := h.kc.CoreV1().ConfigMaps(kns.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if errors.IsNotFound(err) {
		log.Ctx(ctx).Info().Msg("not found, so doesn't exist here now")
		return nil
	}
	return err
}
