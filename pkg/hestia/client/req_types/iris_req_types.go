package hestia_req_types

import iris_programmable_proxy_v1_beta "github.com/zeus-fyi/zeus/zeus/iris_programmable_proxy/v1beta"

type IrisOrgGroupRoutesRequest struct {
	GroupName string   `json:"groupName,omitempty"`
	Routes    []string `json:"routes"`
}

type IrisRoutingProcedureRequest struct {
	ProcedureName         string                                                     `json:"procedureHeader"`
	OrderedProcedureSteps []iris_programmable_proxy_v1_beta.IrisRoutingProcedureStep `json:"orderedProcedureSteps,omitempty"`
}
