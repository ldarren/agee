package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Web struct{
	router map[string]HandleFunc
}

func NewRouter() *Web{
	return &Web{router: make(map[string]HandleFunc)}
}

func (web *Web) addRoute(method string, path string, handler HandleFunc){
	key := method + ":" + path
	web.router[key] = handler
}

func (web *Web) GET(path string, handler HandleFunc){
	web.addRoute("GET",  path, handler)
}

func (web *Web) POST(path string, handler HandleFunc){
	web.addRoute("POST",  path, handler)
}

func (web *Web) ServeHTTP(res http.ResponseWriter, req *http.Request){
	key := req.Method + ":" + req.URL.Path
	if handler, ok := web.router[key]; ok {
		handler(newContext(req, res))
	} else {
		fmt.Fprintf(res, "404 Not Found: %s\n", req.URL)
	}
}

func (web *Web) Start(port string) (err error){
	return http.ListenAndServe(port, web)
}
