package api

type api struct {
	router Router
}

func NewAPI(router Router) *api {
	return &api{router: router}
}

func (a *api) Start(addr string) error {
	var routes = RouterGroup{}
	routes = append(routes, homeRoutes...)
	a.addRoutes(routes)
	return a.router.Run(addr)
}

func (a *api) addRoutes(routes RouterGroup) {
	routes.forEach(func(route Route) {
		a.router.AddRoute(route.Method, route.Path, route.HandlerFunc)
	})
}
