package hades_core

import (
	"context"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetServiceMonitor(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.ServiceMonitor, error) {
	h.SetContext(kns.Context)
	sm, err := h.mc.MonitoringV1().ServiceMonitors(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("kns", kns).Str("name", name).Msg("GetServiceMonitor")
		return nil, err
	}
	return sm, err
}

func (h *Hades) CreateServiceMonitor(ctx context.Context, kns zeus_common_types.CloudCtxNs, sm *v1.ServiceMonitor, filter *strings_filter.FilterOpts) (*v1.ServiceMonitor, error) {
	h.SetContext(kns.Context)
	opts := metav1.CreateOptions{}
	sm, err := h.mc.MonitoringV1().ServiceMonitors(kns.Namespace).Create(ctx, sm, opts)
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("kns", kns).Msg("CreateServiceMonitor")
		return nil, err
	}
	return sm, err
}

func (h *Hades) DeleteServiceMonitor(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	h.SetContext(kns.Context)
	opts := metav1.DeleteOptions{}
	err := h.mc.MonitoringV1().ServiceMonitors(kns.Namespace).Delete(ctx, name, opts)
	if errors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("kns", kns).Str("name", name).Msg("DeleteServiceMonitor")
		return err
	}
	return nil
}
