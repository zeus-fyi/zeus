package hestia_req_types

type IrisOrgGroupRoutesRequest struct {
	GroupName string   `json:"groupName,omitempty"`
	Routes    []string `json:"routes"`
}
