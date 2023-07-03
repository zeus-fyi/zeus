package hades_core

import (
	"context"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (h *Hades) GetStatefulSetList(ctx context.Context, kns zeus_common_types.CloudCtxNs, filter *strings_filter.FilterOpts) (*v1.StatefulSetList, error) {
	h.SetContext(kns.Context)
	opts := metav1.ListOptions{}
	ssl, err := h.kc.AppsV1().StatefulSets(kns.Namespace).List(ctx, opts)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetStatefulSetList")
		return ssl, err
	}
	return ssl, err
}

func (h *Hades) GetStatefulSet(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) (*v1.StatefulSet, error) {
	h.SetContext(kns.Context)
	opts := metav1.GetOptions{}
	ss, err := h.kc.AppsV1().StatefulSets(kns.Namespace).Get(ctx, name, opts)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetStatefulSet")
		return ss, err
	}
	return ss, err
}

func (h *Hades) DeleteStatefulSet(ctx context.Context, kns zeus_common_types.CloudCtxNs, name string, filter *strings_filter.FilterOpts) error {
	h.SetContext(kns.Context)
	opts := metav1.DeleteOptions{}
	err := h.kc.AppsV1().StatefulSets(kns.Namespace).Delete(ctx, name, opts)
	if errors.IsNotFound(err) {
		log.Ctx(ctx).Info().Msg("not found, so doesn't exist here now")
		return nil
	}
	return err
}

func (h *Hades) CreateStatefulSet(ctx context.Context, kns zeus_common_types.CloudCtxNs, ss *v1.StatefulSet, filter *strings_filter.FilterOpts) (*v1.StatefulSet, error) {
	h.SetContext(kns.Context)
	opts := metav1.CreateOptions{}
	ss, err := h.kc.AppsV1().StatefulSets(kns.Namespace).Create(ctx, ss, opts)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateStatefulSet")
		return ss, err
	}
	return ss, err
}
