package gee

import (
	"strings"
)

type HandleFunc func(ctx *Context)

type Pipeline struct {
	prefix string
	middlewares []HandleFunc
	router *router
	children []*Pipeline
}

func NewGroup(p *Pipeline, prefix string) (*Pipeline) {
	child := &Pipeline {
		prefix: p.prefix + prefix,
		middlewares: make([]HandleFunc, 0),
		router: p.router,
		children: make([]*Pipeline, 0),
	}
	p.children = append(p.children, child)
	return child
}

func (p *Pipeline) Use(middlewares ...HandleFunc) {
	p.middlewares = append(p.middlewares, middlewares...)
}

func (p *Pipeline) getMiddlewares(path string, children []*Pipeline, middlewares []HandleFunc) {
	for _, child := range children {
		if strings.HasPrefix(path, child.prefix) {
			middlewares = append(middlewares, child.middlewares...)
			child.getMiddlewares(path, child.children, middlewares)
			return
		}
	}
	return
}

func (p *Pipeline) addRoute(method string, path string, handler HandleFunc) {
	p.router.add(method, p.prefix + path, handler)
}

func (p *Pipeline) GET(path string, handler HandleFunc) {
	p.addRoute("GET", path, handler)
}

func (p *Pipeline) POST(path string, handler HandleFunc) {
	p.addRoute("POST", path, handler)
}
