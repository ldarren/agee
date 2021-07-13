package gee

type RouterGroup struct {
	prefix string
	middlewares []HandleFunc
	router *router
}

func NewRouterGroup(ctx *Web, prefix string) (*RouterGroup) {
	return &RouterGroup {
		prefix: prefix,
		middlewares: make([]HandleFunc, 0),
		router: ctx.router,
	}
}

func (g *RouterGroup) addRoute(method string, path string, handler HandleFunc) {
	g.router.add(method, g.prefix + path, handler)
}

func (g *RouterGroup) GET(path string, handler HandleFunc) {
	g.addRoute("GET", path, handler)
}

func (g *RouterGroup) POST(path string, handler HandleFunc) {
	g.addRoute("POST", path, handler)
}
