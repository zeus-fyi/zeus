package hades_core

import (
	"context"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetServiceListWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.ServiceList, error) {
	return h.kc.CoreV1().Services(kns.Namespace).List(ctx, metav1.ListOptions{})
}

func (h *Hades) GetServiceWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.Service, error) {
	return h.kc.CoreV1().Services(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
}

func (h *Hades) CreateServiceWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, s *v1.Service, filter *strings_filter.FilterOpts) (*v1.Service, error) {
	return h.kc.CoreV1().Services(kns.Namespace).Create(ctx, s, metav1.CreateOptions{})
}

func (h *Hades) DeleteServiceWithKns(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	err := h.kc.CoreV1().Services(kns.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if errors.IsNotFound(err) {
		return nil
	}
	return err
}
