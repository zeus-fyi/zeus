package hades_core

import (
	"context"

	"github.com/zeus-fyi/zeus/zeus/client/zeus_common_types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetNamespaces(ctx context.Context, kns zeus_common_types.CloudCtxNs) (*v1.NamespaceList, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
}

func (h *Hades) CreateNamespace(ctx context.Context, kns zeus_common_types.CloudCtxNs, namespace *v1.Namespace) (*v1.Namespace, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
}

func (h *Hades) DeleteNamespace(ctx context.Context, kns zeus_common_types.CloudCtxNs) error {
	h.SetContext(kns.Context)
	err := h.kc.CoreV1().Namespaces().Delete(ctx, kns.Namespace, metav1.DeleteOptions{})
	if errors.IsNotFound(err) {
		return nil
	}
	return err
}

func (h *Hades) GetNamespace(ctx context.Context, kns zeus_common_types.CloudCtxNs) (*v1.Namespace, error) {
	h.SetContext(kns.Context)
	return h.kc.CoreV1().Namespaces().Get(ctx, kns.Namespace, metav1.GetOptions{})
}

func (h *Hades) CreateNamespaceIfDoesNotExist(ctx context.Context, kns zeus_common_types.CloudCtxNs) (*v1.Namespace, error) {
	h.SetContext(kns.Context)
	ns, err := h.GetNamespace(ctx, kns)
	if errors.IsNotFound(err) {
		ns.Name = kns.Namespace
		return h.kc.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	}
	return ns, err
}
