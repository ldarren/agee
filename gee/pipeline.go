package gee

import (
	"strings"
	"path"
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

func (p *Pipeline) getMiddlewares(url string, children []*Pipeline, middlewares []HandleFunc) string {
	for _, child := range children {
		if strings.HasPrefix(url, child.prefix) {
			middlewares = append(middlewares, child.middlewares...)
			return child.getMiddlewares(url, child.children, middlewares)
		}
	}
	return p.prefix
}

func (p *Pipeline) addRoute(method string, url string, handler HandleFunc) {
	p.router.add(method, path.Join(p.prefix, url), handler)
}

func (p *Pipeline) GET(url string, handler HandleFunc) {
	p.addRoute("GET", url, handler)
}

func (p *Pipeline) POST(url string, handler HandleFunc) {
	p.addRoute("POST", url, handler)
}
