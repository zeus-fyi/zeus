package hades_core

import (
	"context"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetDeploymentList(ctx context.Context, kns zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.DeploymentList, error) {
	h.SetContext(kns.Context)
	d, err := h.kc.AppsV1().Deployments(kns.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetDeploymentList")
		return d, err
	}
	return d, err
}

func (h *Hades) GetDeployment(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.Deployment, error) {
	h.SetContext(kns.Context)
	d, err := h.kc.AppsV1().Deployments(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetDeployment")
		return d, err
	}
	return d, err
}

func (h *Hades) CreateDeployment(ctx context.Context, kns zeus_common_types.CloudCtxNs, d *v1.Deployment, filter *strings_filter.FilterOpts) (*v1.Deployment, error) {
	h.SetContext(kns.Context)
	opts := metav1.CreateOptions{}
	d, err := h.kc.AppsV1().Deployments(kns.Namespace).Create(ctx, d, opts)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateDeployment")
		return d, err
	}
	return d, err
}

func (h *Hades) DeleteDeployment(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	h.SetContext(kns.Context)
	opts := metav1.DeleteOptions{}
	err := h.kc.AppsV1().Deployments(kns.Namespace).Delete(ctx, name, opts)
	if errors.IsNotFound(err) {
		log.Ctx(ctx).Info().Msg("not found, so doesn't exist here now")
		return nil
	}
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("DeleteDeployment")
		return err
	}
	return err
}
