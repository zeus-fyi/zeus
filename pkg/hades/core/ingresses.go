package hades_core

import (
	"context"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetIngressListWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.IngressList, error) {
	h.SetContext(kns.Context)
	return h.kc.NetworkingV1().Ingresses(kns.Namespace).List(ctx, metav1.ListOptions{})
}

func (h *Hades) GetIngressWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.Ingress, error) {
	h.SetContext(kns.Context)
	return h.kc.NetworkingV1().Ingresses(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
}

func (h *Hades) CreateIngressWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, ing *v1.Ingress, filter *strings_filter.FilterOpts) (*v1.Ingress, error) {
	h.SetContext(kns.Context)
	return h.kc.NetworkingV1().Ingresses(kns.Namespace).Create(ctx, ing, metav1.CreateOptions{})
}

func (h *Hades) DeleteIngressWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	h.SetContext(kns.Context)
	err := h.kc.NetworkingV1().Ingresses(kns.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if errors.IsNotFound(err) {
		log.Ctx(ctx).Info().Msg("not found, so doesn't exist here now")
		return nil
	}
	return err
}
