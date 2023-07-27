package hestia_req_types

type CreateOrgRoutesRequest struct {
	Routes []string `json:"routes"`
}

type CreateOrgGroupRoutesRequest struct {
	GroupName string   `json:"groupName"`
	Routes    []string `json:"routes"`
}
