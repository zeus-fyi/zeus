package hades_core

import (
	"context"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetDeploymentList(ctx context.Context, kns zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.DeploymentList, error) {
	d, err := h.kc.AppsV1().Deployments(kns.Namespace).List(ctx, metav1.ListOptions{})
	return d, err
}

func (h *Hades) GetDeployment(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.Deployment, error) {
	d, err := h.kc.AppsV1().Deployments(kns.Namespace).Get(ctx, name, metav1.GetOptions{})
	return d, err
}

func (h *Hades) CreateDeployment(ctx context.Context, kns zeus_common_types.CloudCtxNs, d *v1.Deployment, filter *strings_filter.FilterOpts) (*v1.Deployment, error) {
	opts := metav1.CreateOptions{}
	d, err := h.kc.AppsV1().Deployments(kns.Namespace).Create(ctx, d, opts)
	return d, err
}

func (h *Hades) DeleteDeployment(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	opts := metav1.DeleteOptions{}
	err := h.kc.AppsV1().Deployments(kns.Namespace).Delete(ctx, name, opts)
	if errors.IsNotFound(err) {
		return nil
	}
	return err
}
