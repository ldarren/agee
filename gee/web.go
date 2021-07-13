package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Web struct {
	router *router
}

func NewWeb() *Web {
	return &Web{
		router: newRouter(),
	}
}

func (web *Web) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	handler, params := web.router.get(req.Method, req.URL.Path)
	if nil != handler {
		handler(newContext(req, res, params))
	} else {
		fmt.Fprintf(res, "404 Not Found: %s\n", req.URL)
	}
}

func (web *Web) Start(port string) (err error) {
	return http.ListenAndServe(port, web)
}
