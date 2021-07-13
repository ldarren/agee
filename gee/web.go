package gee

import (
	"fmt"
	"net/http"
)

type Web struct {
	*Pipeline
}

func NewWeb() *Web {
	return &Web{
		Pipeline: &Pipeline{
			prefix: "",
			middlewares: make([]HandleFunc, 0),
			router: newRouter(),
			children: make([]*Pipeline, 0),
		},
	}
}

func (web *Web) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	handler, params := web.router.get(req.Method, path)
	if nil == handler {
		fmt.Fprintf(res, "404 Not Found: %s\n", req.URL)
		return
	}
	middlewares := make([]HandleFunc, 0)
	middlewares = append(middlewares, web.middlewares...)
	web.getMiddlewares(path, web.children, middlewares)
	middlewares = append(middlewares, handler)
	ctx := newContext(req, res, params, middlewares)
	ctx.Next()
}

func (web *Web) Start(port string) (err error) {
	return http.ListenAndServe(port, web)
}
