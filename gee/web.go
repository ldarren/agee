package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Web struct{
	routes router
}

func NewRouter() *Web{
	return &Web{
		routes: router{
			head: &node{},
			handlers: make(map[string]HandleFunc),
		},
	}
}

func (web *Web) addRoute(method string, path string, handler HandleFunc){
	web.routes.add(method, path, handler)
}

func (web *Web) GET(path string, handler HandleFunc){
	web.addRoute("GET",  path, handler)
}

func (web *Web) POST(path string, handler HandleFunc){
	web.addRoute("POST",  path, handler)
}

func (web *Web) ServeHTTP(res http.ResponseWriter, req *http.Request){
	handler, params := web.routes.get(req.Method, req.URL.Path)
	if nil != handler {
		handler(newContext(req, res, params))
	} else {
		fmt.Fprintf(res, "404 Not Found: %s\n", req.URL)
	}
}

func (web *Web) Start(port string) (err error){
	return http.ListenAndServe(port, web)
}
