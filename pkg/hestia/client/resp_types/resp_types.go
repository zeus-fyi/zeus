package hestia_resp_types

type Response struct {
	Message string `json:"message"`
}

type OrgRoutesResponse struct {
	Routes []string `json:"routes"`
}

type OrgGroupRoutesResponse struct {
	GroupName string   `json:"groupName"`
	Routes    []string `json:"routes"`
}

type OrgGroupsRoutesResponse struct {
	Map map[string][]string `json:"orgGroupsRoutes"`
}
