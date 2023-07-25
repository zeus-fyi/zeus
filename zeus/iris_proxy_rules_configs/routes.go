package iris_proxy_rules_configs

import (
	"context"

	hestia_client "github.com/zeus-fyi/zeus/pkg/hestia/client"
)

type Routes struct {
	RouteID   int    `json:"routeID,omitempty"`
	RoutePath string `json:"routePath"`
}

type RoutingGroups struct {
	RouteGroupName string            `json:"routeGroupName"`
	RouteGroupID   int               `json:"routeGroupID,omitempty"`
	Map            map[string]Routes `json:"routingMap,omitempty"`
	Slice          []Routes          `json:"routingSlice,omitempty"`
}

func (i *Iris) CreateRoutingEndpoints(ctx context.Context, routingEndpoints []string) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.CreateIrisRoutes(ctx, routingEndpoints)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) ReadRoutingEndpoints(ctx context.Context) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.ReadIrisRoutes(ctx, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) UpdateRoutingEndpoints(ctx context.Context, rr any) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.UpdateIrisRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) DeleteRoutingEndpoints(ctx context.Context, rr any) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.DeleteIrisRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) CreateRoutingGroup(ctx context.Context, rr any) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.CreateIrisGroupRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) ReadRoutingGroupEndpoints(ctx context.Context, rr any) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.ReadIrisGroupRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) UpdateRoutingGroupEndpoints(ctx context.Context, rr any) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.UpdateIrisGroupRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) DeleteRoutingGroupEndpoints(ctx context.Context, rr any) (any, error) {
	hc := hestia_client.NewDefaultHestiaClient(i.Token)
	resp, err := hc.DeleteIrisGroupRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
