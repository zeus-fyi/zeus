package iris_proxy_rules_configs

type Routes struct {
	RouteID   int    `json:"routeID"`
	RoutePath string `json:"routePath"`
}

type RoutingGroups struct {
	RouteGroupID   int               `json:"routeGroupID"`
	RouteGroupName string            `json:"routeGroupName"`
	Map            map[string]Routes `json:"routingMap"`
	Slice          []Routes          `json:"routingSlice"`
}

func (i *Iris) RegisterRoutingEndpoints() {
	// TODO
}

func (i *Iris) ReadRoutingEndpoints() {
	// TODO
}

func (i *Iris) DeleteRoutingEndpoints() {
	// TODO
}
