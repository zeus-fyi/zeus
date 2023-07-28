package iris_proxy_rules_configs

import (
	"context"

	hestia_client "github.com/zeus-fyi/zeus/pkg/hestia/client"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	hestia_resp_types "github.com/zeus-fyi/zeus/pkg/hestia/client/resp_types"
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

func (i *Iris) CreateRoutingEndpoints(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) error {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	err := hc.CreateIrisRoutes(ctx, rr)
	if err != nil {
		return err
	}
	return nil
}

func (i *Iris) ReadRoutingEndpoints(ctx context.Context) (any, error) {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	resp, err := hc.ReadIrisRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) UpdateRoutingEndpoints(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) (any, error) {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	resp, err := hc.UpdateIrisRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) DeleteRoutingEndpoints(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) (any, error) {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	resp, err := hc.DeleteIrisRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) CreateRoutingGroup(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) error {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	err := hc.CreateIrisGroupRoutes(ctx, rr)
	if err != nil {
		return err
	}
	return nil
}

func (i *Iris) ReadRoutingGroupEndpoints(ctx context.Context, groupName string) (hestia_resp_types.OrgGroupRoutesResponse, error) {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	resp, err := hc.ReadIrisGroupRoutes(ctx, groupName)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (i *Iris) ReadAllRoutingGroupsEndpoints(ctx context.Context) (hestia_resp_types.OrgGroupsRoutesResponse, error) {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	resp, err := hc.ReadIrisGroupsRoutes(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (i *Iris) UpdateRoutingGroupEndpoints(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) (any, error) {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	resp, err := hc.UpdateIrisGroupRoutes(ctx, rr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *Iris) DeleteRoutingGroupEndpoints(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) error {
	hc := hestia_client.NewHestia(i.BaseURL, i.Token)
	err := hc.DeleteIrisGroupRoutes(ctx, rr)
	if err != nil {
		return err
	}
	return nil
}
